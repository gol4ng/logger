package middleware

import (
	"fmt"

	"github.com/gol4ng/logger"
)

func Error(passThrough bool) logger.MiddlewareInterface {
	return func(handler logger.HandlerInterface) logger.HandlerInterface {
		return func(entry logger.Entry) error {
			err := handler(entry)
			if err != nil {
				fmt.Printf("[Error middleware] an error occured : %v\n", err)
				if passThrough == true {
					return err
				}
			}
			return nil
		}
	}
}
