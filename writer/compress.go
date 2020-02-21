package writer

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"io"
	"sync"
)

// CompressType is the compress writer compression type
type CompressType int

// Compression type allowed
const (
	CompressGzip CompressType = iota
	CompressZlib
	CompressNone
)

// CompressOption is compress writer option setter function
type CompressOption func(*CompressWriter)

// CompressionType set compression type on compress writer
func CompressionType(t CompressType) CompressOption {
	return func(w *CompressWriter) {
		w.compressionType = t
	}
}

// CompressionLevel set compression level on compress writer
func CompressionLevel(level int) CompressOption {
	return func(w *CompressWriter) {
		w.compressionLevel = level
	}
}

// CompressWriter will decorate io.Writer in order to compress the data
type CompressWriter struct {
	io.Writer
	compressionType  CompressType
	compressionLevel int
}

// Write will compress data and pass it to the underlying io.Writer
func (w *CompressWriter) Write(p []byte) (int, error) {
	if w.compressionType == CompressNone {
		return w.Writer.Write(p)
	}

	var compressWriter io.WriteCloser
	var err error
	buf := newBuffer()
	defer bufPool.Put(buf)

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
		_ = compressWriter.Close()
		return n, err
	}
	_ = compressWriter.Close()

	return w.Writer.Write(buf.Bytes())
}

// NewCompressWriter will return a new compress writer
func NewCompressWriter(writer io.Writer, options ...CompressOption) *CompressWriter {
	w := &CompressWriter{Writer: writer, compressionType: CompressNone}
	for _, option := range options {
		option(w)
	}
	return w
}

// 1k bytes buffer by default
var bufPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 1024))
	},
}

func newBuffer() *bytes.Buffer {
	b := bufPool.Get().(*bytes.Buffer)
	if b != nil {
		b.Reset()
		return b
	}
	return bytes.NewBuffer(nil)
}
