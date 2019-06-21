package handler

import (
	"github.com/gol4ng/logger"
)

// Chan will send Entry into a given go channel
func Chan(entryChan chan<- logger.Entry) logger.HandlerInterface {
	return func(entry logger.Entry) error {
		entryChan <- entry
		return nil
	}
}
