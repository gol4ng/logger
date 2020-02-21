package writer

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
)

const (
	// ChunkSize sets to maximal chunk size in bytes.
	chunkSize = 8192
	// ChunkDataSize is ChunkSize minus header's size.
	chunkDataSize = chunkSize - 12
	// MaxChunkCount maximal chunk per message count.
	maxChunkCount = 128
)

var (
	// ChunkedMagicBytes chunked message magic bytes.
	// See http://docs.graylog.org/en/2.4/pages/gelf.html.
	chunkedMagicBytes = []byte{0x1e, 0x0f}
)

// GelfChunkWriter will decorate io.Writer in order to split the data into chunk
type GelfChunkWriter struct {
	writer io.Writer
}

// Writer will split the data and pass it to the underlying io.Writer
func (w *GelfChunkWriter) Write(p []byte) (int, error) {
	lenB := len(p)
	chunkedNb := 1
	if lenB > chunkSize {
		chunkedNb = lenB/chunkDataSize + 1
	}

	if chunkedNb > maxChunkCount {
		return 0, fmt.Errorf("chunk count should be %d or less, %d given", maxChunkCount, chunkedNb)
	}

	if chunkedNb > 1 {
		messageID := make([]byte, 8)
		if n, err := io.ReadFull(rand.Reader, messageID); err != nil {
			return 0, fmt.Errorf("message id can not be generated : %w", err)
		} else if n != 8 {
			return 0, fmt.Errorf("message id can not be generated : cannot write 8 bytes (only %d)", n)
		}
		buffer := bytes.NewBuffer(make([]byte, 0, chunkSize))
		bytesLeft := lenB
		writtenBytes := 0
		for i := 0; i < chunkedNb; i++ {
			off := i * chunkDataSize
			chunkLen := chunkDataSize
			if chunkLen > bytesLeft {
				chunkLen = bytesLeft
			}

			buffer.Reset()
			buffer.Write(chunkedMagicBytes)
			buffer.Write(messageID)
			buffer.WriteByte(uint8(i))
			buffer.WriteByte(uint8(chunkedNb))
			buffer.Write(p[off : off+chunkLen])

			n, err := w.writer.Write(buffer.Bytes())
			writtenBytes += n
			if err != nil {
				return writtenBytes, err
			}
			bytesLeft -= chunkLen
		}
		return writtenBytes, nil
	}
	return w.writer.Write(p)
}

// NewGelfChunkWriter will return a new GelfChunkWriter
func NewGelfChunkWriter(writer io.Writer) *GelfChunkWriter {
	return &GelfChunkWriter{writer: writer}
}
