package logger_test

import (
	"testing"

	"github.com/pkg/errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/instabledesign/logger"
	"github.com/instabledesign/logger/mocks"
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
			name: "test Warn(...)",
			lvl:  logger.WarnLevel,
		},
		{
			name: "test Error(...)",
			lvl:  logger.ErrorLevel,
		},
		{
			name: "test Panic(...)",
			lvl:  logger.PanicLevel,
		},
		{
			name: "test Fatal(...)",
			lvl:  logger.FatalLevel,
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
			case logger.WarnLevel:
				err = log.Warn("log message", nil)
				break
			case logger.ErrorLevel:
				err = log.Error("log message", nil)
				break
			case logger.PanicLevel:
				err = log.Panic("log message", nil)
				break
			case logger.FatalLevel:
				err = log.Fatal("log message", nil)
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
	assert.Equal(t, err, log.Warn("log message", nil))
	assert.Equal(t, err, log.Error("log message", nil))
	assert.Equal(t, err, log.Panic("log message", nil))
	assert.Equal(t, err, log.Fatal("log message", nil))
	assert.Equal(t, err, log.Log("log message", logger.Level(127), nil))
}

func TestNewNilLogger_Log(t *testing.T) {
	log := logger.NewNilLogger()

	assert.Nil(t, log.Debug("log message", nil))
	assert.Nil(t, log.Info("log message", nil))
	assert.Nil(t, log.Warn("log message", nil))
	assert.Nil(t, log.Error("log message", nil))
	assert.Nil(t, log.Panic("log message", nil))
	assert.Nil(t, log.Fatal("log message", nil))
	assert.Nil(t, log.Log("log message", logger.Level(127), nil))
}
