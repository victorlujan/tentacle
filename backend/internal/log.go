package internal

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
	once   sync.Once
)

func NewLoger(logFile string) *logrus.Logger {
	once.Do(func() {
		logger = logrus.New()
		f, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
		if err != nil {
			panic("Error creating log file" + err.Error())
		}

		logger.SetOutput(f)
		logger.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})
	})

	return logger
}
