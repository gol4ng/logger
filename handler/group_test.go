package handler_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/instabledesign/logger"
	"github.com/instabledesign/logger/handler"
	"github.com/instabledesign/logger/mocks"
)

func TestGroup_Handle(t *testing.T) {
	logEntry := logger.Entry{}

	mockHandlerA := mocks.HandlerInterface{}
	mockHandlerA.On("Handle", logEntry).Return(nil)

	mockHandlerB := mocks.HandlerInterface{}
	mockHandlerB.On("Handle", logEntry).Return(nil)

	f := handler.NewGroup([]logger.HandlerInterface{&mockHandlerA, &mockHandlerB})

	assert.Nil(t, f.Handle(logEntry))

	mockHandlerA.AssertCalled(t, "Handle", logEntry)
	mockHandlerB.AssertCalled(t, "Handle", logEntry)
}
