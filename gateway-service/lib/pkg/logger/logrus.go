package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

func GetLogger() *logrus.Logger {

	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{})
	if os.Getenv("ENV") != "production" {
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				// s := strings.Split(f.Function, ".")
				// funcname := s[len(s)-1]
				_, filename := path.Split(f.File)
				return "", fmt.Sprintf("%s:%d", filename, f.Line)
			},
		})
	}
	log.SetReportCaller(true)
	log.Level = logrus.TraceLevel
	log.Out = os.Stderr
	return log
}
