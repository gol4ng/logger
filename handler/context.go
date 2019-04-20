package handler

import (
	"github.com/gol4ng/logger"
)

type Context struct {
	handler    logger.HandlerInterface
	defaultCtx *logger.Context
}

func (c Context) Handle(e logger.Entry) error {
	newCtx := logger.NewContext()
	//copy original context
	if c.defaultCtx != nil {
		for name, field := range *c.defaultCtx {
			(*newCtx)[name] = field
		}
	}
	//merge original context with given context
	if e.Context != nil {
		for name, field := range *e.Context {
			(*newCtx)[name] = field
		}
	}

	return c.handler.Handle(logger.Entry{
		Message: e.Message,
		Level:   e.Level,
		Context: newCtx,
	})
}

func NewContext(handler logger.HandlerInterface, defaultContext *logger.Context) Context {
	return Context{handler: handler, defaultCtx: defaultContext}
}
