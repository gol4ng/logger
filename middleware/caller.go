package middleware

import (
	"github.com/gol4ng/logger"
	"runtime"
)

// This handler will add 2 fields in the handled entry context
// eg:
//     _file:/..../file_of_the_handle_caller.go
//     _line:72
func Caller(skip int) logger.Middleware {
	return func(handler logger.HandlerInterface) logger.HandlerInterface {
		return func(entry logger.Entry) error {
			_, file, line, ok := runtime.Caller(skip)

			if ok {
				if entry.Context == nil {
					entry.Context = &logger.Context{}
				}
				entry.Context.
					Add("_file", file).
					Add("_line", line)
			}
			return handler(entry)
		}
	}
}
