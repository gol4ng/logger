package writer

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

func TimeFileProvider(format string, timeFormat string) func(f *os.File) (*os.File, error) {
	return func(f *os.File) (*os.File, error) {
		if f != nil {
			err := f.Close()
			if err != nil {
				return nil, err
			}
		}
		return os.Create(fmt.Sprintf(format, time.Now().Format(timeFormat)))
	}
}

func LogFileProvider(name string, format string, timeFormat string) func(f *os.File) (*os.File, error) {
	basePath := fmt.Sprintf(format, name)
	return func(f *os.File) (*os.File, error) {
		if f != nil {
			err := os.Rename(basePath, fmt.Sprintf(format, time.Now().Format(timeFormat)))
			if err != nil {
				return nil, err
			}
			err = f.Close()
			if err != nil {
				return nil, err
			}
		}
		return os.Create(basePath)
	}
}

type TimeRotateWriter struct {
	RotateWriter
	Interval     time.Duration
	PanicHandler func(error)
}

func (t *TimeRotateWriter) Start() {
	ticker := time.NewTicker(t.Interval)
	go func() {
		for range ticker.C {
			if err := t.Rotate(); err != nil {
				if t.PanicHandler != nil {
					t.PanicHandler(err)
				}
			}
		}
	}()
}

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

func NewTimeRotateWriter(writer RotateWriter, interval time.Duration) *TimeRotateWriter {
	w := &TimeRotateWriter{RotateWriter: writer, Interval: interval}
	w.Start()
	return w
}

func NewTimeRotateFileWriter(fileProvider func(*os.File) (*os.File, error), interval time.Duration) (*TimeRotateWriter, error) {
	w, err := NewRotateFileWriter(fileProvider)
	return NewTimeRotateWriter(w, interval), err
}
