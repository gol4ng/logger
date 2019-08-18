package writer

import (
	"compress/gzip"
	"compress/zlib"
	"io"
)

// CompressType is the compress writer compression type
type CompressType int

// Compression type allowed
const (
	CompressGzip CompressType = iota
	CompressZlib
	CompressNone
)

// CompressWriter will decorate io.Writer in order to compress the data
type CompressWriter struct {
	io.Writer
	compressionType  CompressType
	compressionLevel int
}

// Write will compress data and pass it to the underlaying io.Writer
func (w *CompressWriter) Write(p []byte) (int, error) {
	if w.compressionType == CompressNone {
		return w.Writer.Write(p)
	}

	var writer io.WriteCloser
	var err error

	switch w.compressionType {
	case CompressGzip:
		writer, err = gzip.NewWriterLevel(w.Writer, w.compressionLevel)
	case CompressZlib:
		writer, err = zlib.NewWriterLevel(w.Writer, w.compressionLevel)
	}
	if err != nil {
		return 0, err
	}

	n, err := writer.Write(p)
	if err != nil {
		return n, err
	}
	return n, writer.Close()
}

// NewCompressWriter will return a new compress writer
func NewCompressWriter(writer io.Writer, compressionType CompressType, compressionLevel int) *CompressWriter {
	return &CompressWriter{Writer: writer, compressionType: compressionType, compressionLevel: compressionLevel}
}
