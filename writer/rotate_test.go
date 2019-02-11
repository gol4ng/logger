package writer_test

import (
	"errors"
	"github.com/gol4ng/logger/mocks"
	"os"
	"reflect"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/logger/writer"
)

func TestTimeFileProvider_WithError(t *testing.T) {
	var f *os.File
	monkey.PatchInstanceMethod(reflect.TypeOf(f), "Close", func(_ *os.File) error {
		return errors.New("fake_file_close_error")
	})
	defer monkey.UnpatchAll()

	w := writer.TimeFileProvider("unused", "unused")
	newFile, err := w(&os.File{})
	assert.EqualError(t, err, "fake_file_close_error")
	assert.Nil(t, newFile)
}

func TestTimeFileProvider(t *testing.T) {
	f := os.File{}
	monkey.Patch(os.Create, func(name string) (*os.File, error) {
		assert.Equal(t, "fake_format_fake_time_format", name)
		return &f, nil
	})
	monkey.Patch(time.Now, func() time.Time { return time.Unix(513216000, 0) })
	defer monkey.UnpatchAll()

	w := writer.TimeFileProvider("fake_format_%s", "fake_time_format")
	newFile, err := w(nil)
	assert.Nil(t, err)
	assert.Equal(t, &f, newFile)
}

func TestLogFileProvider_RenameWithError(t *testing.T) {
	monkey.Patch(os.Rename, func(oldpath, newpath string) error {
		assert.Equal(t, "fake_format_fake_name", oldpath)
		assert.Equal(t, "fake_format_Apr  7 02:00:00", newpath)
		return errors.New("fake_rename_error")
	})
	monkey.Patch(time.Now, func() time.Time { return time.Unix(513216000, 0) })
	defer monkey.UnpatchAll()

	w := writer.LogFileProvider("fake_name", "fake_format_%s", time.Stamp)
	newFile, err := w(&os.File{})
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
		assert.Equal(t, "fake_format_Apr  7 02:00:00", newpath)
		return nil
	})
	monkey.Patch(time.Now, func() time.Time { return time.Unix(513216000, 0) })
	defer monkey.UnpatchAll()

	w := writer.LogFileProvider("fake_name", "fake_format_%s", time.Stamp)
	newFile, err := w(&os.File{})
	assert.EqualError(t, err, "fake_file_close_error")
	assert.Nil(t, newFile)
}

func TestTimeRotateWriter_StartWithError(t *testing.T) {
	mocksRotateWriter := mocks.RotateWriter{}
	mocksRotateWriter.On("Rotate").Return(func() error { return errors.New("fake_rotate_error") })
	tr := writer.TimeRotateWriter{RotateWriter: &mocksRotateWriter, Interval: 50 * time.Millisecond, PanicHandler: func(err error) {
		assert.EqualError(t, err, "fake_rotate_error")
	}}
	tr.Start()
	time.Sleep(100 * time.Millisecond)
}

func TestNewTimeRotateFileWriter_TimeFileProvider_Write(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time { return time.Unix(513216000, 0) })
	defer monkey.UnpatchAll()

	w, err := writer.NewTimeRotateFileWriter(writer.TimeFileProvider(os.TempDir()+"%s.log", time.Stamp), 1*time.Second)
	assert.Nil(t, err)

	n, err := w.Write([]byte("test"))
	assert.Nil(t, err)
	assert.Equal(t, 4, n)
	_, err = os.Stat(os.TempDir() + "Apr  7 02:00:00.log")
	assert.Nil(t, err)
}

func TestNewRotateFileWriter_TimeFileProvider_Rotate(t *testing.T) {
	i := 0
	monkey.Patch(time.Now, func() time.Time {
		r := []time.Time{
			time.Unix(513216000, 0),
			time.Unix(513217000, 0),
		}[i]
		i++
		return r
	})
	defer monkey.UnpatchAll()

	w, err := writer.NewTimeRotateFileWriter(writer.TimeFileProvider(os.TempDir()+"%s.log", time.Stamp), 1*time.Second)
	assert.Nil(t, err)

	n, err := w.Write([]byte("test"))
	assert.Nil(t, err)
	assert.Equal(t, 4, n)
	_, err = os.Stat(os.TempDir() + "Apr  7 02:00:00.log")
	assert.Nil(t, err)

	assert.Nil(t, w.Rotate())

	n, err = w.Write([]byte("test"))
	assert.Nil(t, err)
	assert.Equal(t, 4, n)
	_, err = os.Stat(os.TempDir() + "Apr  7 02:16:40.log")
	assert.Nil(t, err)

	//TODO test file content
}

func TestNewTimeRotateFileWriter_LogFileProvider_Write(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time { return time.Unix(513216000, 0) })
	defer monkey.UnpatchAll()

	w, err := writer.NewTimeRotateFileWriter(writer.LogFileProvider("test", os.TempDir()+"%s.log", time.Stamp), 1*time.Second)
	assert.Nil(t, err)

	n, err := w.Write([]byte("test"))
	assert.Nil(t, err)
	assert.Equal(t, 4, n)
	_, err = os.Stat(os.TempDir() + "Apr  7 02:00:00.log")
	assert.Nil(t, err)
}

func TestNewRotateFileWriter_LogFileProvider_Rotate(t *testing.T) {
	i := 0
	monkey.Patch(time.Now, func() time.Time {
		r := []time.Time{
			time.Unix(513216000, 0),
			time.Unix(513217000, 0),
		}[i]
		i++
		return r
	})
	defer monkey.UnpatchAll()

	w, err := writer.NewTimeRotateFileWriter(writer.LogFileProvider("test", os.TempDir()+"%s.log", time.Stamp), 1*time.Second)
	assert.Nil(t, err)

	n, err := w.Write([]byte("test"))
	assert.Nil(t, err)
	assert.Equal(t, 4, n)
	_, err = os.Stat(os.TempDir() + "Apr  7 02:00:00.log")
	assert.Nil(t, err)

	assert.Nil(t, w.Rotate())

	n, err = w.Write([]byte("test"))
	assert.Nil(t, err)
	assert.Equal(t, 4, n)
	_, err = os.Stat(os.TempDir() + "Apr  7 02:16:40.log")
	assert.Nil(t, err)

	//TODO test file content
}
