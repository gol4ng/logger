package middleware_test

import (
	"testing"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/middleware"
	"github.com/stretchr/testify/assert"
)

func TestTimestampable_Handle(t *testing.T) {
	mockHandler := func(entry logger.Entry) error {
		assert.Equal(t, "my_log_message", entry.Message)
		assert.Equal(t, logger.DebugLevel, entry.Level)
		assert.True(t, entry.Context.Has("timestamp"))
		contextStr := entry.Context.String()
		assert.Contains(t, contextStr, "<timestamp:")

		return nil
	}

	caller := middleware.Timestamp()

	logEntry := logger.Entry{
		Message: "my_log_message",
		Level:   logger.DebugLevel,
		Context: nil,
	}
	assert.Nil(t, caller(mockHandler)(logEntry))
}
