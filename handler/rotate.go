package handler

import (
	"time"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/writer"
	"github.com/gol4ng/logger/writer/provider"
)

// TimeRotateFileStream handler will create a TimeRotateWriter that creates a file rotator with a given logFileName format and a rotation interval
// it will create a new file each time rotate occurs
func TimeRotateFileStream(fileNameFormat string, timeFormat string, formatter logger.FormatterInterface, interval time.Duration) (logger.HandlerInterface, error) {
	return TimeRotateFromProvider(provider.TimeFileProvider(fileNameFormat, timeFormat), formatter, interval)
}

// LogRotateFileStream will create a TimeRotateFileStream that create file rotator with a given format name and rotation interval
// this handler will create one static file with the given name and backup file when rotate occurs over the time (with the given interval)
func LogRotateFileStream(fileNameFormat string, format string, timeFormat string, formatter logger.FormatterInterface, interval time.Duration) (logger.HandlerInterface, error) {
	return TimeRotateFromProvider(provider.LogFileProvider(fileNameFormat, format, timeFormat), formatter, interval)
}

// TimeRotateFromProvider will create a TimeRotateFromProvider that create io.Writer rotator with a given rotation interval
func TimeRotateFromProvider(provider writer.Provider, formatter logger.FormatterInterface, interval time.Duration) (logger.HandlerInterface, error) {
	w, err := writer.NewTimeRotateFromProvider(provider, interval)
	if err != nil {
		return nil, err
	}
	return Stream(w, formatter), nil
}
