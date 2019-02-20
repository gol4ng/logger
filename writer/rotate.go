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
	fileProvider func(*os.File) (*os.File, error)
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

func NewRotateFileWriter(fileProvider func(*os.File) (*os.File, error)) (*RotateFileWriter, error) {
	file, err := fileProvider(nil)
	fw := &RotateFileWriter{file: file, fileProvider: fileProvider}
	return fw, err
}
