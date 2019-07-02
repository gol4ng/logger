package middleware_test

import (
	"errors"
	"testing"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/middleware"
	"github.com/stretchr/testify/assert"
)

func TestError_PassThrought(t *testing.T) {
	logEntry := logger.Entry{}
	e := errors.New("my_fake_error")

	mockHandler := func(entry logger.Entry) error {
		assert.Equal(t, logEntry, entry)
		return e
	}

	errorMiddleware := middleware.Error(true)

	assert.Equal(t, e, errorMiddleware(mockHandler)(logEntry))
}

func TestError(t *testing.T) {
	logEntry := logger.Entry{}
	e := errors.New("my_fake_error")

	mockHandler := func(entry logger.Entry) error {
		assert.Equal(t, logEntry, entry)
		return e
	}

	errorMiddleware := middleware.Error(false)

	assert.Nil(t, errorMiddleware(mockHandler)(logEntry))
}

func TestError_WithoutError(t *testing.T) {
	logEntry := logger.Entry{}

	mockHandler := func(entry logger.Entry) error {
		assert.Equal(t, logEntry, entry)
		return nil
	}

	errorMiddleware := middleware.Error(true)

	assert.Nil(t, errorMiddleware(mockHandler)(logEntry))
}
