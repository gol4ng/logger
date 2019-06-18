package handler

import (
	"io"

	"github.com/gol4ng/logger"
)

// handler that allows you to write logs into a an io.writer (stream)
func Stream(writer io.Writer, formatter logger.FormatterInterface) logger.HandlerInterface {
	return func(entry logger.Entry) error {
		_, err := writer.Write(append([]byte(formatter.Format(entry)), []byte("\n")...))
		return err
	}
}
