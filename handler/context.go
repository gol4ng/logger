package handler

import (
	"github.com/gol4ng/logger"
)

type Context struct {
	handler    logger.HandlerInterface
	defaultCtx *logger.Context
}

func (c Context) Handle(entry logger.Entry) error {
	newCtx := logger.NewContext()
	//copy original context
	if c.defaultCtx != nil {
		for name, field := range *c.defaultCtx {
			(*newCtx)[name] = field
		}
	}
	//merge original context with given context
	if entry.Context != nil {
		for name, field := range *entry.Context {
			(*newCtx)[name] = field
		}
	}

	return c.handler.Handle(logger.Entry{
		Message: entry.Message,
		Level:   entry.Level,
		Context: newCtx,
	})
}

func NewContext(handler logger.HandlerInterface, defaultContext *logger.Context) Context {
	return Context{handler: handler, defaultCtx: defaultContext}
}

func NewContextWrapper(defaultContext *logger.Context) logger.WrapperHandler {
	return func(handler logger.HandlerInterface) logger.HandlerInterface {
		return NewContext(handler, defaultContext)
	}
}
