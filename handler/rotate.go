package handler

import (
	"time"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/writer"
)

func NewRotateStream(writer writer.RotateWriter, formatter logger.FormatterInterface, interval time.Duration) *Stream {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			if err := writer.Rotate(); err != nil {
				panic(err)
			}
		}
	}()
	return &Stream{writer: writer, formatter: formatter}
}

func NewTimeRotateFileStream(format string, timeFormat string, formatter logger.FormatterInterface, interval time.Duration) (*Stream, error) {
	w, err := writer.NewRotateFileWriter(writer.TimeFileProvider(format, timeFormat))

	return NewRotateStream(w, formatter, interval), err
}

func NewLogRotateFileStream(name string, format string, timeFormat string, formatter logger.FormatterInterface, interval time.Duration) (*Stream, error) {
	w, err := writer.NewRotateFileWriter(writer.LogFileProvider(name, format, timeFormat))

	return NewRotateStream(w, formatter, interval), err
}
