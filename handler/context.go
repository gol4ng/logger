package handler

import (
	"github.com/gol4ng/logger"
)

type Context struct {
	h          logger.HandlerInterface
	defaultCtx map[string]interface{}
}

func NewContext(h logger.HandlerInterface, ctxToMerge map[string]interface{}) Context {
	return Context{h, ctxToMerge}
}

func (c Context) Handle(e logger.Entry) error {
	newCtx := map[string]interface{}{}
	//copy original context
	for k, v := range *e.Context {
		newCtx[k] = v
	}
	//merge original context with given context
	for k, v := range c.defaultCtx {
		newCtx[k] = v
	}

	return c.h.Handle(logger.Entry{
		Context: &newCtx,
		Level:   e.Level,
		Message: e.Message,
	})
}
