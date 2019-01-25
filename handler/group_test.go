package handler_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/handler"
	"github.com/gol4ng/logger/mocks"
)

func TestGroup_Handle(t *testing.T) {
	logEntry := logger.Entry{}

	mockHandlerA := mocks.HandlerInterface{}
	mockHandlerA.On("Handle", logEntry).Return(nil)

	mockHandlerB := mocks.HandlerInterface{}
	mockHandlerB.On("Handle", logEntry).Return(nil)

	h := handler.NewGroup(&mockHandlerA, &mockHandlerB)

	assert.Nil(t, h.Handle(logEntry))

	mockHandlerA.AssertCalled(t, "Handle", logEntry)
	mockHandlerB.AssertCalled(t, "Handle", logEntry)
}

func TestNewGroupBlocking_HandleWithError(t *testing.T) {
	logEntry := logger.Entry{}
	err := errors.New("my error")

	mockHandlerA := mocks.HandlerInterface{}
	mockHandlerA.On("Handle", logEntry).Return(err)

	mockHandlerB := mocks.HandlerInterface{}
	mockHandlerB.On("Handle", logEntry)

	h := handler.NewGroupBlocking([]logger.HandlerInterface{&mockHandlerA, &mockHandlerB})

	assert.Equal(t, err, h.Handle(logEntry))

	mockHandlerA.AssertCalled(t, "Handle", logEntry)
	mockHandlerB.AssertNotCalled(t, "Handle", logEntry)
}
