package logger

import (
	"log"
	"sync"
)

type Logger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Debugln(args ...interface{})
	Debugw(msg string, keysAndValues ...interface{})

	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Infoln(args ...interface{})
	Infow(msg string, keysAndValues ...interface{})

	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Warnln(args ...interface{})
	Warnw(msg string, keysAndValues ...interface{})

	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Errorln(args ...interface{})
	Errorw(msg string, keysAndValues ...interface{})

	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	DPanicln(args ...interface{})
	DPanicw(msg string, keysAndValues ...interface{})

	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
	Fatalln(args ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
}

var (
	singleLogger Logger
	once         sync.Once
)

func SetLogger(newLogger Logger) {
	once.Do(func() {
		singleLogger = newLogger
	})
}

func Log() Logger {
	if singleLogger == nil {
		log.Fatalf("can't use logger without setting it")
	}
	return singleLogger
}
