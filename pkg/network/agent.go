package network

// Agent represents an individual connection to a unique client.
type Agent interface {
	ID() string

	Run()

	Read(packet *Packet) (int, error)

	Write(packet *Packet) (int, error)

	Close() error

	String() string
}
