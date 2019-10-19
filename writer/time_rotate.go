package writer

import (
	"time"
)

// TimeRotateWriter will call RotateWriter::Rotate func over the time with the given interval
// it delegate the io.writer creation to the Provider
type TimeRotateWriter struct {
	RotateWriter
	Interval     time.Duration
	PanicHandler func(error)
}

// Start will start listening the interval ticker
// rotate is gonna be called for each tick of the ticker
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
// it will call the given RotateWriter::Rotate func over the time with the given interval
func NewTimeRotateWriter(writer RotateWriter, interval time.Duration) *TimeRotateWriter {
	w := &TimeRotateWriter{RotateWriter: writer, Interval: interval}
	w.Start()
	return w
}

// NewTimeRotateFromProvider will create a TimeRotateFileWriter that rotate an io.writer over the time with the given interval
func NewTimeRotateFromProvider(provider Provider, interval time.Duration) (*TimeRotateWriter, error) {
	w, err := NewRotateIoWriter(provider)
	return NewTimeRotateWriter(w, interval), err
}
