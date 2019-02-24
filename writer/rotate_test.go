package writer_test

import (
	"bou.ke/monkey"
	"errors"
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"testing"

	"github.com/gol4ng/logger/writer"
)

func TestRotateFileWriter_Write(t *testing.T) {
	var f *os.File
	monkey.PatchInstanceMethod(reflect.TypeOf(f), "Write", func(_ *os.File, b []byte) (n int, err error) {
		assert.Equal(t, []byte("fake_data"), b)
		return 99, nil
	})
	defer monkey.UnpatchAll()

	file := os.File{}
	w, err := writer.NewRotateFileWriter(func(*os.File) (*os.File, error) {
		return &file, nil
	})
	assert.Nil(t, err)

	i, err := w.Write([]byte("fake_data"))
	assert.Equal(t, 99, i)
	assert.Nil(t, err)
}

func TestRotateFileWriter_Rotate(t *testing.T) {
	w, err := writer.NewRotateFileWriter(func(*os.File) (*os.File, error) {
		return &os.File{}, nil
	})
	assert.Nil(t, err)

	err = w.Rotate()
	assert.Nil(t, err)
}

func TestNewRotateFileWriter_WithError(t *testing.T) {
	_, err := writer.NewRotateFileWriter(func(*os.File) (*os.File, error) {
		return nil, errors.New("fake_file_provider_error")
	})
	assert.EqualError(t, err, "fake_file_provider_error")
}

func TestNewRotateFileWriter(t *testing.T) {
	_, err := writer.NewRotateFileWriter(func(*os.File) (*os.File, error) {
		return &os.File{}, nil
	})
	assert.Nil(t, err)
}
