package middleware

import (
	"time"

	"github.com/gol4ng/logger"
)

// Timestamp middleware will add the timestamp to the log context
func Timestamp() logger.MiddlewareInterface {
	return func(handler logger.HandlerInterface) logger.HandlerInterface {
		return func(entry logger.Entry) error {
			if entry.Context == nil {
				entry.Context = &logger.Context{}
			}
			entry.Context.SetField(logger.Time("timestamp", time.Now()))
			return handler(entry)
		}
	}
}
