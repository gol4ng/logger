package handler

import (
	"time"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/writer"
	"github.com/gol4ng/logger/writer/provider"
)

func TimeRotateFileStream(format string, timeFormat string, formatter logger.FormatterInterface, interval time.Duration) (logger.HandlerInterface, error) {
	w, err := writer.NewTimeRotateFileWriter(provider.TimeFileProvider(format, timeFormat), interval)
	return Stream(w, formatter), err
}

func LogRotateFileStream(name string, format string, timeFormat string, formatter logger.FormatterInterface, interval time.Duration) (logger.HandlerInterface, error) {
	w, err := writer.NewTimeRotateFileWriter(provider.LogFileProvider(name, format, timeFormat), interval)
	return Stream(w, formatter), err
}
