package handler_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/instabledesign/logger"
	"github.com/instabledesign/logger/handler"
	"github.com/instabledesign/logger/mocks"
)

func TestNopStream_Handle(t *testing.T) {
	assert.Nil(t, handler.NewNopStream().Handle(logger.Entry{}))
}

func TestStream_Handle(t *testing.T) {
	var b bytes.Buffer
	mockFormatter := mocks.FormatterInterface{}
	mockFormatter.On("Format", mock.AnythingOfType("logger.Entry")).Return("my formatter return")

	h := handler.NewStream(&b, &mockFormatter)

	assert.Nil(t, h.Handle(logger.Entry{Message: "test message", Level: logger.WarningLevel, Context: nil}))
	assert.Equal(t, "my formatter return\n", b.String())
}

func TestStream_HandleWithError(t *testing.T) {
	err := errors.New("my error")
	mockFormatter := mocks.FormatterInterface{}
	mockFormatter.On("Format", mock.AnythingOfType("logger.Entry")).Return("my formatter return")

	h := handler.NewStream(&WriterError{Error: err}, &mockFormatter)

	assert.Equal(t, err, h.Handle(logger.Entry{}))
}

type WriterError struct {
	Number int
	Error  error
}

func (w *WriterError) Write(p []byte) (n int, err error) {
	return w.Number, w.Error
}
