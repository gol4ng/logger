package middleware

import (
	"runtime"

	"github.com/gol4ng/logger"
)

// Caller will add 2 fields in the handled entry context
// eg:
//     _file:/..../file_of_the_handle_caller.go
//     _line:72
func Caller(skip int) logger.MiddlewareInterface {
	return func(handler logger.HandlerInterface) logger.HandlerInterface {
		return func(entry logger.Entry) error {
			_, file, line, ok := runtime.Caller(skip)

			if ok {
				if entry.Context == nil {
					entry.Context = &logger.Context{}
				}
				entry.Context.
					SetField(logger.String("file", file)).
					Add("line", line)
			}
			return handler(entry)
		}
	}
}
