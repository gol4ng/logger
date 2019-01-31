package handler

import (
	"fmt"
	"io"

	"github.com/gol4ng/logger"
)

type Stream struct {
	writer    io.Writer
	formatter logger.FormatterInterface
}

func (s *Stream) Handle(e logger.Entry) error {
	_, err := fmt.Fprintln(s.writer, s.formatter.Format(e))

	return err
}

func NewStream(w io.Writer, f logger.FormatterInterface) *Stream {
	return &Stream{writer: w, formatter: f}
}
