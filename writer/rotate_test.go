package writer_test

import (
	"bou.ke/monkey"
	"errors"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/logger/mocks"
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
	file := os.File{}
	w, err := writer.NewRotateFileWriter(func(*os.File) (*os.File, error) {
		return &file, nil
	})
	assert.Nil(t, err)

	err = w.Rotate()
	assert.Nil(t, err)
}

func TestTimeRotateWriter_StartWithError(t *testing.T) {
	mockRotateWriter := mocks.RotateWriter{}
	mockRotateWriter.On("Rotate").Return(func() error { return errors.New("fake_rotate_error") })
	tr := writer.TimeRotateWriter{RotateWriter: &mockRotateWriter, Interval: 50 * time.Millisecond, PanicHandler: func(err error) {
		assert.EqualError(t, err, "fake_rotate_error")
	}}
	tr.Start()
	time.Sleep(100 * time.Millisecond)
	mockRotateWriter.AssertCalled(t, "Rotate")
}

func TestNewRotateFileWriter(t *testing.T) {
	f := os.File{}
	_, err := writer.NewRotateFileWriter(func(*os.File) (*os.File, error) {
		return &f, errors.New("fake_file_provider_error")
	})
	assert.EqualError(t, err, "fake_file_provider_error")
}

func TestNewTimeRotateWriter(t *testing.T) {
	mockRotateWriter := mocks.RotateWriter{}
	mockRotateWriter.On("Rotate").Return(func() error { return nil })
	writer.NewTimeRotateWriter(&mockRotateWriter, 100*time.Millisecond)
	time.Sleep(300 * time.Millisecond)
	mockRotateWriter.AssertCalled(t, "Rotate")
}

func TestNewTimeRotateFileWriter(t *testing.T) {
	f := os.File{}
	_, err := writer.NewTimeRotateFileWriter(func(*os.File) (*os.File, error) {
		return &f, errors.New("fake_file_provider_error")
	}, 1*time.Second)
	assert.EqualError(t, err, "fake_file_provider_error")
}
