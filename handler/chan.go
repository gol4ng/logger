package handler

import (
	"github.com/gol4ng/logger"
)

func Chan(entryChan chan<- logger.Entry) logger.HandlerInterface {
	return func(entry logger.Entry) error {
		entryChan <- entry
		return nil
	}
}
