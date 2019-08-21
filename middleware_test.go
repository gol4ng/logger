package logger_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/logger"
)

func TestMiddlewareInterfaceDecorate(t *testing.T) {
	myEntry := logger.Entry{}

	fakeHandler := func(e logger.Entry) error {
		assert.Equal(t, myEntry, e)
		return errors.New("my fake handler error")
	}
	fakeMiddleware := func(h logger.HandlerInterface) logger.HandlerInterface {
		return func(e logger.Entry) error {
			assert.Equal(t, myEntry, e)
			err := h(e)
			assert.Error(t, err)
			return errors.New("middleware prepend " + err.Error())
		}
	}
	middlewares := logger.MiddlewareStack(fakeMiddleware)
	handler := middlewares.Decorate(fakeHandler)

	err := handler(myEntry)
	assert.Error(t, err)
	assert.Equal(t, "middleware prepend my fake handler error", err.Error())
}

func TestDecorateHandler(t *testing.T) {
	myEntry := logger.Entry{}

	fakeHandler := func(e logger.Entry) error {
		assert.Equal(t, myEntry, e)
		return errors.New("my fake handler error")
	}
	fakeMiddleware := func(h logger.HandlerInterface) logger.HandlerInterface {
		return func(e logger.Entry) error {
			assert.Equal(t, myEntry, e)
			err := h(e)
			assert.Error(t, err)
			return errors.New("middleware prepend " + err.Error())
		}
	}
	handler := logger.DecorateHandler(fakeHandler, fakeMiddleware)

	err := handler(myEntry)
	assert.Error(t, err)
	assert.Equal(t, "middleware prepend my fake handler error", err.Error())
}

func TestMiddlewareStack(t *testing.T) {
	fakeMiddleware := func(logger.HandlerInterface) logger.HandlerInterface {
		return func(entry logger.Entry) error {
			return nil
		}
	}

	assert.IsType(t, logger.Middlewares{}, logger.MiddlewareStack(fakeMiddleware))
}
