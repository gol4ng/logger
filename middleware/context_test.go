package middleware_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/gol4ng/logger/handler"
	"github.com/gol4ng/logger/middleware"
)

func TestContext_Handle(t *testing.T) {
	called := false
	mockHandler := func(entry logger.Entry) error {
		called = true
		assert.Equal(t, "my_log_message", entry.Message)
		assert.Equal(t, logger.DebugLevel, entry.Level)
		contextStr := entry.Context.String()
		assert.Contains(t, contextStr, "<my_key:my_overwritten_value>")
		assert.Contains(t, contextStr, "<my_entry_key:my_entry_value>")
		assert.Contains(t, contextStr, "<my_default_key:my_default_value>")

		return nil
	}

	defaultContext := logger.Context(map[string]logger.Field{
		"my_key":         {Value: "my_value"},
		"my_default_key": {Value: "my_default_value"},
	})

	context := middleware.Context(&defaultContext)

	entryContext := logger.Context(map[string]logger.Field{
		"my_key":       {Value: "my_overwritten_value"},
		"my_entry_key": {Value: "my_entry_value"},
	})
	logEntry := logger.Entry{
		Message: "my_log_message",
		Level:   logger.DebugLevel,
		Context: &entryContext,
	}
	assert.Nil(t, context(mockHandler)(logEntry))
	assert.True(t, called)
}

// =====================================================================================================================
// ================================================= EXAMPLES ==========================================================
// =====================================================================================================================

func ExampleContext() {
	streamHandler := handler.Stream(os.Stdout, formatter.NewJSONEncoder())
	contextHandler := middleware.Context(logger.Ctx("my_value_1", "value 1"))

	myLogger := logger.NewLogger(handler.Group(contextHandler(streamHandler), streamHandler))
	myLogger.Debug("will be printed", logger.Any("my_value_1", "overwritten value 1"))
	myLogger.Debug("only context handler values will be printed")

	//Output:
	//{"Message":"will be printed","Level":7,"Context":{"my_value_1":"overwritten value 1"}}
	//{"Message":"will be printed","Level":7,"Context":{"my_value_1":"overwritten value 1"}}
	//{"Message":"only context handler values will be printed","Level":7,"Context":{"my_value_1":"value 1"}}
	//{"Message":"only context handler values will be printed","Level":7,"Context":null}
}
