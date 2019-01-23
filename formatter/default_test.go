package formatter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/instabledesign/logger"
	"github.com/instabledesign/logger/formatter"
)

func TestDefaultFormatter_Format(t *testing.T) {
	tests := []struct {
		name      string
		formatter *formatter.DefaultFormatter
		entry     logger.Entry
		expected  string
	}{
		{name: "test default formatter struct", formatter: &formatter.DefaultFormatter{}, entry: logger.Entry{}, expected: "emergency "},
		{name: "test NewNilFormatter()", formatter: formatter.NewDefaultFormatter(), entry: logger.Entry{}, expected: "emergency "},
		{name: "test NewNilFormatter()", formatter: formatter.NewDefaultFormatter(), entry: logger.Entry{Message: "my message", Level: logger.DebugLevel}, expected: "debug my message"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.formatter.Format(tt.entry))
		})
	}
}
