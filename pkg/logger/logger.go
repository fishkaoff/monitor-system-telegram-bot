package logger

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

type Logger struct {
	sugar *zap.SugaredLogger
}

func (l *Logger) Start() *zap.Logger {
	logger, _ := zap.NewProduction()
	l.sugar = logger.Sugar()
	return logger
}

func (l *Logger) LogMessage(message string) {
	currentTime := time.Now()
	messageForLog := fmt.Sprintf("***%s*** [%s]: %s", "MESSAGE", currentTime, message)
	l.sugar.Info(messageForLog)
}

func (l *Logger) LogError(err string) {
	currentTime := time.Now()
	messageForLog := fmt.Sprintf("***%s*** [%s]: %s", "ERROR", currentTime, err)
	l.sugar.Info(messageForLog)
}

func (l *Logger) LogErrorAndQuit(err string) {
	currentTime := time.Now()
	messageForLog := fmt.Sprintf("***%s*** [%s]: %s", "ERROR", currentTime, err)
	panic(messageForLog)
}
