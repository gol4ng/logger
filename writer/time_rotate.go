package writer

import (
	"os"
	"time"
)

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

func NewTimeRotateWriter(writer RotateWriter, interval time.Duration) *TimeRotateWriter {
	w := &TimeRotateWriter{RotateWriter: writer, Interval: interval}
	w.Start()
	return w
}

func NewTimeRotateFileWriter(fileProvider func(*os.File) (*os.File, error), interval time.Duration) (*TimeRotateWriter, error) {
	w, err := NewRotateFileWriter(fileProvider)
	return NewTimeRotateWriter(w, interval), err
}
