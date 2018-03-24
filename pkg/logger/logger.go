package logger

import (
	"log"
	"os"
	"sync"
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
	return &logger{
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}
