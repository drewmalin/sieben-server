package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// NewConfig re-instantiates viper using the specified config filepath.
func NewConfig(path string) {
	viper.SetConfigName("config")
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Error reading config file: %s", err))
	}

	viper.SetDefault("port", 8080)
	viper.SetDefault("maxConnections", 16)
	viper.SetDefault("log", "FILE")
}
