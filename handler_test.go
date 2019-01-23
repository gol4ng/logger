package logger_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/instabledesign/logger"
)

func TestNopHandler_Handle(t *testing.T) {
	tests := []struct {
		name    string
		handler *logger.NopHandler
	}{
		{name: "test nil handler struct", handler: &logger.NopHandler{}},
		{name: "test NewNopHandler()", handler: logger.NewNopHandler()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Nil(t, tt.handler.Handle(logger.Entry{}))
		})
	}
}
