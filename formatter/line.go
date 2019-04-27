package formatter

import (
	"fmt"

	"github.com/gol4ng/logger"
)

type Line struct {
	format string
}

func (l *Line) Format(entry logger.Entry) string {
	return fmt.Sprintf(l.format, entry.Message, entry.Level, entry.Context)
}

func NewLine(format string) *Line {
	return &Line{format}
}
