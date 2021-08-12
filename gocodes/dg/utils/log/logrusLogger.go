package log

import "github.com/sirupsen/logrus"

func NewLogrusLogger() *logrus.Logger {
	formatter := logrus.TextFormatter{
		FullTimestamp: true,
	}
	logger := logrus.New()
	logger.Formatter = &formatter
	return logger
}
