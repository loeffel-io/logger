package logger

import (
	"github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
	"sync"
)

type Logger struct {
	Debug bool
	*sync.RWMutex
}

func (logger *Logger) GetDebug() bool {
	logger.RLock()
	defer logger.RUnlock()

	return logger.Debug
}

func (logger *Logger) SetDebug(debug bool) {
	logger.Lock()
	defer logger.Unlock()

	logger.Debug = debug
}

func (logger *Logger) Error(err error) {
	sentry.CaptureException(err)
	log.Fatal(err)
}

func (logger *Logger) Log(err error) {
	if !logger.GetDebug() {
		return
	}

	sentry.CaptureException(err)
	log.Warn(err)
}

func (logger *Logger) Print(msg string) {
	if !logger.GetDebug() {
		return
	}

	log.Info(msg)
}
