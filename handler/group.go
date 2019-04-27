package handler

import (
	"github.com/gol4ng/logger"
)

type Group struct {
	stopOnError bool
	handlers    []logger.HandlerInterface
}

func (g *Group) Handle(entry logger.Entry) error {
	var err error
	for _, h := range g.handlers {
		if err = h.Handle(entry); err == nil {
			continue
		}
		if g.stopOnError {
			return err
		}
		//TODO handle multiple err
	}

	return err
}

func NewGroupBlocking(handlers ...logger.HandlerInterface) *Group {
	return &Group{handlers: handlers, stopOnError: true}
}

func NewGroup(handlers ...logger.HandlerInterface) *Group {
	return &Group{handlers: handlers}
}
