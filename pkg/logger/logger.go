package logger

import (
	"io"
	"log"
	"os"
	"sync"

	"github.com/spf13/viper"
)

type logger struct {
	*log.Logger
}

var loggerSingleton *logger
var once sync.Once

func Get() *logger {
	once.Do(func() {
		loggerSingleton = newLogger()
	})
	return loggerSingleton
}

func newLogger() *logger {
	var writer io.Writer
	if viper.GetString("log") == "STDOUT" {
		writer = os.Stdout
	} else {
		// create file
	}
	return &logger{
		Logger: log.New(writer, "", log.LstdFlags),
	}
}
