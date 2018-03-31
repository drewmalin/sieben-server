package network

import (
	"net"
	"testing"
)

func TestNewTCPAgent(t *testing.T) {
	mockServer := new(MockServer)
	mockConnection := new(MockConn)

	tcpAgent1 := NewTCPAgent(mockConnection, mockServer)
	tcpAgent2 := NewTCPAgent(mockConnection, mockServer)

	var actualConn net.Conn
	var expectedConn net.Conn

	actualConn = tcpAgent1.conn
	expectedConn = mockConnection

	if actualConn != expectedConn {
		t.Errorf("Conn was expected to be: %s, but was: %s.", expectedConn, actualConn)
	}
	if tcpAgent1.id == "" {
		t.Errorf("ID was unexpectedly empty.")
	}

	actualConn = tcpAgent2.conn
	expectedConn = mockConnection

	if actualConn != expectedConn {
		t.Errorf("Conn was expected to be: %s, but was: %s.", expectedConn, actualConn)
	}
	if tcpAgent2.id == "" {
		t.Errorf("ID was unexpectedly empty.")
	}

	if tcpAgent1.id == tcpAgent2.id {
		t.Errorf("ID1 and ID2 were unexpectedly equal.")
	}
}
