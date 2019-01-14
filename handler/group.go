package handler

import (
	"github.com/instabledesign/logger"
)

type Group struct {
	stopOnError bool
	handlers    []logger.HandlerInterface
}

func (g *Group) Handle(e logger.Entry) error {
	var err error
	for _, h := range g.handlers {
		//TODO GO ROUTINE
		if err = h.Handle(e); err == nil {
			continue
		}
		if g.stopOnError {
			return err
		}
		//TODO handle multiple err
	}

	return err
}

func NewGroupBlocking(h []logger.HandlerInterface) *Group {
	return &Group{handlers: h, stopOnError: true}
}

func NewGroup(h []logger.HandlerInterface) *Group {
	return &Group{handlers: h}
}
