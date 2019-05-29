package middleware

import (
	"fmt"

	"github.com/gol4ng/logger"
)

type PanicError struct {
	message string
	Data    interface{}
}

func (p *PanicError) Error() string {
	return fmt.Sprintf("%s : %v", p.message, p.Data)
}

func NewPanicError(message string, data interface{}) *PanicError {
	return &PanicError{message: message, Data: data}
}

func Recover() logger.MiddlewareInterface {
	return func(handler logger.HandlerInterface) logger.HandlerInterface {
		return func(entry logger.Entry) (e error) {
			defer func() {
				if errData := recover(); errData != nil {
					e = NewPanicError("[Recover middleware] recover panic", errData)
				}
			}()
			e = handler(entry)
			return e
		}
	}
}
