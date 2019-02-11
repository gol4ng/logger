package handler_test

import (
	"bou.ke/monkey"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/handler"
	"github.com/gol4ng/logger/mocks"
)

func TestNewTimeRotateFileStream_Handle(t *testing.T) {
	var f *os.File
	monkey.PatchInstanceMethod(reflect.TypeOf(f), "Write", func(_ *os.File, p []byte) (n int, err error) {
		assert.Equal(t, []uint8("my formatter return\n"), p)
		return 99, nil
	})

	file := os.File{}
	monkey.Patch(os.Create, func(name string) (*os.File, error) {
		assert.Equal(t, "fake_format_Mon Apr  7 1986", name)
		return &file, nil
	})

	i := 0
	monkey.Patch(time.Now, func() time.Time {
		r := []time.Time{
			time.Unix(513216000, 0),
			time.Unix(513217000, 0),
			time.Unix(513217000, 0),
		}[i]
		i++
		return r
	})
	defer monkey.UnpatchAll()

	mockFormatter := mocks.FormatterInterface{}
	mockFormatter.On("Format", mock.AnythingOfType("logger.Entry")).Return("my formatter return")

	h, err := handler.NewTimeRotateFileStream("fake_format_%s", "Mon Jan _2 2006", &mockFormatter, 100*time.Millisecond)
	assert.Nil(t, err)

	assert.Nil(t, h.Handle(logger.Entry{Message: "test message", Level: logger.WarningLevel, Context: nil}))
	time.Sleep(300 * time.Millisecond)
	assert.Nil(t, h.Handle(logger.Entry{Message: "test message", Level: logger.WarningLevel, Context: nil}))
}

func TestNewLogRotateFileStream_Handle(t *testing.T) {
	var f *os.File
	monkey.PatchInstanceMethod(reflect.TypeOf(f), "Write", func(_ *os.File, p []byte) (n int, err error) {
		assert.Equal(t, []uint8("my formatter return\n"), p)
		return 99, nil
	})

	file := os.File{}
	monkey.Patch(os.Create, func(name string) (*os.File, error) {
		assert.Equal(t, "fake_format_test", name)
		return &file, nil
	})
	monkey.Patch(os.Rename, func(oldpath, newpath string) error {
		assert.Equal(t, "fake_format_test", oldpath)
		assert.Equal(t, "fake_format_Mon Apr  7 1986", newpath)
		return nil
	})
	monkey.Patch(time.Now, func() time.Time { return time.Unix(513216000, 0) })
	defer monkey.UnpatchAll()

	mockFormatter := mocks.FormatterInterface{}
	mockFormatter.On("Format", mock.AnythingOfType("logger.Entry")).Return("my formatter return")

	h, err := handler.NewLogRotateFileStream("test", "fake_format_%s", "Mon Jan _2 2006", &mockFormatter, 100*time.Millisecond)
	assert.Nil(t, err)

	assert.Nil(t, h.Handle(logger.Entry{Message: "test message", Level: logger.WarningLevel, Context: nil}))
	time.Sleep(200 * time.Millisecond)
	assert.Nil(t, h.Handle(logger.Entry{Message: "test message", Level: logger.WarningLevel, Context: nil}))
}
