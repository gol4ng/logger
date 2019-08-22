package writer_test

import (
	"io"
	"strings"
	"testing"

	"bou.ke/monkey"

	"github.com/gol4ng/logger/mocks"
	"github.com/gol4ng/logger/writer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewGelfChunkWriter_Simple(t *testing.T) {
	fakeData := []byte(strings.Repeat("A", 1000))

	writerMock := &mocks.Writer{}
	writerMock.On("Write", mock.MatchedBy(func(p []byte) bool {
		assert.Equal(t, fakeData, p)
		return true
	})).Return(33, nil)

	i, err := writer.NewGelfChunkWriter(writerMock).Write(fakeData)
	assert.Equal(t, 33, i)
	assert.Nil(t, err)
}

func TestNewGelfChunkWriter_Multiple(t *testing.T) {
	fakeData := []byte(strings.Repeat("A", 10000))
	fakeMessageId := []byte("BBBBBBBB")
	chunkedMagicBytes := []byte{0x1e, 0x0f}

	chunk1 := chunkedMagicBytes
	chunk1 = append(chunk1, fakeMessageId...)      // add message ID
	chunk1 = append(chunk1, []byte{0x00, 0x02}...) // add chunk number, total chunk
	prefixLen := len(chunk1)
	chunk1 = append(chunk1, []byte(strings.Repeat("A", 8192-prefixLen))...) // generate chunk data

	chunk2 := chunkedMagicBytes
	chunk2 = append(chunk2, fakeMessageId...)      // add message ID
	chunk2 = append(chunk2, []byte{0x01, 0x02}...) // add chunk number, total chunk
	prefixLen2 := len(chunk2)
	chunk2 = append(chunk2, []byte(strings.Repeat("A", len(fakeData)-8192+prefixLen2))...) // generate chunk data

	chunksExpected := [][]byte{
		chunk1,
		chunk2,
	}

	chunkedNb := 0
	writerMock := &mocks.Writer{}
	writerMock.On("Write", mock.MatchedBy(func(p []byte) bool {
		assert.Equal(t, chunksExpected[chunkedNb], p)
		chunkedNb++
		return true
	})).Return(33, nil)

	monkey.Patch(io.ReadFull, func(reader io.Reader, buf []byte) (int, error) {
		copy(buf, fakeMessageId)
		return len(fakeMessageId), nil
	})
	defer monkey.UnpatchAll()

	w := writer.NewGelfChunkWriter(writerMock)

	i, err := w.Write(fakeData)
	assert.Equal(t, 33*chunkedNb, i)
	assert.Nil(t, err)
}

