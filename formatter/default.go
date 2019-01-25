package formatter

import "github.com/gol4ng/logger"

type DefaultFormatter struct{}

func (n *DefaultFormatter) Format(e logger.Entry) string {
	return e.Level.String() + " " + e.Message
}

func NewDefaultFormatter() *DefaultFormatter {
	return &DefaultFormatter{}
}
