package provider_test

import (
	"errors"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/gol4ng/logger/writer/provider"
	"github.com/stretchr/testify/assert"
)

func TestTimeFileProvider_CloseWithError(t *testing.T) {
	var f *os.File
	existingFile := os.File{}
	patch := gomonkey.NewPatches()
	patch.ApplyMethod(reflect.TypeOf(f), "Close", func(f *os.File) error {
		assert.Equal(t, &existingFile, f)
		return errors.New("fake_file_close_error")
	})
	defer patch.Reset()

	fileProvider := provider.TimeFileProvider("unused", "unused")
	newFile, err := fileProvider(&existingFile)
	assert.EqualError(t, err, "fake_file_close_error")
	assert.Nil(t, newFile)
}

func TestTimeFileProvider(t *testing.T) {
	var f *os.File
	createdFile := os.File{}
	existingFile := os.File{}
	patch := gomonkey.NewPatches()
	patch.ApplyMethod(reflect.TypeOf(f), "Close", func(f *os.File) error {
		assert.Equal(t, &existingFile, f)
		return nil
	})

	patch.ApplyFunc(os.OpenFile, func(name string, flag int, perm os.FileMode) (*os.File, error) {
		assert.Equal(t, "fake_format_Thu Jan  1 1970 00", name)
		assert.Equal(t, os.O_CREATE|os.O_APPEND|os.O_WRONLY, flag)
		assert.Equal(t, os.FileMode(0666), perm)
		return &createdFile, nil
	})

	patch.ApplyFunc(time.Now, func() time.Time { return time.Unix(0, 0) })
	defer patch.Reset()

	fileProvider := provider.TimeFileProvider("fake_format_%s", "Mon Jan _2 2006 05")
	newFile, err := fileProvider(&existingFile)
	assert.Nil(t, err)
	assert.Equal(t, &createdFile, newFile)
}
