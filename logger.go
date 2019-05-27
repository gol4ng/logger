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
		return fmt.Sprintf("level(%d)", l)
	}
}

type LogInterface interface {
	Log(message string, level Level, context *Context) error
}

type LoggerInterface interface {
	LogInterface
	Debug(message string, context *Context) error
	Info(message string, context *Context) error
	Notice(message string, context *Context) error
	Warning(message string, context *Context) error
	Error(message string, context *Context) error
	Critical(message string, context *Context) error
	Alert(message string, context *Context) error
	Emergency(message string, context *Context) error
}

type WrappableLoggerInterface interface {
	LoggerInterface
	Wrap(wrapper MiddlewareInterface) LoggerInterface
	WrapNew(wrapper MiddlewareInterface) LoggerInterface
}

type Logger struct {
	handler HandlerInterface
}

func (l *Logger) Debug(message string, context *Context) error {
	return l.Log(message, DebugLevel, context)
}
func (l *Logger) Info(message string, context *Context) error {
	return l.Log(message, InfoLevel, context)
}
func (l *Logger) Notice(message string, context *Context) error {
	return l.Log(message, NoticeLevel, context)
}
func (l *Logger) Warning(message string, context *Context) error {
	return l.Log(message, WarningLevel, context)
}
func (l *Logger) Error(message string, context *Context) error {
	return l.Log(message, ErrorLevel, context)
}
func (l *Logger) Critical(message string, context *Context) error {
	return l.Log(message, CriticalLevel, context)
}
func (l *Logger) Alert(message string, context *Context) error {
	return l.Log(message, AlertLevel, context)
}
func (l *Logger) Emergency(message string, context *Context) error {
	return l.Log(message, EmergencyLevel, context)
}

func (l *Logger) Log(message string, level Level, context *Context) error {
	return l.handler(Entry{message, level, context})
}

func (l *Logger) Wrap(wrapper MiddlewareInterface) LoggerInterface {
	l.handler = wrapper(l.handler)
	return l
}

func (l *Logger) WrapNew(wrapper MiddlewareInterface) LoggerInterface {
	return &Logger{handler: wrapper(l.handler)}
}

func NewNopLogger() WrappableLoggerInterface {
	return &Logger{handler: NopHandler}
}

func NewLogger(handler HandlerInterface) WrappableLoggerInterface {
	return &Logger{handler: handler}
}
