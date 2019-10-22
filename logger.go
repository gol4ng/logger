package logger

import (
	"errors"
	"fmt"
	"strings"
)

// Log Severity level
//
// https://github.com/freebsd/freebsd/blob/master/sys/sys/syslog.h#L51
// From /usr/include/sys/syslog.h.
// These are the same on Linux, BSD, and OS X.
const (
	EmergencyLevel Level = iota
	AlertLevel
	CriticalLevel
	ErrorLevel
	WarningLevel
	NoticeLevel
	InfoLevel
	DebugLevel
)

var stringToLevel = map[string]Level{
	"emergency": EmergencyLevel,
	"alert":     AlertLevel,
	"critical":  CriticalLevel,
	"error":     ErrorLevel,
	"warning":   WarningLevel,
	"notice":    NoticeLevel,
	"info":      InfoLevel,
	"debug":     DebugLevel,
}

var levelToString = map[Level]string{
	EmergencyLevel: "emergency",
	AlertLevel:     "alert",
	CriticalLevel:  "critical",
	ErrorLevel:     "error",
	WarningLevel:   "warning",
	NoticeLevel:    "notice",
	InfoLevel:      "info",
	DebugLevel:     "debug",
}

// LevelString represent log Entry level as string
type LevelString string

// Level will return log Level for string or DebugLevel if unknown value
func (l LevelString) Level() Level {
	if v, ok := stringToLevel[strings.ToLower(string(l))]; ok {
		return v
	}
	return DebugLevel
}

// Level represent log Entry levelString
type Level int8

// String will return Level as string
func (l Level) String() string {
	if v, ok := levelToString[l]; ok {
		return v
	}
	return fmt.Sprintf("level(%d)", l)
}

// LogInterface define simplest logger contract
// See LoggerInterface for a more convenient one
type LogInterface interface {
	Log(message string, level Level, context *Context) error
}

// LoggerInterface define a convenient logger contract
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

// WrappableLoggerInterface define a contract to wrap or decorate the logger with given middleware
type WrappableLoggerInterface interface {
	LoggerInterface
	// Implementation should return the same logger after wrapping it
	Wrap(middleware MiddlewareInterface) LoggerInterface
	// Implementation should return a new decorated logger
	WrapNew(middleware MiddlewareInterface) LoggerInterface
}

// Logger is basic implementation of WrappableLoggerInterface
type Logger struct {
	handler HandlerInterface
}

// Debug will log a debug message
func (l *Logger) Debug(message string, context *Context) error {
	return l.Log(message, DebugLevel, context)
}

// Info will log a info message
func (l *Logger) Info(message string, context *Context) error {
	return l.Log(message, InfoLevel, context)
}

// Notice will log a notice message
func (l *Logger) Notice(message string, context *Context) error {
	return l.Log(message, NoticeLevel, context)
}

// Warning will log a warning message
func (l *Logger) Warning(message string, context *Context) error {
	return l.Log(message, WarningLevel, context)
}

// Error will log a error message
func (l *Logger) Error(message string, context *Context) error {
	return l.Log(message, ErrorLevel, context)
}

// Critical will log a critical message
func (l *Logger) Critical(message string, context *Context) error {
	return l.Log(message, CriticalLevel, context)
}

// Alert will log a alert message
func (l *Logger) Alert(message string, context *Context) error {
	return l.Log(message, AlertLevel, context)
}

// Emergency will log a emergency message
func (l *Logger) Emergency(message string, context *Context) error {
	return l.Log(message, EmergencyLevel, context)
}

// Log will log a message with a given level
func (l *Logger) Log(message string, level Level, context *Context) error {
	return l.handler(Entry{message, level, context})
}

// Wrap will return the logger after decorate his handler with given middleware
func (l *Logger) Wrap(middleware MiddlewareInterface) LoggerInterface {
	l.handler = middleware(l.handler)
	return l
}

// WrapNew will return a new logger after wrap his handler with given middleware
func (l *Logger) WrapNew(middleware MiddlewareInterface) LoggerInterface {
	return &Logger{handler: middleware(l.handler)}
}

// NewNopLogger will create a new no operating logger that log nowhere
func NewNopLogger() *Logger {
	return &Logger{handler: NopHandler}
}

// NewLogger will return a new logger
func NewLogger(handler HandlerInterface) *Logger {
	if handler == nil {
		panic(errors.New("handler must not be <nil>"))
	}
	return &Logger{handler: handler}
}
