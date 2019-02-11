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
	monkey.PatchInstanceMethod(reflect.TypeOf(f), "Close", func(_ *os.File) error {
		return errors.New("fake_file_close_error")
	})
	defer monkey.UnpatchAll()

	fileProvider := file_provider.TimeFileProvider("unused", "unused")
	newFile, err := fileProvider(&os.File{})
	assert.EqualError(t, err, "fake_file_close_error")
	assert.Nil(t, newFile)
}

func TestTimeFileProvider(t *testing.T) {
	f := os.File{}
	monkey.Patch(os.Create, func(name string) (*os.File, error) {
		assert.Equal(t, "fake_format_Mon Apr  7 1986", name)
		return &f, nil
	})
	monkey.Patch(time.Now, func() time.Time { return time.Unix(513216000, 0) })
	defer monkey.UnpatchAll()

	fileProvider := file_provider.TimeFileProvider("fake_format_%s", "Mon Jan _2 2006")
	newFile, err := fileProvider(nil)
	assert.Nil(t, err)
	assert.Equal(t, &f, newFile)
}
