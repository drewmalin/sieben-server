package network

// Server represents a handler of multiple connections.
type Server interface {
	Start()

	OnClosedAgent(agent Agent)

	OnReceivePacket(packet *Packet)

	Stop()
}
