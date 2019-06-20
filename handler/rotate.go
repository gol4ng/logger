package handler

import (
	"time"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/writer"
	"github.com/gol4ng/logger/writer/provider"
)

// TimeRotateFileStream will create a TimeRotateFileStream that create file rotator with given format name and rotation interval
// it will create a new file each time rotate occurred
func TimeRotateFileStream(format string, timeFormat string, formatter logger.FormatterInterface, interval time.Duration) (logger.HandlerInterface, error) {
	w, err := writer.NewTimeRotateFileWriter(provider.TimeFileProvider(format, timeFormat), interval)
	if err != nil {
		return nil, err
	}
	return Stream(w, formatter), nil
}

// LogRotateFileStream will create a TimeRotateFileStream that create file rotator with given format name and rotation interval
// this handler will create one static file with the given name and backup file when rotate occurred over the time (with given interval)
func LogRotateFileStream(name string, format string, timeFormat string, formatter logger.FormatterInterface, interval time.Duration) (logger.HandlerInterface, error) {
	w, err := writer.NewTimeRotateFileWriter(provider.LogFileProvider(name, format, timeFormat), interval)
	if err != nil {
		return nil, err
	}
	return Stream(w, formatter), nil
}
