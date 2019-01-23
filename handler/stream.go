package handler

import (
	"fmt"
	"io"
	"os"
	"syscall"

	"github.com/instabledesign/logger"
)

type Stream struct {
	writer    io.Writer
	formatter logger.FormatterInterface
}

func (s *Stream) Handle(e logger.Entry) error {
	_, err := fmt.Fprintln(s.writer, s.formatter.Format(e))

	return err
}

func NewNopStream() *Stream {
	return &Stream{writer: os.NewFile(uintptr(syscall.Stderr), "/dev/null"), formatter: &logger.NopFormatter{}}
}

func NewStream(w io.Writer, f logger.FormatterInterface) *Stream {
	return &Stream{writer: w, formatter: f}
}
