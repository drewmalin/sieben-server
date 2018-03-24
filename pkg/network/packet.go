package network

// Packet represents a minimal bundle of data.
type Packet struct {
	SenderID string
	Data     []byte
}
