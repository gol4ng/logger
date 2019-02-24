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

func TestTimeFileProvider_CloseWithError(t *testing.T) {
	var f *os.File
	existingFile := os.File{}
	monkey.PatchInstanceMethod(reflect.TypeOf(f), "Close", func(f *os.File) error {
		assert.Equal(t, &existingFile, f)
		return errors.New("fake_file_close_error")
	})
	defer monkey.UnpatchAll()

	fileProvider := file_provider.TimeFileProvider("unused", "unused")
	newFile, err := fileProvider(&existingFile)
	assert.EqualError(t, err, "fake_file_close_error")
	assert.Nil(t, newFile)
}

func TestTimeFileProvider(t *testing.T) {
	var f *os.File
	createdFile := os.File{}
	existingFile := os.File{}
	monkey.PatchInstanceMethod(reflect.TypeOf(f), "Close", func(f *os.File) error {
		assert.Equal(t, &existingFile, f)
		return nil
	})

	monkey.Patch(os.Create, func(name string) (*os.File, error) {
		assert.Equal(t, "fake_format_Thu Jan  1 1970 00", name)
		return &createdFile, nil
	})
	monkey.Patch(time.Now, func() time.Time { return time.Unix(0, 0) })
	defer monkey.UnpatchAll()

	fileProvider := file_provider.TimeFileProvider("fake_format_%s", "Mon Jan _2 2006 05")
	newFile, err := fileProvider(&existingFile)
	assert.Nil(t, err)
	assert.Equal(t, &createdFile, newFile)
}
