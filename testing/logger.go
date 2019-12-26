package testing

import (
	"sync"

	"github.com/gol4ng/logger"
)

// Logger can be use in third party in order to do assertion on logged entry
type Logger struct {
	mu      sync.Mutex
	entries []logger.Entry
}

// CleanEntries will reset the in memory entries list
func (l *Logger) CleanEntries() {
	defer l.lock()()
	l.entries = []logger.Entry{}
}

// GetEntries will return the in memory entries list
func (l *Logger) GetEntries() []logger.Entry {
	return l.entries
}

// GetAndCleanEntries will return and clean the in memory entries list
func (l *Logger) GetAndCleanEntries() []logger.Entry {
	defer l.lock()()
	entries := l.GetEntries()
	l.CleanEntries()
	return entries
}

// Debug will log a debug message
func (l *Logger) Debug(message string, context *logger.Context) error {
	return l.Log(message, logger.DebugLevel, context)
}

// Info will log a info message
func (l *Logger) Info(message string, context *logger.Context) error {
	return l.Log(message, logger.InfoLevel, context)
}

// Notice will log a notice message
func (l *Logger) Notice(message string, context *logger.Context) error {
	return l.Log(message, logger.NoticeLevel, context)
}

// Warning will log a warning message
func (l *Logger) Warning(message string, context *logger.Context) error {
	return l.Log(message, logger.WarningLevel, context)
}

// Error will log a error message
func (l *Logger) Error(message string, context *logger.Context) error {
	return l.Log(message, logger.ErrorLevel, context)
}

// Critical will log a critical message
func (l *Logger) Critical(message string, context *logger.Context) error {
	return l.Log(message, logger.CriticalLevel, context)
}

// Alert will log a alert message
func (l *Logger) Alert(message string, context *logger.Context) error {
	return l.Log(message, logger.AlertLevel, context)
}

// Emergency will log a emergency message
func (l *Logger) Emergency(message string, context *logger.Context) error {
	return l.Log(message, logger.EmergencyLevel, context)
}

// Log will log a message with a given level
func (l *Logger) Log(message string, level logger.Level, context *logger.Context) error {
	defer l.lock()()
	l.entries = append(l.entries, logger.Entry{
		Message: message,
		Level:   level,
		Context: context,
	})
	return nil
}

func (l *Logger) lock() func() {
	l.mu.Lock()
	return func() { l.mu.Unlock() }
}
