package logger

import (
	"log"
	"os"
)

type Logger struct {
	infoLog *log.Logger
	errorLog *log.Logger
}

func New() *Logger {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return &Logger{infoLog: infoLog, errorLog: errorLog}
}

func (l Logger) Info(msg string) {
	l.infoLog.Println(msg)
}

func (l Logger) Error(msg string) {
	l.errorLog.Fatal(msg)
}