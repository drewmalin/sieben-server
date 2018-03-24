package main

import (
	"github.com/sieben-server/pkg/network"
)

const (
	port           = 9123
	maxConnections = 16
)

func main() {
	server := network.NewTCPServer(port, maxConnections)
	server.Start()
}
