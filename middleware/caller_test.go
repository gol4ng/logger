package middleware_test

import (
	"github.com/gol4ng/logger/middleware"
	"testing"

	"github.com/gol4ng/logger"
	"github.com/stretchr/testify/assert"
)

func TestCaller_Handle(t *testing.T) {
	mockHandler := func(entry logger.Entry) error {
		assert.Equal(t, "my_log_message", entry.Message)
		assert.Equal(t, logger.DebugLevel, entry.Level)
		contextStr := entry.Context.String()
		assert.Contains(t, contextStr, "<_file:")
		assert.Contains(t, contextStr, "caller_test.go")
		assert.Contains(t, contextStr, "<_line:30>")

		return nil
	}

	caller := middleware.Caller(1)

	logEntry := logger.Entry{
		Message: "my_log_message",
		Level:   logger.DebugLevel,
		Context: nil,
	}
	assert.Nil(t, caller(mockHandler)(logEntry))
}
