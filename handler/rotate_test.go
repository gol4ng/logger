package handler_test

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/handler"
	"github.com/gol4ng/logger/mocks"
)

func TestNewTimeRotateFileStream_Handle(t *testing.T) {
	i := 0
	monkey.Patch(time.Now, func() time.Time {
		r := []time.Time{
			time.Unix(513216000, 0),
			time.Unix(513217000, 0),
		}[i]
		i++
		return r
	})
	defer monkey.UnpatchAll()

	mockFormatter := mocks.FormatterInterface{}
	mockFormatter.On("Format", mock.AnythingOfType("logger.Entry")).Return("my formatter return")

	h, err := handler.NewTimeRotateFileStream( "./%s.log", time.Stamp, &mockFormatter, 1*time.Second)
	//h, err := handler.NewLogRotateFileStream("test", os.TempDir()+"%s.log", time.Stamp, &mockFormatter, 1*time.Second)
	assert.Nil(t, err)

	assert.Nil(t, h.Handle(logger.Entry{Message: "test message", Level: logger.WarningLevel, Context: nil}))

	//TODO test file content
}

func TestNewLogRotateFileStream_Handle(t *testing.T) {
	mockFormatter := mocks.FormatterInterface{}
	mockFormatter.On("Format", mock.AnythingOfType("logger.Entry")).Return("my formatter return")

	h, err := handler.NewLogRotateFileStream("test", "./%s.log", time.Stamp, &mockFormatter, 1*time.Second)
	//h, err := handler.NewLogRotateFileStream("test", os.TempDir()+"%s.log", time.Stamp, &mockFormatter, 1*time.Second)
	assert.Nil(t, err)

	assert.Nil(t, h.Handle(logger.Entry{Message: "test message", Level: logger.WarningLevel, Context: nil}))

	//TODO test file content
}
