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
	i := int64(0)

	var f *os.File
	createdFile1 := &os.File{}
	createdFile2 := &os.File{}
	//createdFile1, _ := os.Create("file1")//os.File{}
	//createdFile2, _ := os.Create("file2")//os.File{}
	monkey.PatchInstanceMethod(reflect.TypeOf(f), "Write", func(aaa *os.File, p []byte) (n int, err error) {
		if i == 0 {
		println("writeFILE1", i, aaa)
			assert.Equal(t, createdFile1, aaa)
		}
		if i == 1 {
		println("writeFILE2", i, aaa)
			assert.Equal(t, createdFile2, aaa)
		}
		assert.Equal(t, []uint8("my formatter return\n"), p)
		return 99, nil
	})

	monkey.Patch(os.Create, func(name string) (*os.File, error) {
		println("create", i)
		if i == 0 {
			assert.Equal(t, "fake_format_Thu Jan  1 1970 00", name)
			return createdFile1, nil
		}
		if i == 1 {
			assert.Equal(t, "fake_format_Thu Jan  1 1970 01", name)
			return createdFile2, nil
		}

		print("OOPS")
		return nil, nil
	})

	tickerChan := make(chan time.Time, 1)
	monkey.Patch(time.NewTicker, func(d time.Duration) *time.Ticker {
		assert.Equal(t, 100*time.Millisecond, d)

		return &time.Ticker{
			C: tickerChan,
		}
	})

	monkey.Patch(time.Now, func() time.Time {return time.Unix(i, 0)})
	defer monkey.UnpatchAll()

	mockFormatter := mocks.FormatterInterface{}
	mockFormatter.On("Format", mock.AnythingOfType("logger.Entry")).Return("my formatter return")

	h, err := handler.NewTimeRotateFileStream("fake_format_%s", "Mon Jan _2 2006 05", &mockFormatter, 100*time.Millisecond)
	//file1
	assert.Nil(t, err)
	assert.Nil(t, h.Handle(logger.Entry{}))
	i++
	println(i)
	time.Sleep(200*time.Millisecond)
	tickerChan <- time.Now()
	//tickerChan <- time.Now()
	time.Sleep(200*time.Millisecond)
	//time.Sleep(100*time.Millisecond)
	//file2
	assert.Nil(t, h.Handle(logger.Entry{}))
	//
	//tickerChan <- time.Unix(i, 0)
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
