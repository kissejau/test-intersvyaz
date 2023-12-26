package logger

import (
	"log"
	"os"
)

type Logger struct {
	logInfo *log.Logger
	logWarn *log.Logger
	logErr  *log.Logger
}

func NewLogger() *Logger {
	flags := log.Ldate | log.Ltime
	return &Logger{
		logInfo: log.New(os.Stdout, "INFO: ", flags),
		logWarn: log.New(os.Stdout, "WARN: ", flags),
		logErr:  log.New(os.Stdout, "ERROR: ", flags),
	}
}

func (l Logger) Info(format string) {
	l.logInfo.Print(format)
}

func (l Logger) Infof(format string, args ...any) {
	l.logInfo.Printf(format, args...)
}

func (l Logger) Warn(format string) {
	l.logWarn.Print(format)
}

func (l Logger) Warnf(format string, args ...any) {
	l.logWarn.Printf(format, args...)
}

func (l Logger) Error(format string) {
	l.logErr.Fatal(format)
}

func (l Logger) Errorf(format string, args ...any) {
	l.logErr.Fatalf(format, args...)
}
