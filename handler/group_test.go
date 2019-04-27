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
	entry := logger.Entry{}

	mockHandlerA := mocks.HandlerInterface{}
	mockHandlerA.On("Handle", entry).Return(nil)

	mockHandlerB := mocks.HandlerInterface{}
	mockHandlerB.On("Handle", entry).Return(nil)

	h := handler.NewGroup(&mockHandlerA, &mockHandlerB)

	assert.Nil(t, h.Handle(entry))

	mockHandlerA.AssertCalled(t, "Handle", entry)
	mockHandlerB.AssertCalled(t, "Handle", entry)
}

func TestNewGroupBlocking_HandleWithError(t *testing.T) {
	entry := logger.Entry{}
	err := errors.New("my error")

	mockHandlerA := mocks.HandlerInterface{}
	mockHandlerA.On("Handle", entry).Return(err)

	mockHandlerB := mocks.HandlerInterface{}
	mockHandlerB.On("Handle", entry)

	h := handler.NewGroupBlocking(&mockHandlerA, &mockHandlerB)

	assert.Equal(t, err, h.Handle(entry))

	mockHandlerA.AssertCalled(t, "Handle", entry)
	mockHandlerB.AssertNotCalled(t, "Handle", entry)
}
