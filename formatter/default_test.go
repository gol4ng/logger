package formatter_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
)

func TestDefaultFormatter_Format(t *testing.T) {
	tests := []struct {
		name      string
		formatter *formatter.DefaultFormatter
		entry     logger.Entry
		expected  string
	}{
		{name: "test default formatter struct", formatter: &formatter.DefaultFormatter{}, entry: logger.Entry{}, expected: "<emergency>"},
		{name: "test NewDefaultFormatter()", formatter: formatter.NewDefaultFormatter(), entry: logger.Entry{}, expected: "<emergency>"},
		{
			name:      "test NewDefaultFormatter(formatter.WithContext(true)",
			formatter: formatter.NewDefaultFormatter(formatter.WithContext(true)),
			entry:     logger.Entry{Message: "my message", Level: logger.DebugLevel, Context: logger.Ctx("my_key", "my_value")},
			expected:  "<debug> my message {\"my_key\":\"my_value\"}",
		},
		{
			name:      "test formatter.NewDefaultFormatter(formatter.WithColor(true))",
			formatter: formatter.NewDefaultFormatter(formatter.WithColor(true)),
			entry:     logger.Entry{Message: "my message", Level: logger.DebugLevel, Context: logger.Ctx("my_key", "my_value")},
			expected:  "\x1b[1;36m<debug>\x1b[m my message",
		},
		{
			name:      "test formatter.NewDefaultFormatter(formatter.WithColor(true), formatter.WithContext(true)",
			formatter: formatter.NewDefaultFormatter(formatter.WithColor(true), formatter.WithContext(true)),
			entry:     logger.Entry{Message: "my message", Level: logger.DebugLevel, Context: logger.Ctx("my_key", "my_value")},
			expected:  "\x1b[1;36m<debug>\x1b[m my message {\"my_key\":\"my_value\"}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.formatter.Format(tt.entry))
		})
	}
}

func TestDefaultFormatter_Format_AllColor(t *testing.T) {
	tests := []struct {
		level    logger.Level
		expected string
	}{
		{level: logger.EmergencyLevel, expected: "\x1b[1;37;41m<emergency>\x1b[m my message"},
		{level: logger.AlertLevel, expected: "\x1b[1;30;43m<alert>\x1b[m my message"},
		{level: logger.CriticalLevel, expected: "\x1b[1;30;47m<critical>\x1b[m my message"},
		{level: logger.ErrorLevel, expected: "\x1b[1;31m<error>\x1b[m my message"},
		{level: logger.WarningLevel, expected: "\x1b[1;33m<warning>\x1b[m my message"},
		{level: logger.NoticeLevel, expected: "\x1b[1;34m<notice>\x1b[m my message"},
		{level: logger.InfoLevel, expected: "\x1b[1;32m<info>\x1b[m my message"},
		{level: logger.DebugLevel, expected: "\x1b[1;36m<debug>\x1b[m my message"},
	}
	defaultFormatter := formatter.NewDefaultFormatter(formatter.WithColor(true))
	for _, tt := range tests {
		t.Run(tt.level.String(), func(t *testing.T) {
			assert.Equal(t, tt.expected, defaultFormatter.Format(logger.Entry{Level: tt.level, Message: "my message"}))
		})
	}
}

// =====================================================================================================================
// ================================================= EXAMPLES ==========================================================
// =====================================================================================================================

func ExampleDefaultFormatter() {
	defaultFormatter := formatter.NewDefaultFormatter(formatter.WithContext(true))

	fmt.Println(defaultFormatter.Format(
		logger.Entry{
			Message: "My log message",
			Level:   logger.InfoLevel,
			Context: logger.NewContext().Add("my_key", "my_value"),
		},
	))

	//Output:
	//<info> My log message {"my_key":"my_value"}
}
