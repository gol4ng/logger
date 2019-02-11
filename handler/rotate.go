package handler

import (
	"time"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/writer"
)

func NewTimeRotateFileStream(format string, timeFormat string, formatter logger.FormatterInterface, interval time.Duration) (*Stream, error) {
	w, err := writer.NewTimeRotateFileWriter(writer.TimeFileProvider(format, timeFormat), interval)
	return &Stream{writer: w, formatter: formatter}, err
}

func NewLogRotateFileStream(name string, format string, timeFormat string, formatter logger.FormatterInterface, interval time.Duration) (*Stream, error) {
	w, err := writer.NewTimeRotateFileWriter(writer.LogFileProvider(name, format, timeFormat), interval)
	return &Stream{writer: w, formatter: formatter}, err
}
