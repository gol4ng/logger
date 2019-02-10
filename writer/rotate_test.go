package writer_test

import (
	"os"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/logger/writer"
)

func TestNewTimeRotateFileWriter_Write(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time { return time.Unix(513216000, 0) })
	defer monkey.UnpatchAll()

	w, err := writer.NewRotateFileWriter(writer.TimeFileProvider(os.TempDir()+"%s.log", time.Stamp))
	assert.Nil(t, err)

	n, err := w.Write([]byte("test"))
	assert.Nil(t, err)
	assert.Equal(t, 4, n)
	_, err = os.Stat(os.TempDir() + "Apr  7 02:00:00.log")
	assert.Nil(t, err)
}

func TestNewRotateFileWriter_Rotate(t *testing.T) {
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

	w, err := writer.NewRotateFileWriter(writer.TimeFileProvider(os.TempDir()+"%s.log", time.Stamp))
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

	w, err := writer.NewRotateFileWriter(writer.LogFileProvider("test", os.TempDir()+"%s.log", time.Stamp))
	assert.Nil(t, err)

	n, err := w.Write([]byte("first"))
	assert.Nil(t, err)
	assert.Equal(t, 5, n)
	if _, err := os.Stat(os.TempDir() + "test.log"); os.IsNotExist(err) {
		t.Error("file \"/tmp/test.log\" must exist")
	}

	assert.Nil(t, w.Rotate())
	if _, err := os.Stat(os.TempDir() + "Apr  7 02:00:00.log"); os.IsNotExist(err) {
		t.Error("file \"/tmp/test.log\" must exist")
	}
	if _, err := os.Stat(os.TempDir() + "test.log"); os.IsNotExist(err) {
		t.Error("file \"/tmp/test.log\" must exist")
	}

	n, err = w.Write([]byte("second"))
	assert.Nil(t, err)
	assert.Equal(t, 6, n)

	//TODO test file content
}
