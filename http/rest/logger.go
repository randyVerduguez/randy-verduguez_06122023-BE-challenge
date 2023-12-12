package rest

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func NewLogger() (*logrus.Logger, error) {
	log := logrus.New()
	absPath, err := filepath.Abs("../../logs")

	if err != nil {
		error := fmt.Errorf("Error reading given path: %s", err)
		return log, error
	}

	filename := fmt.Sprintf("%s/%s", absPath, "http-log.log")
	flags := os.O_RDWR | os.O_CREATE | os.O_TRUNC
	logFile, err := os.OpenFile(filename, flags, 0666)

	if err != nil {
		error := fmt.Errorf("Error opening file: %s", err)
		return log, error
	}

	log.SetOutput(logFile)
	log.SetLevel(logrus.TraceLevel)
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:      false,
		DisableTimestamp: true,
	})
	return log, nil
}
