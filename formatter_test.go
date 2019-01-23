package logger_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/instabledesign/logger"
)

func TestNopFormatter_Format(t *testing.T) {
	tests := []struct {
		name      string
		formatter *logger.NopFormatter
	}{
		{name: "test nil formatter struct", formatter: &logger.NopFormatter{}},
		{name: "test NewNopFormatter()", formatter: logger.NewNopFormatter()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := logger.Entry{}
			assert.Equal(t, "", tt.formatter.Format(e))
		})
	}
}
