package testing

import (
	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/handler"
)

// NewLogger will return a new Memory Logger in order to do some assertion during test
func NewLogger() (*logger.Logger, *handler.Memory) {
	store := handler.NewMemory()
	return logger.NewLogger(store.Handle), store
}
