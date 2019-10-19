package writer_test

import (
	"bytes"
	"compress/flate"
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

	w := writer.NewCompressWriter(buffer)

	i, err := w.Write([]byte("fake_data"))
	assert.Equal(t, 99, i)
	assert.Nil(t, err)
}

func TestCompressWriter_Write_CompressGzip(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{})

	w := writer.NewCompressWriter(buffer, writer.CompressionType(writer.CompressGzip), writer.CompressionLevel(flate.BestSpeed))

	i, err := w.Write([]byte("fake_data"))
	assert.Equal(t, 37, i)
	assert.Nil(t, err)
}

func TestCompressWriter_Write_CompressZlib(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{})

	w := writer.NewCompressWriter(buffer, writer.CompressionType(writer.CompressZlib), writer.CompressionLevel(flate.BestSpeed))

	i, err := w.Write([]byte("fake_data"))
	assert.Equal(t, 25, i)
	assert.Nil(t, err)
}
