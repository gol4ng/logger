package handler

import (
	"github.com/gol4ng/logger"
)

type Chan struct {
	entryChan chan<- logger.Entry
}

func (c Chan) Handle(entry logger.Entry) error {
	c.entryChan <- entry
	return nil
}

func NewChan(entryChan chan<- logger.Entry) Chan {
	return Chan{entryChan: entryChan}
}
