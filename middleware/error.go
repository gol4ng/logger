package middleware

import (
	"fmt"

	"github.com/gol4ng/logger"
)

// Error provide a middleware that will catch underlying handler errors and print it to stdout
// if passThrough is set to true, the error will be printed and then returned in order to be handled elsewhere (in the overlying handler for example)
func Error(passThrough bool) logger.MiddlewareInterface {
	return func(handler logger.HandlerInterface) logger.HandlerInterface {
		return func(entry logger.Entry) error {
			err := handler(entry)
			if err != nil {
				fmt.Printf("[Error middleware] an error occured : %v\n", err)
				if passThrough {
					return err
				}
			}
			return nil
		}
	}
}
