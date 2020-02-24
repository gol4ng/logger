package logger

import (
	"errors"
	"log"
)

// ErrorHandler will print error and entry when logging error occurred
func ErrorHandler(error error, entry Entry) {
	log.Println(error, entry)
}

// LogInterface define simplest logger contract
// See LoggerInterface for a more convenient one
type LogInterface interface {
	Log(message string, level Level, field ...Field)
}

// LoggerInterface define a convenient logger contract
type LoggerInterface interface {
	LogInterface
	Debug(message string, field ...Field)
	Info(message string, field ...Field)
	Notice(message string, field ...Field)
	Warning(message string, field ...Field)
	Error(message string, field ...Field)
	Critical(message string, field ...Field)
	Alert(message string, field ...Field)
	Emergency(message string, field ...Field)
}

// WrappableLoggerInterface define a contract to wrap or decorate the logger with given middleware
type WrappableLoggerInterface interface {
	LoggerInterface
	// Implementation should return the same logger after wrapping it
	Wrap(middlewares ...MiddlewareInterface) LoggerInterface
	// Implementation should return a new decorated logger
	WrapNew(middlewares ...MiddlewareInterface) LoggerInterface
}

// Logger is basic implementation of WrappableLoggerInterface
type Logger struct {
	handler      HandlerInterface
	ErrorHandler func(error, Entry)
}

// onError will handle when error occur during log
func (l *Logger) onError(error error, entry Entry) {
	l.ErrorHandler(error, entry)
}

// Debug will log a debug message
func (l *Logger) Debug(message string, fields ...Field) {
	l.log(message, DebugLevel, fields...)
}

// Info will log a info message
func (l *Logger) Info(message string, fields ...Field) {
	l.log(message, InfoLevel, fields...)
}

// Notice will log a notice message
func (l *Logger) Notice(message string, fields ...Field) {
	l.log(message, NoticeLevel, fields...)
}

// Warning will log a warning message
func (l *Logger) Warning(message string, fields ...Field) {
	l.log(message, WarningLevel, fields...)
}

// Error will log a message
func (l *Logger) Error(message string, fields ...Field) {
	l.log(message, ErrorLevel, fields...)
}

// Critical will log a critical message
func (l *Logger) Critical(message string, fields ...Field) {
	l.log(message, CriticalLevel, fields...)
}

// Alert will log a alert message
func (l *Logger) Alert(message string, fields ...Field) {
	l.log(message, AlertLevel, fields...)
}

// Emergency will log a emergency message
func (l *Logger) Emergency(message string, fields ...Field) {
	l.log(message, EmergencyLevel, fields...)
}

// Log will log a message with a given level
func (l *Logger) Log(message string, level Level, fields ...Field) {
	l.log(message, level, fields...)
}

func (l *Logger) log(message string, level Level, fields ...Field) {
	entry := Entry{message, level, NewContext(fields...)}
	if err := l.handler(entry); err != nil {
		l.onError(err, entry)
	}
}

// Wrap will return the logger after decorate his handler with given middleware
func (l *Logger) Wrap(middlewares ...MiddlewareInterface) LoggerInterface {
	for _, middleware := range middlewares {
		l.handler = middleware(l.handler)
	}
	return l
}

// WrapNew will return a new logger after wrap his handler with given middleware
func (l *Logger) WrapNew(middlewares ...MiddlewareInterface) LoggerInterface {
	handler := l.handler
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return &Logger{handler: handler}
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
	return &Logger{handler: handler, ErrorHandler: ErrorHandler}
}
