package logging

import (
	"log"
	"os"
)

type Logger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

func NewLogger() *Logger {
	flags := log.LstdFlags | log.Lshortfile

	return &Logger{
		infoLogger:  log.New(os.Stdout, "INFO: ", flags),
		warnLogger:  log.New(os.Stdout, "WARN: ", flags),
		errorLogger: log.New(os.Stdout, "ERROR: ", flags),
	}
}

func (l *Logger) Info(message string) {
	l.infoLogger.Println(message)
}

func (l *Logger) Warn(message string) {
	l.warnLogger.Println(message)
}

func (l *Logger) Error(message string) {
	l.errorLogger.Println(message)
}
