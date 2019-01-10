package handler

import (
	"github.com/instabledesign/logger"
)

type Group struct {
	handlers []logger.HandlerInterface
}

func (g *Group) Handle(e logger.Entry) error {
	var err error
	for _, h := range g.handlers {
		//TODO GO ROUTINE
		err = h.Handle(e)
	}

	return err
}

func NewGroup(h []logger.HandlerInterface) *Group {
	return &Group{handlers: h}
}
