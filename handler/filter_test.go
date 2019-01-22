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
	logEntry := logger.Entry{}

	mockHandler := mocks.HandlerInterface{}

	h := handler.NewFilter(&mockHandler, func(e logger.Entry) bool {
		return true
	})

	assert.Nil(t, h.Handle(logEntry))

	mockHandler.AssertNotCalled(t, "Handle", logEntry)
}

func TestFilter_HandleWithExclusion(t *testing.T) {
	logEntry := logger.Entry{}

	mockHandler := mocks.HandlerInterface{}
	mockHandler.On("Handle", logEntry).Return(nil)

	h := handler.NewFilter(&mockHandler, func(e logger.Entry) bool {
		return false
	})

	assert.Nil(t, h.Handle(logEntry))

	mockHandler.AssertCalled(t, "Handle", logEntry)
}

func TestNewMinLevelFilter(t *testing.T) {
	logLevels := [8]logger.Level{
		logger.DebugLevel,
		logger.InfoLevel,
		logger.NoticeLevel,
		logger.WarningLevel,
		logger.ErrorLevel,
		logger.CriticalLevel,
		logger.AlertLevel,
		logger.EmergencyLevel,
	}

	tests := []struct {
		name             string
		lvl              logger.Level
		logLevelsHandled [8]bool
	}{
		{name: "test min lvl DEBUG", lvl: logger.DebugLevel, logLevelsHandled: [8]bool{true, true, true, true, true, true, true, true}},
		{name: "test min lvl INFO", lvl: logger.InfoLevel, logLevelsHandled: [8]bool{false, true, true, true, true, true, true, true}},
		{name: "test min lvl NOTICE", lvl: logger.NoticeLevel, logLevelsHandled: [8]bool{false, false, true, true, true, true, true, true}},
		{name: "test min lvl WARNING", lvl: logger.WarningLevel, logLevelsHandled: [8]bool{false, false, false, true, true, true, true, true}},
		{name: "test min lvl ERROR", lvl: logger.ErrorLevel, logLevelsHandled: [8]bool{false, false, false, false, true, true, true, true}},
		{name: "test min lvl CRITICAL", lvl: logger.CriticalLevel, logLevelsHandled: [8]bool{false, false, false, false, false, true, true, true}},
		{name: "test min lvl ALERT", lvl: logger.AlertLevel, logLevelsHandled: [8]bool{false, false, false, false, false, false, true, true}},
		{name: "test min lvl EMERGENCY", lvl: logger.EmergencyLevel, logLevelsHandled: [8]bool{false, false, false, false, false, false, false, true}},
	}
	for _, tt := range tests {
		for i, logLevel := range logLevels {
			t.Run(tt.name, func(t *testing.T) {
				e := logger.Entry{Level: logLevel}
				mockHandler := mocks.HandlerInterface{}
				mockHandler.On("Handle", e).Return(nil)

				h := handler.NewMinLevelFilter(&mockHandler, tt.lvl)

				assert.Nil(t, h.Handle(e))

				if tt.logLevelsHandled[i] {
					mockHandler.AssertCalled(t, "Handle", e)
					return
				}

				mockHandler.AssertNotCalled(t, "Handle", e)
			})
			t.Run(tt.name+" wrapped", func(t *testing.T) {
				e := logger.Entry{Level: logLevel}
				mockHandler := mocks.HandlerInterface{}
				mockHandler.On("Handle", e).Return(nil)

				h := handler.NewMinLevelWrapper(tt.lvl)(&mockHandler)

				assert.Nil(t, h.Handle(e))

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
	logLevels := [8]logger.Level{
		logger.DebugLevel,
		logger.InfoLevel,
		logger.NoticeLevel,
		logger.WarningLevel,
		logger.ErrorLevel,
		logger.CriticalLevel,
		logger.AlertLevel,
		logger.EmergencyLevel,
	}

	tests := []struct {
		name             string
		minLvl           logger.Level
		maxLvl           logger.Level
		logLevelsHandled [8]bool
	}{
		{name: "test between DEBUG/EMERGENCY with log level %s", minLvl: logger.DebugLevel, maxLvl: logger.EmergencyLevel, logLevelsHandled: [8]bool{true, true, true, true, true, true, true, true}},
		{name: "test between INFO/ALERT with log level %s", minLvl: logger.InfoLevel, maxLvl: logger.AlertLevel, logLevelsHandled: [8]bool{false, true, true, true, true, true, true, false}},
	}
	for _, tt := range tests {
		for i, logLevel := range logLevels {
			t.Run(fmt.Sprintf(tt.name, logLevel), func(t *testing.T) {
				e := logger.Entry{Level: logLevel}
				mockHandler := mocks.HandlerInterface{}
				mockHandler.On("Handle", e).Return(nil)

				h := handler.NewRangeLevelFilter(&mockHandler, tt.minLvl, tt.maxLvl)

				assert.Nil(t, h.Handle(e))

				if tt.logLevelsHandled[i] {
					mockHandler.AssertCalled(t, "Handle", e)
					return
				}

				mockHandler.AssertNotCalled(t, "Handle", e)
			})
			t.Run(fmt.Sprintf(tt.name+" wrapped", logLevel), func(t *testing.T) {
				e := logger.Entry{Level: logLevel}
				mockHandler := mocks.HandlerInterface{}
				mockHandler.On("Handle", e).Return(nil)

				h := handler.NewRangeLevelWrapper(tt.minLvl, tt.maxLvl)(&mockHandler)

				assert.Nil(t, h.Handle(e))

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
		func() { handler.NewRangeLevelFilter(&mockHandler, logger.ErrorLevel, logger.DebugLevel) })
}

func TestNewFilterWrapped(t *testing.T) {
	logLevels := [8]logger.Level{
		logger.DebugLevel,
		logger.InfoLevel,
		logger.NoticeLevel,
		logger.WarningLevel,
		logger.ErrorLevel,
		logger.CriticalLevel,
		logger.AlertLevel,
		logger.EmergencyLevel,
	}

	tests := []struct {
		name             string
		filterFn         func(e logger.Entry) bool
		logLevelsHandled [8]bool
	}{
		{
			name:             "test filter DEBUG/ERROR with log level %s",
			filterFn:         func(e logger.Entry) bool { return e.Level != logger.DebugLevel && e.Level != logger.ErrorLevel },
			logLevelsHandled: [8]bool{true, false, false, false, true, false, false, false},
		},
		{
			name:             "test filter INFO/ALERT with log level %s",
			filterFn:         func(e logger.Entry) bool { return e.Level != logger.InfoLevel && e.Level != logger.AlertLevel },
			logLevelsHandled: [8]bool{false, true, false, false, false, false, true, false},
		},
	}
	for _, tt := range tests {
		for i, logLevel := range logLevels {
			t.Run(fmt.Sprintf(tt.name, logLevel), func(t *testing.T) {
				e := logger.Entry{Level: logLevel}
				mockHandler := mocks.HandlerInterface{}
				mockHandler.On("Handle", e).Return(nil)

				h := handler.NewFilterWrapper(tt.filterFn)(&mockHandler)

				assert.Nil(t, h.Handle(e))

				if tt.logLevelsHandled[i] {
					mockHandler.AssertCalled(t, "Handle", e)
					return
				}

				mockHandler.AssertNotCalled(t, "Handle", e)
			})
		}
	}
}
