package logger

import (
	"fmt"
)

const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

type Level int8

func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case PanicLevel:
		return "panic"
	case FatalLevel:
		return "fatal"
	default:
		return fmt.Sprintf("Level(%d)", l)
	}
}

type LoggerInterface interface {
	Log(msg string, lvl Level, ctx *map[string]interface{})
}

type Logger struct {
	HandlerInterface
}

func (l *Logger) Debug(msg string, ctx *map[string]interface{}) error {
	return l.Log(msg, DebugLevel, ctx)
}
func (l *Logger) Info(msg string, ctx *map[string]interface{}) error {
	return l.Log(msg, InfoLevel, ctx)
}
func (l *Logger) Warn(msg string, ctx *map[string]interface{}) error {
	return l.Log(msg, WarnLevel, ctx)
}
func (l *Logger) Error(msg string, ctx *map[string]interface{}) error {
	return l.Log(msg, ErrorLevel, ctx)
}
func (l *Logger) Panic(msg string, ctx *map[string]interface{}) error {
	return l.Log(msg, PanicLevel, ctx)
}
func (l *Logger) Fatal(msg string, ctx *map[string]interface{}) error {
	return l.Log(msg, FatalLevel, ctx)
}

func (l *Logger) Log(msg string, lvl Level, ctx *map[string]interface{}) error {
	return l.Handle(Entry{msg, lvl, ctx})
}

func NewNilLogger() *Logger {
	return &Logger{HandlerInterface: &NilHandler{}}
}

func NewLogger(h HandlerInterface) *Logger {
	return &Logger{HandlerInterface: h}
}
