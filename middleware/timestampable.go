package middleware

import (
	"time"

	"github.com/gol4ng/logger"
)

func Timestampable() logger.MiddlewareInterface {
	return func(handler logger.HandlerInterface) logger.HandlerInterface {
		return func(entry logger.Entry) error {
			if entry.Context == nil {
				entry.Context = &logger.Context{}
			}
			entry.Context.Set("timestamp", logger.Time(time.Now()))
			return handler(entry)
		}
	}
}
