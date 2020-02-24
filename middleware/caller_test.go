package middleware_test

import (
	"testing"

	"github.com/gol4ng/logger/handler"
	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/middleware"
)

func TestCaller_Handle(t *testing.T) {
	mockHandler := func(entry logger.Entry) error {
		assert.Equal(t, "my_log_message", entry.Message)
		assert.Equal(t, logger.DebugLevel, entry.Level)
		contextStr := entry.Context.String()
		assert.Contains(t, contextStr, "<file:")
		assert.Contains(t, contextStr, "caller_test.go")
		assert.Contains(t, contextStr, "<line:32>")

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

func TestCaller_Handle_SameCallerDepth(t *testing.T) {
	caller := middleware.Caller(3)
	m := handler.NewMemory()
	l := logger.NewLogger(caller(m.Handle))

	l.Debug("my_fake_debug_message")
	l.Log("my_fake_debug_message", logger.DebugLevel)

	entries := m.GetEntries()
	assert.Len(t, entries, 2)

	entry1 := entries[0]
	assert.Contains(t, (*entry1.Context)["file"].Value, "caller_test.go")

	entry2 := entries[1]
	assert.Contains(t, (*entry2.Context)["file"].Value, "caller_test.go")

	assert.Equal(t, (*entry1.Context)["file"].Value, (*entry2.Context)["file"].Value, "This 2 logs must have the same context file value")
}
