package middleware_test

import (
	"testing"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/middleware"
	"github.com/stretchr/testify/assert"
)

func TestContext_Handle(t *testing.T) {
	mockHandler := func(entry logger.Entry) error {
		assert.Equal(t, "my_log_message", entry.Message)
		assert.Equal(t, logger.DebugLevel, entry.Level)
		contextStr := entry.Context.String()
		assert.Contains(t, contextStr, "<my_key:my_overwrited_value>")
		assert.Contains(t, contextStr, "<my_entry_key:my_entry_value>")

		return nil
	}

	defaultContext := logger.Context(map[string]logger.Field{
		"my_key": {Value: "my_value"},
	})

	context := middleware.Context(&defaultContext)

	entryContext := logger.Context(map[string]logger.Field{
		"my_key":       {Value: "my_overwrited_value"},
		"my_entry_key": {Value: "my_entry_value"},
	})
	logEntry := logger.Entry{
		Message: "my_log_message",
		Level:   logger.DebugLevel,
		Context: &entryContext,
	}
	assert.Nil(t, context(mockHandler)(logEntry))
}
