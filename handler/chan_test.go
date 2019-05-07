package handler_test

import (
	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/handler"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChan_Handle(t *testing.T) {
	entryChan := make(chan logger.Entry)
	chanHandler := handler.Chan(entryChan)

	entryContext := logger.Context(map[string]logger.Field{
		"my_key":       {Value: "my_overwrited_value"},
		"my_entry_key": {Value: "my_entry_value"},
	})
	logEntry := logger.Entry{
		Message: "my_log_message",
		Level:   logger.DebugLevel,
		Context: &entryContext,
	}

	go func() {
		assert.Nil(t, chanHandler(logEntry))
	}()

	assert.Equal(t, logEntry, <-entryChan)
}
