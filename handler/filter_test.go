package handler_test

import (
	"testing"

	"github.com/instabledesign/logger/handler"

	"github.com/instabledesign/logger"
	"github.com/instabledesign/logger/mocks"
)

func TestFilter_HandleWithoutExclusion(t *testing.T) {
	logEntry := logger.Entry{Message: "test message", Level: logger.WarnLevel, Context: nil}

	mockHandler := mocks.HandlerInterface{}

	f := handler.NewFilter(&mockHandler, func(e logger.Entry) bool {
		return true
	})

	_ = f.Handle(logEntry)

	mockHandler.AssertNotCalled(t, "Handle", logEntry)
}

func TestFilter_HandleWithExclusion(t *testing.T) {
	logEntry := logger.Entry{Message: "test message", Level: logger.WarnLevel, Context: nil}

	mockHandler := mocks.HandlerInterface{}
	mockHandler.On("Handle", logEntry).Return(nil)

	f := handler.NewFilter(&mockHandler, func(e logger.Entry) bool {
		return false
	})

	_ = f.Handle(logEntry)

	mockHandler.AssertCalled(t, "Handle", logEntry)
}

//func TestNewMinLevelFilter(t *testing.T) {
//	logEntry := logger.Entry{Message: "test message", Level: logger.WarnLevel, Context: nil}
//
//	mockHandler := mocks.HandlerInterface{}
//	mockHandler.On("Handle", logEntry).Return(nil)
//
//	f := handler.NewMinLevelFilter(&mockHandler, logger.WarnLevel)
//
//	_ = f.Handle(logEntry)
//
//	if handleShouldBeCalled {
//		mockHandler.AssertCalled(t, "Handle", logEntry)
//		return
//	}
//
//	mockHandler.AssertNotCalled(t, "Handle", logEntry)
//}

//func TestNewMinLevelFilter(t *testing.T) {
//	logLevels := []logger.Level{
//		logger.DebugLevel,
//		logger.InfoLevel,
//		logger.WarnLevel,
//		logger.ErrorLevel,
//		logger.PanicLevel,
//		logger.FatalLevel,
//	}
//
//	tests := []struct {
//		name               string
//		e                  logger.Entry
//		logLevelsExclusion []bool
//	}{
//		{name: "test min lvl DEBUG", e: logger.Entry{Level: logger.DebugLevel}, logLevelsExclusion: []bool{true, true, true, true, true}},
//		//{name: "test min lvl INFO", e: logger.Entry{Level: logger.InfoLevel}, logLevelsExclusion: []bool{false, true, true, true, true}},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			mockHandler := mocks.HandlerInterface{}
//			mockHandler.On("Handle", tt.e).Return(nil)
//
//			for i, logLevel := range logLevels {
//				f := handler.NewMinLevelFilter(&mockHandler, logLevel)
//
//				_ = f.Handle(tt.e)
//
//				if tt.logLevelsExclusion[i] {
//					mockHandler.AssertCalled(t, "Handle", tt.e)
//					continue
//				}
//
//				mockHandler.AssertNotCalled(t, "Handle", tt.e)
//			}
//		})
//	}
//}
