package network

import (
	"net"
	"time"
)

// MockConn mocks the functionality of net.Conn
type MockConn struct {
}

// Read ...
func (m *MockConn) Read(data []byte) (int, error) {
	return 0, nil
}

// Write ...
func (m *MockConn) Write(data []byte) (int, error) {
	return 0, nil
}

// Close ...
func (m *MockConn) Close() error {
	return nil
}

// LocalAddr ...
func (m *MockConn) LocalAddr() net.Addr {
	return nil
}

// RemoteAddr ...
func (m *MockConn) RemoteAddr() net.Addr {
	return nil
}

// SetDeadline ...
func (m *MockConn) SetDeadline(t time.Time) error {
	return nil
}

// SetReadDeadline ...
func (m *MockConn) SetReadDeadline(t time.Time) error {
	return nil
}

// SetWriteDeadline ...
func (m *MockConn) SetWriteDeadline(t time.Time) error {
	return nil
}
