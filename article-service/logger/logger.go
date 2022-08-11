package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

type IEntry interface {
	Trace(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}

type ICustomLogger interface {
	Setup()
	WithFields(fields map[string]interface{}) IEntry

	Trace(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
}

func InitLogger(config *LoggerConfig) *CustomLogger {
	logger := &CustomLogger{
		Config:  config,
		Adapter: logrus.New(),
	}
	logger.Setup()

	return logger
}

func (l *CustomLogger) Setup() {
	l.Adapter.SetFormatter(&logrus.JSONFormatter{})

	if l.Config.TimestampFormat == "" {
		l.Config.TimestampFormat = "2006-01-02 15:04:05"
	}
	if l.Config.Writer == nil {
		l.Config.Writer = os.Stdout
	}

	l.Adapter.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   l.Config.Fulltimestamp,
		TimestampFormat: l.Config.TimestampFormat,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			_, filename := path.Split(f.File)
			return "", fmt.Sprintf(" %s:%d -- ", filename, f.Line)
		},
	})

	l.Adapter.SetReportCaller(true)
	l.Adapter.Level = logrus.Level(l.Config.Level)
	l.Adapter.Out = l.Config.Writer
}

func (l *CustomLogger) WithFields(fields map[string]interface{}) IEntry {
	return l.Adapter.WithFields(fields)
}

func (l *CustomLogger) Trace(args ...interface{}) {
	l.Adapter.Trace(args...)
}

func (l *CustomLogger) Debug(args ...interface{}) {
	l.Adapter.Debug(args...)
}

func (l *CustomLogger) Info(args ...interface{}) {
	l.Adapter.Info(args...)
}

func (l *CustomLogger) Warn(args ...interface{}) {
	l.Adapter.Warn(args...)
}

func (l *CustomLogger) Error(args ...interface{}) {
	l.Adapter.Error(args...)
}

func (l *CustomLogger) Fatal(args ...interface{}) {
	l.Adapter.Fatal(args...)
}

func (l *CustomLogger) Panic(args ...interface{}) {
	l.Adapter.Panic(args...)
}

func (l *CustomLogger) Tracef(format string, args ...interface{}) {
	l.Adapter.Tracef(format, args...)
}

func (l *CustomLogger) Debugf(format string, args ...interface{}) {
	l.Adapter.Debugf(format, args...)
}

func (l *CustomLogger) Infof(format string, args ...interface{}) {
	l.Adapter.Infof(format, args...)
}

func (l *CustomLogger) Warnf(format string, args ...interface{}) {
	l.Adapter.Warnf(format, args...)
}

func (l *CustomLogger) Errorf(format string, args ...interface{}) {
	l.Adapter.Errorf(format, args...)
}

func (l *CustomLogger) Fatalf(format string, args ...interface{}) {
	l.Adapter.Fatalf(format, args...)
}

func (l *CustomLogger) Panicf(format string, args ...interface{}) {
	l.Adapter.Panicf(format, args...)
}
