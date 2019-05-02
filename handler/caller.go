package handler

import (
	"github.com/gol4ng/logger"
	"runtime"
)

// This handler will add 2 fields in the handled entry context
// _file:/..../file_of_the_handle_caller.go
// _line:72
type Caller struct {
	handler logger.HandlerInterface
	skip    int
}

func (c Caller) Handle(entry logger.Entry) error {
	_, file, line, ok := runtime.Caller(c.skip)

	if ok {
		if entry.Context == nil {
			entry.Context = &logger.Context{}
		}
		entry.Context.
			Add("_file", file).
			Add("_line", line)
	}
	return c.handler.Handle(entry)
}

func NewCaller(handler logger.HandlerInterface, skip int) *Caller {
	return &Caller{handler: handler, skip: skip}
}

func NewCallerWrapper(skip int) logger.WrapperHandler {
	return func(handler logger.HandlerInterface) logger.HandlerInterface {
		return NewCaller(handler, skip)
	}
}
