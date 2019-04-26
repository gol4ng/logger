package handler

import (
	"io"

	"github.com/gol4ng/logger"
)

type Stream struct {
	writer    io.Writer
	formatter logger.FormatterInterface
}

func (s *Stream) Handle(entry logger.Entry) error {
	_, err := s.writer.Write(append([]byte(s.formatter.Format(entry)), []byte("\n")...))

	return err
}

func NewStream(writer io.Writer, formatter logger.FormatterInterface) *Stream {
	return &Stream{writer: writer, formatter: formatter}
}
