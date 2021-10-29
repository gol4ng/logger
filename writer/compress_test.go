package writer_test

import (
	"compress/flate"
	"testing"

	"github.com/gol4ng/logger/mocks"
	"github.com/gol4ng/logger/writer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCompressWriter_Write_CompressNone(t *testing.T) {
	buffer := &mocks.Writer{}
	buffer.On("Write", mock.MatchedBy(func(p []byte) bool {
		assert.Equal(t, []byte("fake_data"), p)
		return true
	})).Return(99, nil)

	w := writer.NewCompressWriter(buffer)

	i, err := w.Write([]byte("fake_data"))

	buffer.AssertExpectations(t)
	assert.Equal(t, 99, i)
	assert.Nil(t, err)
}

func TestCompressWriter_Write_CompressGzip(t *testing.T) {
	buffer := &mocks.Writer{}
	buffer.On("Write", mock.MatchedBy(func(p []byte) bool {
		assert.Equal(t, []byte{0x1f, 0x8b, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4, 0xff, 0x0, 0x9, 0x0, 0xf6, 0xff, 0x66, 0x61, 0x6b, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x1, 0x0, 0x0, 0xff, 0xff, 0x36, 0xbe, 0x4d, 0x32, 0x9, 0x0, 0x0, 0x0}, p)
		return true
	})).Return(37, nil)

	w := writer.NewCompressWriter(buffer, writer.CompressionType(writer.CompressGzip), writer.CompressionLevel(flate.BestSpeed))

	i, err := w.Write([]byte("fake_data"))

	buffer.AssertExpectations(t)
	assert.Equal(t, 37, i)
	assert.Nil(t, err)
}

func TestCompressWriter_Write_CompressZlib(t *testing.T) {
	writerMock := &mocks.Writer{}
	writerMock.On("Write", mock.MatchedBy(func(p []byte) bool {
		assert.Equal(t, []byte{0x78, 0x1, 0x0, 0x9, 0x0, 0xf6, 0xff, 0x66, 0x61, 0x6b, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x1, 0x0, 0x0, 0xff, 0xff, 0x11, 0xc9, 0x3, 0x91}, p)
		return true
	})).Return(25, nil)

	w := writer.NewCompressWriter(writerMock, writer.CompressionType(writer.CompressZlib), writer.CompressionLevel(flate.BestSpeed))

	i, err := w.Write([]byte("fake_data"))

	writerMock.AssertExpectations(t)
	assert.Equal(t, 25, i)
	assert.Nil(t, err)
}
