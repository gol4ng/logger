package logger

import (
	"fmt"
)

const (
	// Severity.

	// https://github.com/freebsd/freebsd/blob/master/sys/sys/syslog.h#L51
	// From /usr/include/sys/syslog.h.
	// These are the same on Linux, BSD, and OS X.
	EmergencyLevel Level = iota
	AlertLevel
	CriticalLevel
	ErrorLevel
	WarningLevel
	NoticeLevel
	InfoLevel
	DebugLevel
)

type Level int8

func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case NoticeLevel:
		return "notice"
	case WarningLevel:
		return "warning"
	case ErrorLevel:
		return "error"
	case CriticalLevel:
		return "critical"
	case AlertLevel:
		return "alert"
	case EmergencyLevel:
		return "emergency"
	default:
		return fmt.Sprintf("Level(%d)", l)
	}
}

type LoggerInterface interface {
	Log(msg string, lvl Level, ctx *map[string]interface{})
}

type Logger struct {
	handler HandlerInterface
}

func (l *Logger) Debug(msg string, ctx *map[string]interface{}) error {
	return l.Log(msg, DebugLevel, ctx)
}
func (l *Logger) Info(msg string, ctx *map[string]interface{}) error {
	return l.Log(msg, InfoLevel, ctx)
}
func (l *Logger) Notice(msg string, ctx *map[string]interface{}) error {
	return l.Log(msg, NoticeLevel, ctx)
}
func (l *Logger) Warning(msg string, ctx *map[string]interface{}) error {
	return l.Log(msg, WarningLevel, ctx)
}
func (l *Logger) Error(msg string, ctx *map[string]interface{}) error {
	return l.Log(msg, ErrorLevel, ctx)
}
func (l *Logger) Critical(msg string, ctx *map[string]interface{}) error {
	return l.Log(msg, CriticalLevel, ctx)
}
func (l *Logger) Alert(msg string, ctx *map[string]interface{}) error {
	return l.Log(msg, AlertLevel, ctx)
}
func (l *Logger) Emergency(msg string, ctx *map[string]interface{}) error {
	return l.Log(msg, EmergencyLevel, ctx)
}

func (l *Logger) Log(msg string, lvl Level, ctx *map[string]interface{}) error {
	return l.handler.Handle(Entry{msg, lvl, ctx})
}

func (l *Logger) Wrap(w func(h HandlerInterface) HandlerInterface) {
	l.handler = w(l.handler)
}

func (l *Logger) WrapNew(w func(h HandlerInterface) HandlerInterface) *Logger {
	return &Logger{handler: w(l.handler)}
}

func NewNopLogger() *Logger {
	return &Logger{handler: &NopHandler{}}
}

func NewLogger(h HandlerInterface) *Logger {
	return &Logger{handler: h}
}
