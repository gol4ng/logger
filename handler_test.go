package logger_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/instabledesign/logger"
)

func TestNilHandler_Handle(t *testing.T) {
	tests := []struct {
		name    string
		handler *logger.NilHandler
	}{
		{name: "test nil handler struct", handler: &logger.NilHandler{}},
		{name: "test NewNilHandler()", handler: logger.NewNilHandler()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Nil(t, tt.handler.Handle(logger.Entry{}))
		})
	}
}
