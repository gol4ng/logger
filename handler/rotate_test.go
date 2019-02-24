package handler_test

import (
	"os"
	"reflect"
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
	i := int64(0)

	var f *os.File
	createdFile1 := os.File{}
	createdFile2 := os.File{}
	monkey.PatchInstanceMethod(reflect.TypeOf(f), "Write", func(file *os.File, p []byte) (n int, err error) {
		if i == 0 {
			assert.Equal(t, &createdFile1, file)
			assert.Equal(t, []uint8("my formatter return\n"), p)
			return 99, nil
		}
		if i == 1 {
			assert.Equal(t, &createdFile2, file)
			assert.Equal(t, []uint8("my formatter return\n"), p)
			return 99, nil
		}

		assert.True(t, false, "must not be reached")
		return 0, nil
	})

	monkey.Patch(os.Create, func(name string) (*os.File, error) {
		if i == 0 {
			assert.Equal(t, "fake_format_Thu Jan  1 1970 00", name)
			return &createdFile1, nil
		}
		if i == 1 {
			assert.Equal(t, "fake_format_Thu Jan  1 1970 01", name)
			return &createdFile2, nil
		}

		assert.True(t, false, "must not be reached")
		return nil, nil
	})

	tickerChan := make(chan time.Time, 1)
	monkey.Patch(time.NewTicker, func(d time.Duration) *time.Ticker {
		assert.Equal(t, 100*time.Millisecond, d)

		return &time.Ticker{
			C: tickerChan,
		}
	})

	monkey.Patch(time.Now, func() time.Time { return time.Unix(i, 0) })
	defer monkey.UnpatchAll()

	mockFormatter := mocks.FormatterInterface{}
	mockFormatter.On("Format", mock.AnythingOfType("logger.Entry")).Return("my formatter return")

	h, err := handler.NewTimeRotateFileStream("fake_format_%s", "Mon Jan _2 2006 05", &mockFormatter, 100*time.Millisecond)
	assert.Nil(t, err)
	assert.Nil(t, h.Handle(logger.Entry{}))

	i++
	tickerChan <- time.Now()
	assert.Nil(t, h.Handle(logger.Entry{}))
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
