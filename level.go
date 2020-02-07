package logger

import (
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
