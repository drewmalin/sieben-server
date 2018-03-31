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
	writeWait      time.Duration = 10 * time.Second
	pongWait       time.Duration = 60 * time.Second
	pingInterval   time.Duration = (pongWait * 3) / 4
	maxMessageSize int64         = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
	end     = []byte{'\r', '\n'}
)

// WebSocketAgent represents an individual web socket connection to a unique client.
type WebSocketAgent struct {
	id           string
	conn         *websocket.Conn
	server       Server
	close        chan bool
	lastPingSent int64
	latency      int64
}

// NewWebSocketAgent instantiates a new WebSocketAgent.
func NewWebSocketAgent(conn *websocket.Conn, server Server) *WebSocketAgent {
	id, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	agent := &WebSocketAgent{
		id:     id.String(),
		conn:   conn,
		server: server,
		close:  make(chan bool, 1),
	}

	agent.conn.SetReadLimit(maxMessageSize)
	agent.conn.SetReadDeadline(time.Now().Add(pongWait))
	agent.conn.SetPongHandler(func(appData string) error {
		agent.latency = (time.Now().UnixNano() - agent.lastPingSent) / int64(time.Millisecond)
		agent.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	return agent
}

// ID returns the unique ID of this agent.
func (agent *WebSocketAgent) ID() string {
	return agent.id
}

// Run is the business.
func (agent *WebSocketAgent) Run() {
	go agent.ping()
	agent.handleReceivePackets()
}

// Read reads data from this agent's connection into the specified byte array.
func (agent *WebSocketAgent) Read(packet *Packet) (int, error) {
	//return agent.conn.Read(packet.Data)
	return 0, nil
}

// Write writes data into this agent's connection from the specified byte array.
func (agent *WebSocketAgent) Write(packet *Packet) (int, error) {
	if packet.SenderID != agent.id {
		message := append(packet.Data, end...)
		if err := agent.writeMessage(websocket.TextMessage, message); err != nil {
			return 0, err
		}
		return len(packet.Data), nil
	}
	return 0, nil
}

// Close closes this agent's connection.
func (agent *WebSocketAgent) Close() error {
	agent.close <- true
	return agent.conn.Close()
}

func (agent *WebSocketAgent) String() string {
	return fmt.Sprintf("WebSocketAgent: %s, Connected to: %s, Latency: %d ms",
		agent.id, agent.conn.RemoteAddr().String(), agent.latency)
}

// Periodically sents a PingMesssage via this agent's websocket connection.
// Does not return.
// Should be invoked as a go routine.
func (agent *WebSocketAgent) ping() {
	ticker := time.NewTicker(pingInterval)
	defer func() {
		ticker.Stop()
	}()
	for {
		select {
		case <-ticker.C:
			agent.lastPingSent = time.Now().UnixNano()
			if err := agent.writeMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		case <-agent.close:
			return
		}
	}
}

// Handles all incoming packet data.
// Returns only upon websocket disconnection.
func (agent *WebSocketAgent) handleReceivePackets() {
	for {
		_, message, err := agent.conn.ReadMessage()
		if err != nil {
			logger.Get().Printf("Agent %s connection closed by client", agent.ID())
			agent.server.OnClosedAgent(agent)
			return
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		agent.server.OnReceivePacket(&Packet{agent.id, message})
	}
}

// Writes a message for a given message type.
// Closes the agent if a failure is enountered during write.
func (agent *WebSocketAgent) writeMessage(messageType int, message []byte) error {
	agent.conn.SetWriteDeadline(time.Now().Add(writeWait))
	if err := agent.conn.WriteMessage(messageType, message); err != nil {
		logger.Get().Printf("Agent %s failed to receive timely response from client", agent.ID())
		agent.server.OnClosedAgent(agent)
		return err
	}
	return nil
}
