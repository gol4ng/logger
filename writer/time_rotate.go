package writer

import (
	"time"
)

// TimeRotateWriter will rotate the io.writer over the time with the given interval
// it delegate the io.writer creation to the Provider
type TimeRotateWriter struct {
	RotateWriter
	Interval     time.Duration
	PanicHandler func(error)
}

// Start will start listening the interval ticker
// rotate gonna be called for each tick of the ticker
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

// NewTimeRotateWriter will create a new TimeRotateWriter
func NewTimeRotateWriter(writer RotateWriter, interval time.Duration) *TimeRotateWriter {
	w := &TimeRotateWriter{RotateWriter: writer, Interval: interval}
	w.Start()
	return w
}

// NewTimeRotateFileWriter will create a new TimeRotateFileWriter
func NewTimeRotateFileWriter(provider Provider, interval time.Duration) (*TimeRotateWriter, error) {
	w, err := NewRotateIoWriter(provider)
	return NewTimeRotateWriter(w, interval), err
}
