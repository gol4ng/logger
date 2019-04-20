package handler

import (
	"github.com/gol4ng/logger"
)

type Context struct {
	handler    logger.HandlerInterface
	defaultCtx *logger.Context
}

func NewContext(handler logger.HandlerInterface, defaultContext *logger.Context) Context {
	return Context{handler: handler, defaultCtx: defaultContext}
}

func (c Context) Handle(e logger.Entry) error {
	newCtx := logger.NewContext()
	//copy original context
	if c.defaultCtx != nil {
		for k, v := range *c.defaultCtx {
			(*newCtx)[k] = v
		}
	}
	//merge original context with given context
	if e.Context != nil {
		for k, v := range *e.Context {
			(*newCtx)[k] = v
		}
	}

	return c.handler.Handle(logger.Entry{
		Context: newCtx,
		Level:   e.Level,
		Message: e.Message,
	})
}
