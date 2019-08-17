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
	createdFile1 := &os.File{}
	createdFile2 := &os.File{}
	// mock os.File::Write method
	monkey.PatchInstanceMethod(reflect.TypeOf(f), "Write", func(file *os.File, p []byte) (n int, err error) {
		if i == 0 {
			assert.Equal(t, createdFile1, file)
			assert.Equal(t, []byte("my formatter return\n"), p)
			return 99, nil
		}
		if i == 1 {
			assert.Equal(t, createdFile2, file)
			assert.Equal(t, []byte("my formatter return\n"), p)
			return 99, nil
		}
		assert.True(t, false, "must not be reached")
		return 0, nil
	})
	// mock os.File::Close method in order not to return an error
	// on the second call of rotate, the TimeFileProvider will have a writer that is not nil (it will contain `createdFile1` that is technically nil
	// if we do not mock the `Close` method on `createdFile1` which will and up with a syscall.EINVAL error
	// as the code will pass here https://github.com/golang/go/blob/release-branch.go1.12/src/os/file_unix.go#L242
	monkey.PatchInstanceMethod(reflect.TypeOf(f), "Close", func(file *os.File) error {
		return nil
	})
	// mock os.OpenFile method
	monkey.Patch(os.OpenFile, func(name string, flag int, perm os.FileMode) (*os.File, error) {
		if i == 0 {
			assert.Equal(t, "fake_format_Thu Jan  1 1970 00", name)
			assert.Equal(t, os.O_CREATE|os.O_APPEND|os.O_WRONLY, flag)
			assert.Equal(t, os.FileMode(0666), perm)
			return createdFile1, nil
		}
		if i == 1 {
			assert.Equal(t, "fake_format_Thu Jan  1 1970 01", name)
			assert.Equal(t, os.O_CREATE|os.O_APPEND|os.O_WRONLY, flag)
			assert.Equal(t, os.FileMode(0666), perm)
			return createdFile2, nil
		}
		assert.True(t, false, "must not be reached")
		return nil, nil
	})
	// mock time.NewTicker method : override private ticker channel in order to be able to send ticks manually
	tickerChan := make(chan time.Time, 1)
	monkey.Patch(time.NewTicker, func(d time.Duration) *time.Ticker {
		assert.Equal(t, 10*time.Millisecond, d)
		return &time.Ticker{
			C: tickerChan,
		}
	})
	// mock time.Now method in order to return always the same time whenever the test is launched
	monkey.Patch(time.Now, func() time.Time { return time.Unix(i, 0) })
	defer monkey.UnpatchAll()
	// mock a basic formatter that will return "my formatter return" on any call of `Format`
	mockFormatter := mocks.FormatterInterface{}
	mockFormatter.On("Format", mock.AnythingOfType("logger.Entry")).Return("my formatter return")

	// create the handler to test
	h, err := handler.TimeRotateFileStream("fake_format_%s", "Mon Jan _2 2006 05", &mockFormatter, 10*time.Millisecond)
	assert.Nil(t, err)
	// call the handler a first time (i=0)
	assert.Nil(t, h(logger.Entry{}))
	// send tick to trigger a file rotation
	i++
	tickerChan <- time.Now()
	// call the handler a second time to check the rotation (i=1)
	assert.Nil(t, h(logger.Entry{}))
}

func TestNewLogRotateFileStream_Handle(t *testing.T) {
	var f *os.File
	monkey.PatchInstanceMethod(reflect.TypeOf(f), "Write", func(_ *os.File, p []byte) (n int, err error) {
		assert.Equal(t, []byte("my formatter return\n"), p)
		return 99, nil
	})

	file := os.File{}

	monkey.Patch(os.OpenFile, func(name string, flag int, perm os.FileMode) (*os.File, error) {
		assert.Equal(t, "fake_format_test", name)
		assert.Equal(t, os.O_CREATE|os.O_APPEND|os.O_WRONLY, flag)
		assert.Equal(t, os.FileMode(0666), perm)
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

	h, err := handler.LogRotateFileStream("test", "fake_format_%s", "Mon Jan _2 2006", &mockFormatter, 100*time.Millisecond)
	assert.Nil(t, err)

	assert.Nil(t, h(logger.Entry{Message: "test message", Level: logger.WarningLevel, Context: nil}))
	time.Sleep(200 * time.Millisecond)
	assert.Nil(t, h(logger.Entry{Message: "test message", Level: logger.WarningLevel, Context: nil}))
}
