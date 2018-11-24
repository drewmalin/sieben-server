package network

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/drewmalin/sieben-server/pkg/logger"
)

// WebSocketServer is a Server which handles its connections via web sockets.
type WebSocketServer struct {
	port           int
	maxConnections int
	agentSet       *AgentSet

	newConnChannel   chan *websocket.Conn
	deadAgentChannel chan Agent
	packetChannel    chan *Packet
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// NewWebSocketServer instantiates a new WebSocketServer.
func NewWebSocketServer(port int, maxConnections int) *WebSocketServer {
	agentSet := NewAgentSet()
	newConnChannel := make(chan *websocket.Conn, maxConnections)
	deadAgentChannel := make(chan Agent, maxConnections)
	packetChannel := make(chan *Packet, maxConnections)

	return &WebSocketServer{
		port:             port,
		maxConnections:   maxConnections,
		agentSet:         agentSet,
		newConnChannel:   newConnChannel,
		deadAgentChannel: deadAgentChannel,
		packetChannel:    packetChannel,
	}
}

// Start informs this server to begin listening on its port. New connections will be passed to new agents for handling,
// and the agent lifecycle will be managed. Start will not return.
func (server *WebSocketServer) Start() {
	server.logStatus()

	go server.handleNewConnections()

	for {
		select {
		case conn := <-server.newConnChannel:
			webSocketAgent := NewWebSocketAgent(conn, server)
			logger.Get().Printf("Agent %s opened", webSocketAgent.ID())

			server.agentSet.Add(webSocketAgent)
			server.logStatus()

			go webSocketAgent.Run()

		case agent := <-server.deadAgentChannel:
			server.agentSet.Remove(agent)
			agent.Close()
			logger.Get().Printf("Agent %s closed", agent.ID())

			server.logStatus()

		case packet := <-server.packetChannel:
			for agent := range server.agentSet.Agents {
				agent.Write(packet)
			}
		}
	}
}

func (server *WebSocketServer) handleNewConnections() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		conn, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			logger.Get().Println(err)
			panic(err)
		}
		if server.agentSet.Size() >= server.maxConnections {
			logger.Get().Println("Connection attempted while agent count is at capacity. Connection will be closed.")
			conn.Close()
		} else {
			server.newConnChannel <- conn
		}
	})
	err := http.ListenAndServe(fmt.Sprintf(":%d", server.port), nil)
	if err != nil {
		panic(err)
	}
}

// OnClosedAgent handles the closure of an agent's connection.
func (server *WebSocketServer) OnClosedAgent(agent Agent) {
	server.deadAgentChannel <- agent
}

// OnReceivePacket handles the receipt of a new packet.
func (server *WebSocketServer) OnReceivePacket(packet *Packet) {
	server.packetChannel <- packet
}

// Stop closes all agents associated with this server.
func (server *WebSocketServer) Stop() {
	for agent := range server.agentSet.Agents {
		agent.Close()
	}
}

func (server *WebSocketServer) logStatus() {
	logger.Get().Printf("WebSocketServer status:")
	logger.Get().Printf("\tCurrent agent count: %d / %d", server.agentSet.Size(), server.maxConnections)
	for agent := range server.agentSet.Agents {
		logger.Get().Printf("\t%s", agent.String())
	}
}
