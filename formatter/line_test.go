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
			name:     "test simple format",
			format:   "%s %s",
			expected: "test message warn",
		},
		{
			name:     "test simple format",
			format:   "%s %d",
			expected: "test message 1",
		},
		{
			name:     "test simple format",
			format:   "%[2]s %[1]s",
			expected: "warn test message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := formatter.NewLine(tt.format)

			assert.Equal(t, tt.expected, f.Format(logger.Entry{Message: "test message", Level: logger.WarnLevel, Context: nil}))
		})
	}
}
