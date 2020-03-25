package logger

import (
	"github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type Logger struct {
	SentryHub *sentry.Hub
	Debug     bool
	*sync.RWMutex
}

func (logger *Logger) GetSentryHub() *sentry.Hub {
	logger.RLock()
	defer logger.RUnlock()

	return logger.SentryHub
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
	defer logger.GetSentryHub().Flush(2 * time.Second)

	logger.GetSentryHub().CaptureException(err)
	log.Fatal(err)
}

func (logger *Logger) Log(err error) {
	logger.GetSentryHub().CaptureException(err)

	if logger.GetDebug() {
		log.Warn(err)
	}
}

func (logger *Logger) Print(msg string) {
	logger.GetSentryHub().CaptureMessage(msg)

	if logger.GetDebug() {
		log.Info(msg)
	}
}
