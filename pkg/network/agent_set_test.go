package network

import (
	"testing"
)

func TestAdd(t *testing.T) {
	mockServer := new(MockServer)
	mockConnection := new(MockConn)

	tcpAgent1 := NewTCPAgent(mockConnection, mockServer)
	tcpAgent2 := NewTCPAgent(mockConnection, mockServer)

	tcpAgentSet := NewAgentSet()

	var actualResult bool
	var actualSize int
	var expectedResult bool
	var expectedSize int

	if tcpAgentSet.Size() != 0 {
		t.Errorf("Size was expected to be: %d, but was: %d.", 0, tcpAgentSet.Size())
	}

	actualResult = tcpAgentSet.Add(tcpAgent1)
	actualSize = tcpAgentSet.Size()
	expectedResult = true
	expectedSize = 1

	if actualSize != expectedSize {
		t.Errorf("Size was expected to be: %d, but was: %d.", expectedSize, actualSize)
	}
	if actualResult != expectedResult {
		t.Errorf("Added was expected to be: %t, but was: %t.", expectedResult, actualResult)
	}

	actualResult = tcpAgentSet.Add(tcpAgent1)
	actualSize = tcpAgentSet.Size()
	expectedResult = false
	expectedSize = 1

	if actualSize != expectedSize {
		t.Errorf("Size was expected to be: %d, but was: %d.", expectedSize, actualSize)
	}
	if actualResult != expectedResult {
		t.Errorf("Added was expected to be: %t, but was: %t.", expectedResult, actualResult)
	}

	actualResult = tcpAgentSet.Add(tcpAgent2)
	actualSize = tcpAgentSet.Size()
	expectedResult = true
	expectedSize = 2

	if actualSize != expectedSize {
		t.Errorf("Size was expected to be: %d, but was: %d.", expectedSize, actualSize)
	}
	if actualResult != expectedResult {
		t.Errorf("Added was expected to be: %t, but was: %t.", expectedResult, actualResult)
	}
}

func TestRemove(t *testing.T) {
	mockServer := new(MockServer)
	mockConnection := new(MockConn)

	tcpAgent1 := NewTCPAgent(mockConnection, mockServer)
	tcpAgent2 := NewTCPAgent(mockConnection, mockServer)

	tcpAgentSet := NewAgentSet()

	var actualResult bool
	var actualSize int
	var expectedResult bool
	var expectedSize int

	tcpAgentSet.Add(tcpAgent1)
	tcpAgentSet.Add(tcpAgent2)

	if tcpAgentSet.Size() != 2 {
		t.Errorf("Size was expected to be: %d, but was: %d.", 2, tcpAgentSet.Size())
	}

	actualResult = tcpAgentSet.Remove(tcpAgent1)
	actualSize = tcpAgentSet.Size()
	expectedResult = true
	expectedSize = 1

	if actualSize != expectedSize {
		t.Errorf("Size was expected to be: %d, but was: %d.", expectedSize, actualSize)
	}
	if actualResult != expectedResult {
		t.Errorf("Added was expected to be: %t, but was: %t.", expectedResult, actualResult)
	}

	actualResult = tcpAgentSet.Remove(tcpAgent1)
	actualSize = tcpAgentSet.Size()
	expectedResult = false
	expectedSize = 1

	if actualSize != expectedSize {
		t.Errorf("Size was expected to be: %d, but was: %d.", expectedSize, actualSize)
	}
	if actualResult != expectedResult {
		t.Errorf("Added was expected to be: %t, but was: %t.", expectedResult, actualResult)
	}

	actualResult = tcpAgentSet.Remove(tcpAgent2)
	actualSize = tcpAgentSet.Size()
	expectedResult = true
	expectedSize = 0

	if actualSize != expectedSize {
		t.Errorf("Size was expected to be: %d, but was: %d.", expectedSize, actualSize)
	}
	if actualResult != expectedResult {
		t.Errorf("Added was expected to be: %t, but was: %t.", expectedResult, actualResult)
	}
}

func TestContains(t *testing.T) {
	mockServer := new(MockServer)
	mockConnection := new(MockConn)

	tcpAgent1 := NewTCPAgent(mockConnection, mockServer)
	tcpAgent2 := NewTCPAgent(mockConnection, mockServer)

	tcpAgentSet := NewAgentSet()

	var actualContains1 bool
	var actualContains2 bool
	var expectedContains1 bool
	var expectedContains2 bool

	actualContains1 = tcpAgentSet.Contains(tcpAgent1)
	actualContains2 = tcpAgentSet.Contains(tcpAgent2)
	expectedContains1 = false
	expectedContains2 = false

	if actualContains1 != expectedContains1 {
		t.Errorf("Contains was expected to be: %t, but was: %t", expectedContains1, actualContains1)
	}
	if actualContains2 != expectedContains2 {
		t.Errorf("Contains was expected to be: %t, but was: %t", expectedContains2, actualContains2)
	}

	tcpAgentSet.Add(tcpAgent1)

	actualContains1 = tcpAgentSet.Contains(tcpAgent1)
	actualContains2 = tcpAgentSet.Contains(tcpAgent2)
	expectedContains1 = true
	expectedContains2 = false

	if actualContains1 != expectedContains1 {
		t.Errorf("Contains was expected to be: %t, but was: %t", expectedContains1, actualContains1)
	}
	if actualContains2 != expectedContains2 {
		t.Errorf("Contains was expected to be: %t, but was: %t", expectedContains2, actualContains2)
	}

	tcpAgentSet.Add(tcpAgent2)

	actualContains1 = tcpAgentSet.Contains(tcpAgent1)
	actualContains2 = tcpAgentSet.Contains(tcpAgent2)
	expectedContains1 = true
	expectedContains2 = true

	if actualContains1 != expectedContains1 {
		t.Errorf("Contains was expected to be: %t, but was: %t", expectedContains1, actualContains1)
	}
	if actualContains2 != expectedContains2 {
		t.Errorf("Contains was expected to be: %t, but was: %t", expectedContains2, actualContains2)
	}

	tcpAgentSet.Remove(tcpAgent1)

	actualContains1 = tcpAgentSet.Contains(tcpAgent1)
	actualContains2 = tcpAgentSet.Contains(tcpAgent2)
	expectedContains1 = false
	expectedContains2 = true

	if actualContains1 != expectedContains1 {
		t.Errorf("Contains was expected to be: %t, but was: %t", expectedContains1, actualContains1)
	}
	if actualContains2 != expectedContains2 {
		t.Errorf("Contains was expected to be: %t, but was: %t", expectedContains2, actualContains2)
	}

	tcpAgentSet.Remove(tcpAgent2)

	actualContains1 = tcpAgentSet.Contains(tcpAgent1)
	actualContains2 = tcpAgentSet.Contains(tcpAgent2)
	expectedContains1 = false
	expectedContains2 = false

	if actualContains1 != expectedContains1 {
		t.Errorf("Contains was expected to be: %t, but was: %t", expectedContains1, actualContains1)
	}
	if actualContains2 != expectedContains2 {
		t.Errorf("Contains was expected to be: %t, but was: %t", expectedContains2, actualContains2)
	}
}
