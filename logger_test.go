package logger_test

import (
	"errors"
	"testing"

	"github.com/gol4ng/logger"
	"github.com/stretchr/testify/assert"
)

func TestLevel_String(t *testing.T) {
	tests := []struct {
		level    logger.Level
		expected string
	}{
		{level: logger.DebugLevel, expected: "debug"},
		{level: logger.InfoLevel, expected: "info"},
		{level: logger.NoticeLevel, expected: "notice"},
		{level: logger.WarningLevel, expected: "warning"},
		{level: logger.ErrorLevel, expected: "error"},
		{level: logger.CriticalLevel, expected: "critical"},
		{level: logger.AlertLevel, expected: "alert"},
		{level: logger.EmergencyLevel, expected: "emergency"},
		{level: 123, expected: "level(123)"},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, tt.level.String())
	}
}

func TestLogger_Log(t *testing.T) {
	tests := []struct {
		name  string
		level logger.Level
	}{
		{
			name:  "test Debug(...)",
			level: logger.DebugLevel,
		},
		{
			name:  "test Info(...)",
			level: logger.InfoLevel,
		},
		{
			name:  "test Notice(...)",
			level: logger.NoticeLevel,
		},
		{
			name:  "test Warning(...)",
			level: logger.WarningLevel,
		},
		{
			name:  "test Error(...)",
			level: logger.ErrorLevel,
		},
		{
			name:  "test Critical(...)",
			level: logger.CriticalLevel,
		},
		{
			name:  "test Alert(...)",
			level: logger.AlertLevel,
		},
		{
			name:  "test Emergency(...)",
			level: logger.EmergencyLevel,
		},
		{
			name:  "test Log(custom level)",
			level: logger.Level(127),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := logger.NewLogger(func(entry logger.Entry) error {
				assert.Equal(t, "l message", entry.Message)
				assert.Equal(t, tt.level, entry.Level)
				assert.Nil(t, entry.Context)

				return nil
			})

			var err error
			switch tt.level {
			case logger.DebugLevel:
				err = l.Debug("l message", nil)
				break
			case logger.InfoLevel:
				err = l.Info("l message", nil)
				break
			case logger.NoticeLevel:
				err = l.Notice("l message", nil)
				break
			case logger.WarningLevel:
				err = l.Warning("l message", nil)
				break
			case logger.ErrorLevel:
				err = l.Error("l message", nil)
				break
			case logger.CriticalLevel:
				err = l.Critical("l message", nil)
				break
			case logger.AlertLevel:
				err = l.Alert("l message", nil)
				break
			case logger.EmergencyLevel:
				err = l.Emergency("l message", nil)
				break
			default:
				err = l.Log("l message", tt.level, nil)
			}
			assert.Nil(t, err)
		})
	}
}

func TestNewLogger_LogWithError(t *testing.T) {
	err := errors.New("my error")
	l := logger.NewLogger(func(entry logger.Entry) error {
		return err
	})

	assert.Equal(t, err, l.Debug("l message", nil))
	assert.Equal(t, err, l.Info("l message", nil))
	assert.Equal(t, err, l.Notice("l message", nil))
	assert.Equal(t, err, l.Warning("l message", nil))
	assert.Equal(t, err, l.Error("l message", nil))
	assert.Equal(t, err, l.Critical("l message", nil))
	assert.Equal(t, err, l.Alert("l message", nil))
	assert.Equal(t, err, l.Emergency("l message", nil))
	assert.Equal(t, err, l.Log("l message", logger.Level(127), nil))
}

func TestNewNilLogger_Log(t *testing.T) {
	l := logger.NewNopLogger()

	assert.Nil(t, l.Debug("l message", nil))
	assert.Nil(t, l.Info("l message", nil))
	assert.Nil(t, l.Notice("l message", nil))
	assert.Nil(t, l.Warning("l message", nil))
	assert.Nil(t, l.Error("l message", nil))
	assert.Nil(t, l.Critical("l message", nil))
	assert.Nil(t, l.Alert("l message", nil))
	assert.Nil(t, l.Emergency("l message", nil))
	assert.Nil(t, l.Log("l message", logger.Level(127), nil))
}

func TestNewLogger_Wrap(t *testing.T) {
	mockHandlerCalled := false
	mockHandler := func(entry logger.Entry) error {
		mockHandlerCalled = true
		assert.Equal(t, "l message", entry.Message)
		assert.Equal(t, logger.DebugLevel, entry.Level)
		assert.Nil(t, entry.Context)

		return nil
	}

	l := logger.NewLogger(mockHandler)

	mockHandlerWrapperCalled := false
	mockHandlerWrapper := func(entry logger.Entry) error {
		mockHandlerWrapperCalled = true
		return mockHandler(entry)
	}

	l.Wrap(func(h logger.HandlerInterface) logger.HandlerInterface {
		return mockHandlerWrapper
	})

	assert.Nil(t, l.Debug("l message", nil))

	assert.True(t, mockHandlerCalled)
	assert.True(t, mockHandlerWrapperCalled)
}

func TestNewLogger_WrapNew(t *testing.T) {
	mockHandlerCalled := false
	mockHandler := func(entry logger.Entry) error {
		mockHandlerCalled = true
		assert.Equal(t, "l message", entry.Message)
		assert.Equal(t, logger.DebugLevel, entry.Level)
		assert.Nil(t, entry.Context)

		return nil
	}

	l := logger.NewLogger(mockHandler)

	mockHandlerWrapperCalled := false
	mockHandlerWrapper := func(entry logger.Entry) error {
		mockHandlerWrapperCalled = true
		return mockHandler(entry)
	}

	l2 := l.WrapNew(func(h logger.HandlerInterface) logger.HandlerInterface {
		return mockHandlerWrapper
	})

	assert.Nil(t, l.Debug("l message", nil))
	assert.True(t, mockHandlerCalled)
	assert.False(t, mockHandlerWrapperCalled)

	mockHandlerCalled = false
	mockHandlerWrapperCalled = false

	assert.Nil(t, l2.Debug("l message", nil))
	assert.True(t, mockHandlerCalled)
	assert.True(t, mockHandlerWrapperCalled)
}
