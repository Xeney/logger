package logger

import (
	"log"
	"os"
)

// Уровни логирования
const (
	LevelDebug = "DEBUG"
	LevelInfo  = "INFO"
	LevelWarn  = "WARN"
	LevelError = "ERROR"
)

// Logger представляет собой простой логгер
type Logger struct {
	*log.Logger
}

// New создает новый экземпляр логгера
func New() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
	}
}

// Debug логирует сообщение на уровне Debug
func (l *Logger) Debug(message string) {
	l.printLog(LevelDebug, message)
}

// Info логирует сообщение на уровне Info
func (l *Logger) Info(message string) {
	l.printLog(LevelInfo, message)
}

// Warn логирует сообщение на уровне Warn
func (l *Logger) Warn(message string) {
	l.printLog(LevelWarn, message)
}

// Error логирует сообщение на уровне Error
func (l *Logger) Error(message interface{}) {
	l.printLog(LevelError, message)
}

// printLog выводит сообщение в лог
func (l *Logger) printLog(level string, message interface{}) {
	l.Printf("[%s] %s\n", level, message)
}
