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

	server := network.NewTCPServer(port, maxConnections)
	server.Start()
}
