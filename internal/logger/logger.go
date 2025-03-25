package logger

import (
	"log/slog"
	"os"
)

type Logger struct {
	Log *slog.Logger
	// errorLog *log.Logger
}

func New() *Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	// infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return &Logger{Log: logger}
}

func (l Logger) Info(msg string) {
	l.Log.Info(msg)
}

func (l Logger) Error(msg string) {
	l.Log.Error(msg)
}
