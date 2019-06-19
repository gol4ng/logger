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

// transform a logger.Level (int) value into a human readable string value
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

// this interface is the simplest entry point you can rely on in your app
// if you want a more readable api, you can rely on LoggerInterface
type LogInterface interface {
	Log(message string, level Level, context *Context) error
}

// more readable interface that gives you helper functions in order not to pass the log level in the func param
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

// expose a Wrap function that allows you to wrap a logger with a middleware
type WrappableLoggerInterface interface {
	LoggerInterface
	Wrap(middleware MiddlewareInterface) LoggerInterface
	WrapNew(middleware MiddlewareInterface) LoggerInterface
}

// basic implementation of LogInterface, LoggerInterface and
type Logger struct {
	handler HandlerInterface
}

// helper to log a debug message
func (l *Logger) Debug(message string, context *Context) error {
	return l.Log(message, DebugLevel, context)
}

// helper to log a info message
func (l *Logger) Info(message string, context *Context) error {
	return l.Log(message, InfoLevel, context)
}

// helper to log a notice message
func (l *Logger) Notice(message string, context *Context) error {
	return l.Log(message, NoticeLevel, context)
}

// helper to log a warning message
func (l *Logger) Warning(message string, context *Context) error {
	return l.Log(message, WarningLevel, context)
}

// helper to log a error message
func (l *Logger) Error(message string, context *Context) error {
	return l.Log(message, ErrorLevel, context)
}

// helper to log a critical message
func (l *Logger) Critical(message string, context *Context) error {
	return l.Log(message, CriticalLevel, context)
}

// helper to log a alert message
func (l *Logger) Alert(message string, context *Context) error {
	return l.Log(message, AlertLevel, context)
}

// helper to log a emergency message
func (l *Logger) Emergency(message string, context *Context) error {
	return l.Log(message, EmergencyLevel, context)
}

// log a message with a given level
func (l *Logger) Log(message string, level Level, context *Context) error {
	return l.handler(Entry{message, level, context})
}

// wrap a logger with a middleware
func (l *Logger) Wrap(middleware MiddlewareInterface) LoggerInterface {
	l.handler = middleware(l.handler)
	return l
}

// immutable wrap func
func (l *Logger) WrapNew(middleware MiddlewareInterface) LoggerInterface {
	return &Logger{handler: middleware(l.handler)}
}

// create a logger that logs nowhere
func NewNopLogger() *Logger {
	return &Logger{handler: NopHandler}
}

// basic logger constructor
func NewLogger(handler HandlerInterface) *Logger {
	return &Logger{handler: handler}
}
