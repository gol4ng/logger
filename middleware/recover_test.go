package middleware_test

import (
	"testing"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/middleware"
	"github.com/stretchr/testify/assert"
)

func TestRecover(t *testing.T) {
	logEntry := logger.Entry{}

	mockHandler := func(entry logger.Entry) error {
		assert.Equal(t, logEntry, entry)
		panic("my fake panic message")
		return nil
	}

	errorMiddleware := middleware.Recover()

	err := errorMiddleware(mockHandler)(logEntry)
	assert.IsType(t, &middleware.PanicError{}, err)
	assert.Equal(t, "[Recover middleware] recover panic : my fake panic message", err.Error())
	assert.Equal(t, "my fake panic message", err.(*middleware.PanicError).Data)
}
