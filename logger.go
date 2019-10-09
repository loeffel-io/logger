package logger

import (
	"github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
)

type Logger struct {
	Debug bool
}

func (logger *Logger) Error(err error) {
	sentry.CaptureException(err)
	log.Fatal(err)
}

func (logger *Logger) Log(err error) {
	if !logger.Debug {
		return
	}

	sentry.CaptureException(err)
	log.Warn(err)
}

func (logger *Logger) Print(msg string) {
	if !logger.Debug {
		return
	}

	log.Info(msg)
}
