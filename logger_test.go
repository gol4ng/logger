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
		name string
		lvl  logger.Level
	}{
		{
			name: "test Debug(...)",
			lvl:  logger.DebugLevel,
		},
		{
			name: "test Info(...)",
			lvl:  logger.InfoLevel,
		},
		{
			name: "test Notice(...)",
			lvl:  logger.NoticeLevel,
		},
		{
			name: "test Warning(...)",
			lvl:  logger.WarningLevel,
		},
		{
			name: "test Error(...)",
			lvl:  logger.ErrorLevel,
		},
		{
			name: "test Critical(...)",
			lvl:  logger.CriticalLevel,
		},
		{
			name: "test Alert(...)",
			lvl:  logger.AlertLevel,
		},
		{
			name: "test Emergency(...)",
			lvl:  logger.EmergencyLevel,
		},
		{
			name: "test Log(custom level)",
			lvl:  logger.Level(127),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHandler := mocks.HandlerInterface{}

			mockHandler.On("Handle", mock.AnythingOfType("logger.Entry")).Return(func(e logger.Entry) error {
				assert.Equal(t, "log message", e.Message)
				assert.Equal(t, tt.lvl, e.Level)
				assert.Nil(t, e.Context)

				return nil
			})

			log := logger.NewLogger(&mockHandler)

			var err error
			switch tt.lvl {
			case logger.DebugLevel:
				err = log.Debug("log message", nil)
				break
			case logger.InfoLevel:
				err = log.Info("log message", nil)
				break
			case logger.NoticeLevel:
				err = log.Notice("log message", nil)
				break
			case logger.WarningLevel:
				err = log.Warning("log message", nil)
				break
			case logger.ErrorLevel:
				err = log.Error("log message", nil)
				break
			case logger.CriticalLevel:
				err = log.Critical("log message", nil)
				break
			case logger.AlertLevel:
				err = log.Alert("log message", nil)
				break
			case logger.EmergencyLevel:
				err = log.Emergency("log message", nil)
				break
			default:
				err = log.Log("log message", tt.lvl, nil)
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

	log := logger.NewLogger(&mockHandler)

	assert.Equal(t, err, log.Debug("log message", nil))
	assert.Equal(t, err, log.Info("log message", nil))
	assert.Equal(t, err, log.Notice("log message", nil))
	assert.Equal(t, err, log.Warning("log message", nil))
	assert.Equal(t, err, log.Error("log message", nil))
	assert.Equal(t, err, log.Critical("log message", nil))
	assert.Equal(t, err, log.Alert("log message", nil))
	assert.Equal(t, err, log.Emergency("log message", nil))
	assert.Equal(t, err, log.Log("log message", logger.Level(127), nil))
}

func TestNewNilLogger_Log(t *testing.T) {
	log := logger.NewNopLogger()

	assert.Nil(t, log.Debug("log message", nil))
	assert.Nil(t, log.Info("log message", nil))
	assert.Nil(t, log.Notice("log message", nil))
	assert.Nil(t, log.Warning("log message", nil))
	assert.Nil(t, log.Error("log message", nil))
	assert.Nil(t, log.Critical("log message", nil))
	assert.Nil(t, log.Alert("log message", nil))
	assert.Nil(t, log.Emergency("log message", nil))
	assert.Nil(t, log.Log("log message", logger.Level(127), nil))
}

func TestNewLogger_Wrap(t *testing.T) {
	mockHandler := mocks.HandlerInterface{}
	mockHandler.On("Handle", mock.AnythingOfType("logger.Entry")).Return(func(e logger.Entry) error {
		assert.Equal(t, "log message", e.Message)
		assert.Equal(t, logger.DebugLevel, e.Level)
		assert.Nil(t, e.Context)

		return nil
	})

	log := logger.NewLogger(&mockHandler)

	mockHandlerWrapper := mocks.HandlerInterface{}
	mockHandlerWrapper.On("Handle", mock.AnythingOfType("logger.Entry")).Return(func(e logger.Entry) error {
		return mockHandler.Handle(e)
	})

	log.Wrap(func(h logger.HandlerInterface) logger.HandlerInterface {
		return &mockHandlerWrapper
	})

	assert.Nil(t, log.Debug("log message", nil))

	mockHandler.AssertCalled(t, "Handle", mock.AnythingOfType("logger.Entry"))
	mockHandlerWrapper.AssertCalled(t, "Handle", mock.AnythingOfType("logger.Entry"))
}

func TestNewLogger_WrapNew(t *testing.T) {
	mockHandler := mocks.HandlerInterface{}
	mockHandler.On("Handle", mock.AnythingOfType("logger.Entry")).Return(func(e logger.Entry) error {
		assert.Equal(t, "log message", e.Message)
		assert.Equal(t, logger.DebugLevel, e.Level)
		assert.Nil(t, e.Context)

		return nil
	})

	log := logger.NewLogger(&mockHandler)

	mockHandlerWrapper := mocks.HandlerInterface{}
	mockHandlerWrapper.On("Handle", mock.AnythingOfType("logger.Entry")).Return(func(e logger.Entry) error {
		return mockHandler.Handle(e)
	})

	newLog := log.WrapNew(func(h logger.HandlerInterface) logger.HandlerInterface {
		return &mockHandlerWrapper
	})

	assert.NotEqual(t, log, newLog)
	assert.Nil(t, newLog.Debug("log message", nil))

	mockHandler.AssertCalled(t, "Handle", mock.AnythingOfType("logger.Entry"))
	mockHandlerWrapper.AssertCalled(t, "Handle", mock.AnythingOfType("logger.Entry"))
}
