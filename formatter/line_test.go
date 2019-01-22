package formatter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/instabledesign/logger"
	"github.com/instabledesign/logger/formatter"
)

func TestLine_Format(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		expected string
	}{
		{
			name:     "test simple format 1",
			format:   "%s %s %s",
			expected: "test message warning %!s(*map[string]interface {}=<nil>)",
		},
		{
			name:     "test simple format 2",
			format:   "%s %d %s",
			expected: "test message 4 %!s(*map[string]interface {}=<nil>)",
		},
		{
			name:     "test simple format",
			format:   "%[2]s %[1]s %[3]s",
			expected: "warning test message %!s(*map[string]interface {}=<nil>)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := formatter.NewLine(tt.format)

			assert.Equal(t, tt.expected, f.Format(logger.Entry{Message: "test message", Level: logger.WarningLevel, Context: nil}))
		})
	}
}
