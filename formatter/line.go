package formatter

import (
	"fmt"

	"github.com/gol4ng/logger"
)

// Line formatter will transform log Entry into string
type Line struct {
	format string
}

// Format will return Entry as string
// typically used for stdout output
func (l *Line) Format(entry logger.Entry) string {
	return fmt.Sprintf(l.format, entry.Message, entry.Level, entry.Context)
}

// NewLine will create a new Line with format (fmt)
func NewLine(format string) *Line {
	return &Line{format}
}
