package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewFileLogger(fileLocation string, level logrus.Level, formatter logrus.Formatter) *logrus.Logger {
	output, _ := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0755)
	return NewLogger(output, level, formatter)
}

func JsonFileLogger(fileLocation string, level logrus.Level) *logrus.Logger {
	return NewFileLogger(fileLocation, level, &logrus.JSONFormatter{})
}

func TextFileLogger(fileLocation string, level logrus.Level) *logrus.Logger {
	return NewFileLogger(fileLocation, level, &logrus.TextFormatter{})
}
