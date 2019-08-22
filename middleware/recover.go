package middleware

import (
	"fmt"

	"github.com/gol4ng/logger"
)

// PanicError is an error that contains panic data
type PanicError struct {
	message string
	Data    interface{}
}

// Error implements the error interface for PanicError
func (p *PanicError) Error() string {
	return fmt.Sprintf("%s : %v", p.message, p.Data)
}

// NewPanicError builds a specific error from panic data
func NewPanicError(message string, data interface{}) *PanicError {
	return &PanicError{message: message, Data: data}
}

// Recover provide a middleware func that will handle panics that occur in the underlying handler
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
