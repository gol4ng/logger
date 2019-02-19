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
	i := 0

	monkey.PatchInstanceMethod(reflect.TypeOf(f), "Write", func(_ *os.File, p []byte) (n int, err error) {
		assert.Equal(t, []uint8("my formatter return\n"), p)
		return 99, nil
	})

	file := os.File{}
	monkey.Patch(os.Create, func(name string) (*os.File, error) {
		if i == 1 {
			assert.Equal(t, "fake_format_Mon Apr  7 1986", name)
		}
		if i == 2 {
			assert.Equal(t, "fake_format_Tue Apr  7 1987", name)
		}
		println("titi")

		return &file, nil
	})

	c := make(chan time.Time, 1)

	monkey.Patch(time.NewTicker, func(d time.Duration) *time.Ticker {
		assert.Equal(t, 100*time.Millisecond, d)

		return &time.Ticker{
			C: c,
		}
	})

	r := []time.Time{
		time.Unix(513216000, 0),// 7/4/1986
		time.Unix(544752000, 0),// 7/4/1987
	}
	monkey.Patch(time.Now, func() time.Time {
		t := r[i]
		i++
		return t
	})
	defer monkey.UnpatchAll()

	mockFormatter := mocks.FormatterInterface{}
	mockFormatter.On("Format", mock.AnythingOfType("logger.Entry")).Return("my formatter return")

	h, err := handler.NewTimeRotateFileStream("fake_format_%s", "Mon Jan _2 2006", &mockFormatter, 100*time.Millisecond)
	assert.Nil(t, err)

	assert.Nil(t, h.Handle(logger.Entry{Message: "test message", Level: logger.WarningLevel, Context: nil}))
	c <- r[0]
	c <- r[1]

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
