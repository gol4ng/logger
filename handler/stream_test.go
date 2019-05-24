package handler_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/handler"
	"github.com/gol4ng/logger/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStream_Handle(t *testing.T) {
	var b bytes.Buffer
	mockFormatter := mocks.FormatterInterface{}
	mockFormatter.On("Format", mock.AnythingOfType("logger.Entry")).Return("my formatter return")

	h := handler.Stream(&b, &mockFormatter)

	assert.Nil(t, h(logger.Entry{Message: "test message", Level: logger.WarningLevel, Context: nil}))
	assert.Equal(t, "my formatter return\n", b.String())
}

func TestStream_HandleWithError(t *testing.T) {
	err := errors.New("my error")
	mockFormatter := mocks.FormatterInterface{}
	mockFormatter.On("Format", mock.AnythingOfType("logger.Entry")).Return("my formatter return")

	h := handler.Stream(&WriterError{Error: err}, &mockFormatter)

	assert.Equal(t, err, h(logger.Entry{}))
}

type WriterError struct {
	Number int
	Error  error
}

func (w *WriterError) Write(p []byte) (n int, err error) {
	return w.Number, w.Error
}
