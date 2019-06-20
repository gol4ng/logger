package writer

import (
	"io"
	"sync"
)

type RotateWriter interface {
	io.Writer
	Rotate() error
}

type RotateIoWriter struct {
	mutex          sync.Mutex
	writerProvider Provider
	writer         io.Writer
}

func (w *RotateIoWriter) Write(output []byte) (int, error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	return w.writer.Write(output)
}

func (w *RotateIoWriter) Rotate() (err error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.writer, err = w.writerProvider(w.writer)
	return err
}

func NewRotateIoWriter(provider Provider) (*RotateIoWriter, error) {
	writer, err := provider(nil)
	return &RotateIoWriter{writer: writer, writerProvider: provider}, err
}
