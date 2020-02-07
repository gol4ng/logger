package logger_test

import (
	"testing"

	"github.com/gol4ng/logger"
	"github.com/stretchr/testify/assert"
)

func TestLevelString_Level(t *testing.T) {
	tests := []struct {
		levelString logger.LevelString
		expected    logger.Level
	}{
		{levelString: "DEBUG", expected: logger.DebugLevel},
		{levelString: "debug", expected: logger.DebugLevel},
		{levelString: "info", expected: logger.InfoLevel},
		{levelString: "notice", expected: logger.NoticeLevel},
		{levelString: "warning", expected: logger.WarningLevel},
		{levelString: "error", expected: logger.ErrorLevel},
		{levelString: "critical", expected: logger.CriticalLevel},
		{levelString: "alert", expected: logger.AlertLevel},
		{levelString: "emergency", expected: logger.EmergencyLevel},
		{levelString: "anothervalue", expected: logger.DebugLevel},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, tt.levelString.Level())
	}
}

func TestLevel_String(t *testing.T) {
	tests := []struct {
		level    logger.Level
		expected string
	}{
		{level: logger.DebugLevel, expected: "debug"},
		{level: logger.InfoLevel, expected: "info"},
		{level: logger.NoticeLevel, expected: "notice"},
		{level: logger.WarningLevel, expected: "warning"},
		{level: logger.ErrorLevel, expected: "error"},
		{level: logger.CriticalLevel, expected: "critical"},
		{level: logger.AlertLevel, expected: "alert"},
		{level: logger.EmergencyLevel, expected: "emergency"},
		{level: 123, expected: "level(123)"},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, tt.level.String())
	}
}
