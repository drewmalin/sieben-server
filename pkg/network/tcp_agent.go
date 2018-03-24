package network

import (
	"fmt"
	"net"

	"github.com/google/uuid"
	"github.com/sieben-server/pkg/logger"
)

// TCPAgent represents an individual TCP connection to a unique client.
type TCPAgent struct {
	id   string
	conn net.Conn
}

// NewTCPAgent instantiates a new TCPAgent.
func NewTCPAgent(conn net.Conn) *TCPAgent {
	id, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	agent := &TCPAgent{
		id:   id.String(),
		conn: conn,
	}
	return agent
}

func (agent *TCPAgent) String() string {
	return fmt.Sprintf("(%s - %s)", agent.id, agent.conn.RemoteAddr().String())
}

// Run is the business.
func (agent *TCPAgent) Run(server Server) {
	buffer := make([]byte, 1024)
	for {
		byteCount, err := agent.conn.Read(buffer)
		if err != nil {
			logger.Get().Printf("Agent %s disconnected by client", agent.String())
			server.OnClosedAgent(agent)
			break
		}
		message := make([]byte, byteCount)
		copy(message, buffer[:byteCount])
		server.OnReceivePacket(&Packet{agent.id, message})
	}
}

// Read reads data from this agent's connection into the specified byte array.
func (agent *TCPAgent) Read(packet *Packet) (int, error) {
	return agent.conn.Read(packet.Data)
}

// Write writes data into this agent's connection from the specified byte array.
func (agent *TCPAgent) Write(packet *Packet) (int, error) {
	if packet.SenderID != agent.id {
		return agent.conn.Write(packet.Data)
	}
	return 0, nil
}

// Close closes this agent's connection.
func (agent *TCPAgent) Close() error {
	return agent.conn.Close()
}
