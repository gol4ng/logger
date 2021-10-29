package formatter

import (
	"fmt"

	"github.com/gol4ng/logger"
)

const (
	LineFormatDefault = "%s %s %s"             // test message warning <my_key:my_value>
	LineFormatLevelInt = "%s %d %s"            // test message 4 <my_key:my_value>
	LineFormatLevelFirst = "%[2]s %[1]s %[3]s" // warning test message <my_key:my_value>
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
