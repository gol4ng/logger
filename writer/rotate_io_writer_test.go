package writer_test

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/gol4ng/logger/mocks"
	"github.com/gol4ng/logger/writer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRotateIoWriter_Write(t *testing.T) {
	writerMock := &mocks.Writer{}
	writerMock.On("Write", mock.MatchedBy(func(p []byte) bool {
		assert.Equal(t, []byte("fake_data"), p)
		return true
	})).Return(99, nil)

	w, err := writer.NewRotateIoWriter(func(io.Writer) (io.Writer, error) {
		return writerMock, nil
	})
	assert.Nil(t, err)

	i, err := w.Write([]byte("fake_data"))

	writerMock.AssertExpectations(t)
	assert.Equal(t, 99, i)
	assert.Nil(t, err)
}

func TestRotateIoWriter_Rotate(t *testing.T) {
	writer1Mock := &mocks.Writer{}
	writer1Mock.On("Write", mock.MatchedBy(func(p []byte) bool {
		assert.Equal(t, []byte("fake_data_1"), p)
		return true
	})).Return(99, nil)

	writer2Mock := &mocks.Writer{}
	writer2Mock.On("Write", mock.MatchedBy(func(p []byte) bool {
		assert.Equal(t, []byte("fake_data_2"), p)
		return true
	})).Return(33, nil)

	w, err := writer.NewRotateIoWriter(func(w io.Writer) (io.Writer, error) {
		if w == nil {
			return writer1Mock, nil
		}
		return writer2Mock, nil
	})
	assert.Nil(t, err)

	i1, err1 := w.Write([]byte("fake_data_1"))

	writer1Mock.AssertExpectations(t)
	assert.Equal(t, 99, i1)
	assert.Nil(t, err1)

	err = w.Rotate()

	assert.Nil(t, err)

	i2, err2 := w.Write([]byte("fake_data_2"))

	writer1Mock.AssertExpectations(t)
	assert.Equal(t, 33, i2)
	assert.Nil(t, err2)
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
