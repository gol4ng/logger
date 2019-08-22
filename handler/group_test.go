package handler_test

import (
	"errors"
	"os"
	"testing"

	"github.com/gol4ng/logger/formatter"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/handler"
)

func TestGroup_Handle(t *testing.T) {
	entry := logger.Entry{}

	Acalled := false
	mockHandlerA := func(entry logger.Entry) error {
		Acalled = true
		return nil
	}

	Bcalled := false
	mockHandlerB := func(entry logger.Entry) error {
		Bcalled = true
		return nil
	}

	h := handler.Group(mockHandlerA, mockHandlerB)

	assert.Nil(t, h(entry))

	assert.True(t, Acalled)
	assert.True(t, Bcalled)
}

func TestGroup_HandleWithError(t *testing.T) {
	entry := logger.Entry{}

	err := errors.New("my error")
	Acalled := false
	mockHandlerA := func(entry logger.Entry) error {
		Acalled = true
		return err
	}

	Bcalled := false
	mockHandlerB := func(entry logger.Entry) error {
		Bcalled = true
		return nil
	}

	h := handler.Group(mockHandlerA, mockHandlerB)

	assert.Equal(t, err, h(entry))

	assert.True(t, Acalled)
	assert.False(t, Bcalled)
}

// =====================================================================================================================
// ================================================= EXAMPLES ==========================================================
// =====================================================================================================================

func ExampleGroup() {
	lineFormatter := formatter.NewDefaultFormatter()
	lineLogHandler := handler.Stream(os.Stdout, lineFormatter)

	jsonFormatter := formatter.NewJSONEncoder()
	jsonLogHandler := handler.Stream(os.Stdout, jsonFormatter)

	groupHandler := handler.Group(lineLogHandler, jsonLogHandler)

	groupHandler(logger.Entry{Message: "Log example"})

	//Output:
	// <emergency> Log example
	// {"Message":"Log example","Level":0,"Context":null}
}
