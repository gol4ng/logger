package writer

import (
	"io"
	"os"
	"sync"
)

type RotateWriter interface {
	io.Writer
	Rotate() error
}

type RotateFileWriter struct {
	mutex        sync.Mutex
	fileProvider FileProvider
	file         *os.File
}

func (w *RotateFileWriter) Write(output []byte) (int, error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	return w.file.Write(output)
}

func (w *RotateFileWriter) Rotate() (err error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.file, err = w.fileProvider(w.file)
	return err
}

func NewRotateFileWriter(fileProvider FileProvider) (*RotateFileWriter, error) {
	file, err := fileProvider(nil)
	return &RotateFileWriter{file: file, fileProvider: fileProvider}, err
}
