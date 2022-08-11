package logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

type CustomLogger struct {
	Config  *LoggerConfig
	Adapter *logrus.Logger
}

type Level uint32

type Fields logrus.Fields

const ( // Logrus standard
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

type LoggerConfig struct {
	Level           Level     // Log level (debug, info, warn, error, fatal)
	Fulltimestamp   bool      // true if full timestamp shown
	TimestampFormat string    // timestamp format
	Writer          io.Writer // Log writer
}
