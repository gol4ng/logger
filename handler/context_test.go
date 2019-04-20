package handler_test

import (
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/handler"
	"github.com/gol4ng/logger/mocks"
)

func TestContext_Handle(t *testing.T) {
	mockHandler := &mocks.HandlerInterface{}
	mockHandler.On("Handle", mock.AnythingOfType("logger.Entry")).Return(func(e logger.Entry) error {
		assert.Equal(t, "my_log_message", e.Message)
		assert.Equal(t, logger.DebugLevel, e.Level)
		contextStr := e.Context.String()
		assert.Contains(t, contextStr, "my_key:my_overwrited_value")
		assert.Contains(t, contextStr, "my_entry_key:my_entry_value")

		return nil
	})

	defaultContext := logger.Context(map[string]logger.Field{
		"my_key": logger.Field{Value: "my_value"},
	})

	contextHandler := handler.NewContext(mockHandler, &defaultContext)

	entryContext := logger.Context(map[string]logger.Field{
		"my_key":       logger.Field{Value: "my_overwrited_value"},
		"my_entry_key": logger.Field{Value: "my_entry_value"},
	})
	logEntry := logger.Entry{
		Message: "my_log_message",
		Level:   logger.DebugLevel,
		Context: &entryContext,
	}
	assert.Nil(t, contextHandler.Handle(logEntry))

	mockHandler.AssertCalled(t, "Handle", logEntry)
}
