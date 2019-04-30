package handler_test

import (
	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/handler"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChan_Handle(t *testing.T) {
	entryChan := make(chan logger.Entry)
	chanHandler := handler.NewChan(entryChan)

	entryContext := logger.Context(map[string]logger.Field{
		"my_key":       logger.Field{Value: "my_overwrited_value"},
		"my_entry_key": logger.Field{Value: "my_entry_value"},
	})
	logEntry := logger.Entry{
		Message: "my_log_message",
		Level:   logger.DebugLevel,
		Context: &entryContext,
	}

	go func() {
		assert.Nil(t, chanHandler.Handle(logEntry))
	}()

	assert.Equal(t, logEntry, <-entryChan)
}
