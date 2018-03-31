package network

// MockServer mocks the functionality of network.Server
type MockServer struct {
}

// Start ...
func (m *MockServer) Start() {
}

// OnClosedAgent ...
func (m *MockServer) OnClosedAgent(agent Agent) {
}

// OnReceivePacket ...
func (m *MockServer) OnReceivePacket(packet *Packet) {
}

// Stop ...
func (m *MockServer) Stop() {
}
