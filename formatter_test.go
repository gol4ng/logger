package logger_test

import (
	"testing"

	"github.com/instabledesign/logger"

	"github.com/stretchr/testify/assert"
)

func TestNilFormatter_Format(t *testing.T) {
	tests := []struct {
		name      string
		formatter *logger.NilFormatter
	}{
		{
			name:      "test nil formatter struc",
			formatter: &logger.NilFormatter{},
		},
		{
			name:      "test NewNilFormatter()",
			formatter: logger.NewNilFormatter(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := logger.Entry{}
			assert.Equal(t, e, tt.formatter.Format(e))
		})
	}
}
