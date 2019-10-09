package middleware

import (
	"strings"

	"github.com/gol4ng/logger"
)

// Placeholder will replace message placeholder with context field
func Placeholder() logger.MiddlewareInterface {
	return func(handler logger.HandlerInterface) logger.HandlerInterface {
		return func(entry logger.Entry) error {
			if entry.Context != nil {
				msg := entry.Message
				for n, f := range *entry.Context {
					msg = strings.Replace(msg, "%"+n+"%", f.String(), -1)
					if !strings.Contains(msg, "%") {
						break
					}
				}
				entry.Message = msg
			}
			return handler(entry)
		}
	}
}
