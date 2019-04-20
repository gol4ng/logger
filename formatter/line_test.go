package formatter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
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
			expected: "test message warning <nil>",
		},
		{
			name:     "test simple format 2",
			format:   "%s %d %s",
			expected: "test message 4 <nil>",
		},
		{
			name:     "test simple format",
			format:   "%[2]s %[1]s %[3]s",
			expected: "warning test message <nil>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := formatter.NewLine(tt.format)

			assert.Equal(t, tt.expected, f.Format(logger.Entry{Message: "test message", Level: logger.WarningLevel, Context: nil}))
		})
	}
}
