package handler

import (
	"github.com/gol4ng/logger"
)

// group handler allows you to apply multiple handlers on a logger each time it receives a new log
// it can be seen as a logger.Handler stack
func Group(handlers ...logger.HandlerInterface) logger.HandlerInterface {
	return func(entry logger.Entry) error {
		var err error
		for _, handler := range handlers {
			if err = handler(entry); err != nil {
				return err
			}
		}
		return nil
	}
}
