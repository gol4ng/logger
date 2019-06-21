package writer_test

import (
	"errors"
	"io"
	"os"
	"reflect"
	"testing"

	"bou.ke/monkey"
	"github.com/gol4ng/logger/writer"
	"github.com/stretchr/testify/assert"
)

func TestRotateIoWriter_Write(t *testing.T) {
	f  := &os.File{}
	monkey.PatchInstanceMethod(reflect.TypeOf(f), "Write", func(_ *os.File, b []byte) (n int, err error) {
		assert.Equal(t, []byte("fake_data"), b)
		return 99, nil
	})
	defer monkey.UnpatchAll()

	file := os.File{}
	w, err := writer.NewRotateIoWriter(func(io.Writer) (io.Writer, error) {
		return &file, nil
	})
	assert.Nil(t, err)

	i, err := w.Write([]byte("fake_data"))
	assert.Equal(t, 99, i)
	assert.Nil(t, err)
}

func TestRotateIoWriter_Rotate(t *testing.T) {
	w, err := writer.NewRotateIoWriter(func(io.Writer) (io.Writer, error) {
		return &os.File{}, nil
	})
	assert.Nil(t, err)

	err = w.Rotate()
	assert.Nil(t, err)
}

func TestNewRotateIoWriter_WithError(t *testing.T) {
	_, err := writer.NewRotateIoWriter(func(io.Writer) (io.Writer, error) {
		return nil, errors.New("fake_file_provider_error")
	})
	assert.EqualError(t, err, "fake_file_provider_error")
}

func TestNewRotateIoWriter(t *testing.T) {
	_, err := writer.NewRotateIoWriter(func(io.Writer) (io.Writer, error) {
		return &os.File{}, nil
	})
	assert.Nil(t, err)
}
