package writer

import (
	"bytes"
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

	var compressWriter io.WriteCloser
	var err error
	// TODO create a pool buffer
	buf := bytes.NewBuffer(make([]byte, 0, 1024))

	switch w.compressionType {
	case CompressGzip:
		compressWriter, err = gzip.NewWriterLevel(buf, w.compressionLevel)
	case CompressZlib:
		compressWriter, err = zlib.NewWriterLevel(buf, w.compressionLevel)
	}
	if err != nil {
		return 0, err
	}

	if n, err := compressWriter.Write(p); err != nil {
		compressWriter.Close()
		return n, err
	}
	compressWriter.Close()

	return w.Writer.Write(buf.Bytes())
}

// NewCompressWriter will return a new compress writer
func NewCompressWriter(writer io.Writer, compressionType CompressType, compressionLevel int) *CompressWriter {
	return &CompressWriter{Writer: writer, compressionType: compressionType, compressionLevel: compressionLevel}
}
