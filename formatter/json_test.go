package formatter_test

import (
	"testing"

	"github.com/instabledesign/logger/formatter"
	"github.com/stretchr/testify/assert"

	"github.com/instabledesign/logger"
)

func TestJson_Format(t *testing.T) {
	tests := []struct {
		name     string
		entry    logger.Entry
		expected string
	}{
		{
			name:     "test simple message",
			entry:    logger.Entry{Message: "test message", Level: logger.DebugLevel, Context: nil},
			expected: "{\"Message\":\"test message\",\"Level\":-1,\"Context\":null}",
		},
		{
			name:     "test simple message",
			entry:    logger.Entry{Message: "test message", Level: logger.WarnLevel, Context: &map[string]interface{}{"ctx-field": "ctx-value"}},
			expected: "{\"Message\":\"test message\",\"Level\":1,\"Context\":{\"ctx-field\":\"ctx-value\"}}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := formatter.NewJson()

			assert.Equal(t, tt.expected, f.Format(tt.entry))
		})
	}
}
