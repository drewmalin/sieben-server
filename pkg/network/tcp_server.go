package network

import (
	"fmt"
	"net"

	"github.com/sieben-server/pkg/logger"
)

const (
	protocol      = "tcp"
	addressFormat = "localhost:%d"
)

// TCPServer is a Server which handles its connections via TCP.
type TCPServer struct {
	port           int
	maxConnections int
	agentSet       *AgentSet

	newConnChannel   chan net.Conn
	deadAgentChannel chan Agent
	packetChannel    chan *Packet
}

// NewTCPServer instantiates a new TCPServer.
func NewTCPServer(port int, maxConnections int) *TCPServer {
	agentSet := NewAgentSet()
	newConnChannel := make(chan net.Conn, maxConnections)
	deadAgentChannel := make(chan Agent, maxConnections)
	packetChannel := make(chan *Packet, maxConnections)

	return &TCPServer{
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
func (server *TCPServer) Start() {
	server.logStatus()

	listener, err := net.Listen(protocol, fmt.Sprintf(addressFormat, server.port))
	if err != nil {
		panic(err)
	}

	go server.handleNewConnections(&listener)

	for {
		select {
		case conn := <-server.newConnChannel:

			tcpAgent := NewTCPAgent(conn, server)
			server.agentSet.Add(tcpAgent)
			logger.Get().Printf("Agent %s opened", tcpAgent.ID())
			server.logStatus()

			go tcpAgent.Run()

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

func (server *TCPServer) handleNewConnections(listener *net.Listener) {
	for {
		connection, err := (*listener).Accept()
		if err != nil {
			panic(err)
		}
		if server.agentSet.Size() >= server.maxConnections {
			logger.Get().Println("Connection attempted while agent count is at capacity. Connection will be closed.")
			connection.Close()
		} else {
			server.newConnChannel <- connection
		}
	}
}

// OnClosedAgent handles the closure of an agent's connection.
func (server *TCPServer) OnClosedAgent(agent Agent) {
	server.deadAgentChannel <- agent
}

// OnReceivePacket handles the receipt of a new packet.
func (server *TCPServer) OnReceivePacket(packet *Packet) {
	server.packetChannel <- packet
}

// Stop closes all agents associated with this server.
func (server *TCPServer) Stop() {
	for agent := range server.agentSet.Agents {
		agent.Close()
	}
}

func (server *TCPServer) logStatus() {
	logger.Get().Printf("TCPServer status:")
	logger.Get().Printf("\tCurrent agent count: %d / %d", server.agentSet.Size(), server.maxConnections)
	for agent := range server.agentSet.Agents {
		logger.Get().Printf("\t%s", agent.String())
	}
}
