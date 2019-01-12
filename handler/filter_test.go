package handler_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/instabledesign/logger"
	"github.com/instabledesign/logger/handler"
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

func TestNewMinLevelFilter(t *testing.T) {
	logLevels := [6]logger.Level{
		logger.DebugLevel,
		logger.InfoLevel,
		logger.WarnLevel,
		logger.ErrorLevel,
		logger.PanicLevel,
		logger.FatalLevel,
	}

	tests := []struct {
		name             string
		lvl              logger.Level
		logLevelsHandled [6]bool
	}{
		{name: "test min lvl DEBUG", lvl: logger.DebugLevel, logLevelsHandled: [6]bool{true, true, true, true, true, true}},
		{name: "test min lvl INFO", lvl: logger.InfoLevel, logLevelsHandled: [6]bool{false, true, true, true, true, true}},
		{name: "test min lvl WARN", lvl: logger.WarnLevel, logLevelsHandled: [6]bool{false, false, true, true, true, true}},
		{name: "test min lvl ERRPR", lvl: logger.ErrorLevel, logLevelsHandled: [6]bool{false, false, false, true, true, true}},
		{name: "test min lvl PÄNIC", lvl: logger.PanicLevel, logLevelsHandled: [6]bool{false, false, false, false, true, true}},
		{name: "test min lvl FATAL", lvl: logger.FatalLevel, logLevelsHandled: [6]bool{false, false, false, false, false, true}},
	}
	for _, tt := range tests {
		for i, logLevel := range logLevels {
			t.Run(tt.name, func(t *testing.T) {
				e := logger.Entry{Level: logLevel}
				mockHandler := mocks.HandlerInterface{}
				mockHandler.On("Handle", e).Return(nil)

				f := handler.NewMinLevelFilter(&mockHandler, tt.lvl)

				_ = f.Handle(e)

				if tt.logLevelsHandled[i] {
					mockHandler.AssertCalled(t, "Handle", e)
					return
				}

				mockHandler.AssertNotCalled(t, "Handle", e)
			})
		}
	}
}

func TestNewRangeLevelFilter(t *testing.T) {
	logLevels := [6]logger.Level{
		logger.DebugLevel,
		logger.InfoLevel,
		logger.WarnLevel,
		logger.ErrorLevel,
		logger.PanicLevel,
		logger.FatalLevel,
	}

	tests := []struct {
		name             string
		minLvl           logger.Level
		maxLvl           logger.Level
		logLevelsHandled [6]bool
	}{
		{name: "test between DEBUG/FATAL with log level %s", minLvl: logger.DebugLevel, maxLvl: logger.FatalLevel, logLevelsHandled: [6]bool{true, true, true, true, true, true}},
		{name: "test between INFO/PANIC with log level %s", minLvl: logger.InfoLevel, maxLvl: logger.PanicLevel, logLevelsHandled: [6]bool{false, true, true, true, true, false}},
	}
	for _, tt := range tests {
		for i, logLevel := range logLevels {
			t.Run(fmt.Sprintf(tt.name, logLevel), func(t *testing.T) {
				e := logger.Entry{Level: logLevel}
				mockHandler := mocks.HandlerInterface{}
				mockHandler.On("Handle", e).Return(nil)

				f := handler.NewRangeLevelFilter(&mockHandler, tt.minLvl, tt.maxLvl)

				_ = f.Handle(e)

				if tt.logLevelsHandled[i] {
					mockHandler.AssertCalled(t, "Handle", e)
					return
				}

				mockHandler.AssertNotCalled(t, "Handle", e)
			})
		}
	}
}

func TestNewRangeLevelFilterWithPanic(t *testing.T) {
	mockHandler := mocks.HandlerInterface{}
	assert.PanicsWithValue(
		t,
		"invalid logger range level : Min level must be lower than max level",
		func() { handler.NewRangeLevelFilter(&mockHandler, logger.FatalLevel, logger.DebugLevel) })
}
