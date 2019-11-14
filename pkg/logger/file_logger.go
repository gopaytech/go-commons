package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var logger *logrus.Logger

func NewFileLogger(fileLocation string, level logrus.Level) *logrus.Logger {
	logger = logrus.New()
	output, _ := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0755)
	logger.SetOutput(output)
	logger.SetLevel(level)
	return logger
}
