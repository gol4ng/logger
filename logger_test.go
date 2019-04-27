package logger_test

import (
	"testing"

	"github.com/pkg/errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/mocks"
)

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
			mockHandler := mocks.HandlerInterface{}

			mockHandler.On("Handle", mock.AnythingOfType("logger.Entry")).Return(func(entry logger.Entry) error {
				assert.Equal(t, "l message", entry.Message)
				assert.Equal(t, tt.level, entry.Level)
				assert.Nil(t, entry.Context)

				return nil
			})

			l := logger.NewLogger(&mockHandler)

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

			mockHandler.AssertCalled(t, "Handle", mock.AnythingOfType("logger.Entry"))
		})
	}
}

func TestNewLogger_LogWithError(t *testing.T) {
	mockHandler := mocks.HandlerInterface{}
	err := errors.New("my error")

	mockHandler.On("Handle", mock.AnythingOfType("logger.Entry")).Return(err)

	l := logger.NewLogger(&mockHandler)

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
	mockHandler := mocks.HandlerInterface{}
	mockHandler.On("Handle", mock.AnythingOfType("logger.Entry")).Return(func(entry logger.Entry) error {
		assert.Equal(t, "l message", entry.Message)
		assert.Equal(t, logger.DebugLevel, entry.Level)
		assert.Nil(t, entry.Context)

		return nil
	})

	l := logger.NewLogger(&mockHandler)

	mockHandlerWrapper := mocks.HandlerInterface{}
	mockHandlerWrapper.On("Handle", mock.AnythingOfType("logger.Entry")).Return(func(entry logger.Entry) error {
		return mockHandler.Handle(entry)
	})

	l.Wrap(func(h logger.HandlerInterface) logger.HandlerInterface {
		return &mockHandlerWrapper
	})

	assert.Nil(t, l.Debug("l message", nil))

	mockHandler.AssertCalled(t, "Handle", mock.AnythingOfType("logger.Entry"))
	mockHandlerWrapper.AssertCalled(t, "Handle", mock.AnythingOfType("logger.Entry"))
}

func TestNewLogger_WrapNew(t *testing.T) {
	mockHandler := mocks.HandlerInterface{}
	mockHandler.On("Handle", mock.AnythingOfType("logger.Entry")).Return(func(entry logger.Entry) error {
		assert.Equal(t, "l message", entry.Message)
		assert.Equal(t, logger.DebugLevel, entry.Level)
		assert.Nil(t, entry.Context)

		return nil
	})

	l := logger.NewLogger(&mockHandler)

	mockHandlerWrapper := mocks.HandlerInterface{}
	mockHandlerWrapper.On("Handle", mock.AnythingOfType("logger.Entry")).Return(func(entry logger.Entry) error {
		return mockHandler.Handle(entry)
	})

	newLog := l.WrapNew(func(h logger.HandlerInterface) logger.HandlerInterface {
		return &mockHandlerWrapper
	})

	assert.NotEqual(t, l, newLog)
	assert.Nil(t, newLog.Debug("l message", nil))

	mockHandler.AssertCalled(t, "Handle", mock.AnythingOfType("logger.Entry"))
	mockHandlerWrapper.AssertCalled(t, "Handle", mock.AnythingOfType("logger.Entry"))
}
