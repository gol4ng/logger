package formatter

import (
	"fmt"

	"github.com/instabledesign/logger"
)

type Line struct {
	format string
}

func (l *Line) Format(e logger.Entry) interface{} {
	return fmt.Sprintf(l.format, e.Message, e.Level, e.Context)
}

func NewLine(format string) *Line {
	return &Line{format}
}
