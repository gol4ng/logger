package middleware

import (
	"github.com/gol4ng/logger"
)

// Context will define and merge defaultContext with each Entry Context
// if a default context key is redefined when calling Log/Debug/Info/... func, it will override the default ctx value
func Context(defaultContext *logger.Context) logger.MiddlewareInterface {
	return func(handler logger.HandlerInterface) logger.HandlerInterface {
		return func(entry logger.Entry) error {
			newCtx := logger.NewContext()
			// copy original context
			if defaultContext != nil {
				for name, field := range *defaultContext {
					(*newCtx)[name] = field
				}
			}
			// merge original context with given context
			if entry.Context != nil {
				for name, field := range *entry.Context {
					(*newCtx)[name] = field
				}
			}
			return handler(logger.Entry{
				Message: entry.Message,
				Level:   entry.Level,
				Context: newCtx,
			})
		}
	}
}
