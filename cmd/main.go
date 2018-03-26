package main

import (
	"github.com/sieben-server/config"
	"github.com/sieben-server/pkg/network"
	"github.com/spf13/viper"
)

const (
	defaultConfigPath = "./config"
)

func main() {
	config.NewConfig(defaultConfigPath)

	port := viper.GetInt("port")
	maxConnections := viper.GetInt("maxConnections")

	if viper.GetString("protocol") == "WEBSOCKET" {
		server := network.NewWebSocketServer(port, maxConnections)
		server.Start()
	} else {
		server := network.NewTCPServer(port, maxConnections)
		server.Start()
	}

}
