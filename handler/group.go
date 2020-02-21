package handler

import (
	"github.com/gol4ng/logger"
)

// Group will send Entry to each underlying handlers
// useful when you want to send your log in multiple destination eg stdOut/file/logserver
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
