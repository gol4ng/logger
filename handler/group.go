package handler

import (
	"github.com/gol4ng/logger"
)

func Group(handlers ...logger.HandlerInterface) logger.HandlerInterface {
	stopOnError := false
	return func(entry logger.Entry) error {
		var err error
		for _, handler := range handlers {
			if err = handler(entry); err == nil {
				continue
			}
			if stopOnError {
				return err
			}
			//TODO handle multiple err
		}
		return err
	}
}
