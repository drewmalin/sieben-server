package network

import (
	"bytes"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/sieben-server/pkg/logger"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
	end     = []byte{'\r', '\n'}
)

// WebSocketAgent represents an individual web socket connection to a unique client.
type WebSocketAgent struct {
	id   string
	conn *websocket.Conn
}

// NewWebSocketAgent instantiates a new WebSocketAgent.
func NewWebSocketAgent(conn *websocket.Conn) *WebSocketAgent {
	id, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	agent := &WebSocketAgent{
		id:   id.String(),
		conn: conn,
	}
	agent.conn.SetReadLimit(maxMessageSize)
	agent.conn.SetReadDeadline(time.Now().Add(pongWait))
	agent.conn.SetPongHandler(func(string) error {
		agent.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	return agent
}

func (agent *WebSocketAgent) String() string {
	return fmt.Sprintf("(%s - %s)", agent.id, agent.conn.RemoteAddr().String())
}

// Run is the business.
func (agent *WebSocketAgent) Run(server Server) {
	for {
		_, message, err := agent.conn.ReadMessage()
		if err != nil {
			logger.Get().Printf("Agent %s disconnected by client", agent.String())
			server.OnClosedAgent(agent)
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		server.OnReceivePacket(&Packet{agent.id, message})
	}
}

// Read reads data from this agent's connection into the specified byte array.
func (agent *WebSocketAgent) Read(packet *Packet) (int, error) {
	//return agent.conn.Read(packet.Data)
	return 0, nil
}

// Write writes data into this agent's connection from the specified byte array.
func (agent *WebSocketAgent) Write(packet *Packet) (int, error) {
	if packet.SenderID != agent.id {
		agent.conn.SetWriteDeadline(time.Now().Add(writeWait))
		writer, err := agent.conn.NextWriter(websocket.TextMessage)
		if err != nil {
			panic(err)
		}
		defer writer.Close()

		return writer.Write(append(packet.Data, end...))
	}
	return 0, nil
}

// Close closes this agent's connection.
func (agent *WebSocketAgent) Close() error {
	return agent.conn.Close()
}
