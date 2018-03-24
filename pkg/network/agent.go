package network

// Agent represents an individual connection to a unique client.
type Agent interface {
	Run(server Server)

	Read(packet *Packet) (int, error)

	Write(packet *Packet) (int, error)

	String() string

	Close() error
}
