package handler

import (
	"github.com/gol4ng/logger"
)

// channel handler, sends a logger entry into a given go channel
func Chan(entryChan chan<- logger.Entry) logger.HandlerInterface {
	return func(entry logger.Entry) error {
		entryChan <- entry
		return nil
	}
}
