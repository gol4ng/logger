package mocks

import (
	"strings"

	"github.com/gol4ng/logger"
)

// DecorateMockWrapper will decorate a WrappableLoggerInterface with a real Logger+middlewares in order to have middleware properly working
//
// 	myLogger := &mocks.WrappableLoggerInterface{}
//	newWrapperLogger := &mocks.WrappableLoggerInterface{}
//
//	myLogger.On("WrapNew", mock.AnythingOfType("logger.MiddlewareInterface")).Return(mocks.DecorateMockWrapper(newWrapperLogger))
func DecorateMockWrapper(mockWrapperLoggerInterface *WrappableLoggerInterface) func(middlewares ...logger.MiddlewareInterface) logger.LoggerInterface {
	return func(middlewares ...logger.MiddlewareInterface) logger.LoggerInterface {
		h := func(entry logger.Entry) error {
			mockWrapperLoggerInterface.MethodCalled(strings.Title(entry.Level.String()), entry.Message, entry.Context)
			return nil
		}
		for _, middleware := range middlewares {
			h = middleware(h)
		}
		return logger.NewLogger(h)
	}
}
