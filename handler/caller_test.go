package handler_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/handler"
	"github.com/gol4ng/logger/mocks"
)

func TestCaller_Handle(t *testing.T) {
	mockHandler := &mocks.HandlerInterface{}
	mockHandler.On("Handle", mock.AnythingOfType("logger.Entry")).Return(func(entry logger.Entry) error {
		assert.Equal(t, "my_log_message", entry.Message)
		assert.Equal(t, logger.DebugLevel, entry.Level)
		contextStr := entry.Context.String()
		assert.Contains(t, contextStr, "<_file:")
		assert.Contains(t, contextStr, "caller_test.go")
		assert.Contains(t, contextStr, "<_line:34>")

		return nil
	})

	contextHandler := handler.NewCaller(mockHandler, 1)

	logEntry := logger.Entry{
		Message: "my_log_message",
		Level:   logger.DebugLevel,
		Context: nil,
	}
	assert.Nil(t, contextHandler.Handle(logEntry))
}

func TestNewCallerWrapper(t *testing.T) {
	mockHandler := &mocks.HandlerInterface{}
	mockHandler.On("Handle", mock.AnythingOfType("logger.Entry")).Return(func(entry logger.Entry) error {
		assert.Equal(t, "my_log_message", entry.Message)
		assert.Equal(t, logger.DebugLevel, entry.Level)
		contextStr := entry.Context.String()
		assert.Contains(t, contextStr, "<my_key:my_overwrited_value>")
		assert.Contains(t, contextStr, "<my_entry_key:my_entry_value>")
		assert.Contains(t, contextStr, "<_file:")
		assert.Contains(t, contextStr, "caller_test.go")
		assert.Contains(t, contextStr, "<_line:64>")

		return nil
	})

	wrapper := handler.NewCallerWrapper(1)
	contextHandler := wrapper(mockHandler)

	entryCaller := logger.Context(map[string]logger.Field{
		"my_key":       {Value: "my_overwrited_value"},
		"my_entry_key": {Value: "my_entry_value"},
	})
	logEntry := logger.Entry{
		Message: "my_log_message",
		Level:   logger.DebugLevel,
		Context: &entryCaller,
	}
	assert.Nil(t, contextHandler.Handle(logEntry))

	mockHandler.AssertCalled(t, "Handle", logEntry)
}
