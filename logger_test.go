package logger_test

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/instabledesign/logger"
	"github.com/instabledesign/logger/mocks"
)

func TestLogger_Debug(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test nil handler struct",
		},
		{
			name: "test NewNilHandler()",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHandler := mocks.HandlerInterface{}

			// TODO Need to test the handler argument value
			mockHandler.On("Handle", mock.AnythingOfType("logger.Entry")).Return(nil)

			log := logger.NewLogger(&mockHandler)
			log.Info("log message", nil)

			mockHandler.AssertCalled(t, "Handle", mock.AnythingOfType("logger.Entry"))
		})
	}
}
