package writer

import (
	"io"
	"sync"
)

// RotateIoWriter will rotate the io.writer when rotate was called
// it delegate the io.writer creation to the Provider
type RotateIoWriter struct {
	mutex          sync.Mutex
	writerProvider Provider
	writer         io.Writer
}

// Write will passthrough data to the underlaying io.writer
func (w *RotateIoWriter) Write(output []byte) (int, error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	return w.writer.Write(output)
}

// Rotate will ask the provider to change the underlaing io.writer with a new one
func (w *RotateIoWriter) Rotate() (err error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.writer, err = w.writerProvider(w.writer)
	return err
}

// NewRotateIoWriter will create a RotateIoWriter
func NewRotateIoWriter(provider Provider) (*RotateIoWriter, error) {
	writer, err := provider(nil)
	return &RotateIoWriter{writer: writer, writerProvider: provider}, err
}
