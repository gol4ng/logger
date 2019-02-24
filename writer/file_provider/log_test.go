package file_provider_test

import (
	"errors"
	"os"
	"reflect"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/logger/writer/file_provider"
)

func TestLogFileProvider_RenameWithError(t *testing.T) {
	monkey.Patch(os.Rename, func(oldpath, newpath string) error {
		assert.Equal(t, "fake_format_fake_name", oldpath)
		assert.Equal(t, "fake_format_Thu Jan  1 1970 00", newpath)
		return errors.New("fake_rename_error")
	})
	monkey.Patch(time.Now, func() time.Time {return time.Unix(0, 0)})
	defer monkey.UnpatchAll()

	fileProvider := file_provider.LogFileProvider("fake_name", "fake_format_%s", "Mon Jan _2 2006 05")
	newFile, err := fileProvider(&os.File{})
	assert.EqualError(t, err, "fake_rename_error")
	assert.Nil(t, newFile)
}

func TestLogFileProvider_CloseWithError(t *testing.T) {
	var f *os.File
	monkey.PatchInstanceMethod(reflect.TypeOf(f), "Close", func(_ *os.File) error {
		return errors.New("fake_file_close_error")
	})
	monkey.Patch(os.Rename, func(oldpath, newpath string) error {
		assert.Equal(t, "fake_format_fake_name", oldpath)
		assert.Equal(t, "fake_format_Thu Jan  1 1970 00", newpath)
		return nil
	})
	monkey.Patch(time.Now, func() time.Time { return time.Unix(0, 0) })
	defer monkey.UnpatchAll()

	w := file_provider.LogFileProvider("fake_name", "fake_format_%s", "Mon Jan _2 2006 05")
	newFile, err := w(&os.File{})
	assert.EqualError(t, err, "fake_file_close_error")
	assert.Nil(t, newFile)
}

func TestLogFileProvider(t *testing.T) {
	var f *os.File
	createdFile := os.File{}
	existingFile := os.File{}
	monkey.PatchInstanceMethod(reflect.TypeOf(f), "Close", func(f *os.File) error {
		assert.Equal(t, &existingFile, f)
		return nil
	})
	monkey.Patch(os.Rename, func(oldpath, newpath string) error {
		assert.Equal(t, "fake_format_fake_name", oldpath)
		assert.Equal(t, "fake_format_Thu Jan  1 1970 00", newpath)
		return nil
	})
	monkey.Patch(os.Create, func(name string) (*os.File, error) {
		assert.Equal(t, "fake_format_fake_name", name)
		return &createdFile, nil
	})
	monkey.Patch(time.Now, func() time.Time { return time.Unix(0, 0) })
	defer monkey.UnpatchAll()

	w := file_provider.LogFileProvider("fake_name", "fake_format_%s", "Mon Jan _2 2006 05")
	newFile, err := w(&existingFile)
	assert.Nil(t, err)
	assert.Equal(t, &createdFile, newFile)
}
