package handler

import (
	"io"

	"github.com/gol4ng/logger"
)

// Stream will format and write Entry into io.writer
func Stream(writer io.Writer, formatter logger.FormatterInterface) logger.HandlerInterface {
	return func(entry logger.Entry) error {
		_, err := writer.Write(append([]byte(formatter.Format(entry)), []byte("\n")...))
		return err
	}
}
