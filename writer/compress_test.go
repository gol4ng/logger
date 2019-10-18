package writer_test

import (
	"bytes"
	"reflect"
	"testing"

	"bou.ke/monkey"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/logger/writer"
)

func TestCompressWriter_Write_CompressNone(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{})
	monkey.PatchInstanceMethod(reflect.TypeOf(buffer), "Write", func(_ *bytes.Buffer, b []byte) (n int, err error) {
		assert.Equal(t, []byte("fake_data"), b)
		return 99, nil
	})
	defer monkey.UnpatchAll()

	w := writer.NewCompressWriter(buffer, writer.CompressNone, 46)

	i, err := w.Write([]byte("fake_data"))
	assert.Equal(t, 99, i)
	assert.Nil(t, err)
}

func TestCompressWriter_Write_CompressGzip(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{})

	w := writer.NewCompressWriter(buffer, writer.CompressGzip, 2)

	i, err := w.Write([]byte("fake_data"))
	assert.Equal(t, 33, i)
	assert.Nil(t, err)
}

func TestCompressWriter_Write_CompressZlib(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{})

	w := writer.NewCompressWriter(buffer, writer.CompressZlib, 2)

	i, err := w.Write([]byte("fake_data"))
	assert.Equal(t, 21, i)
	assert.Nil(t, err)
}
