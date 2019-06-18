package formatter

import (
	"fmt"

	"github.com/gol4ng/logger"
)
// allows you to format a logger entry into a human readable string
type Line struct {
	format string
}
// format a logger entry into a human readable string
// typically used for stdout output
func (l *Line) Format(entry logger.Entry) string {
	return fmt.Sprintf(l.format, entry.Message, entry.Level, entry.Context)
}
// create new line formatter instance
func NewLine(format string) *Line {
	return &Line{format}
}
