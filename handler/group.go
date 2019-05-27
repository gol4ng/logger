package handler

import (
	"github.com/gol4ng/logger"
)

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
