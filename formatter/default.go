package formatter

import "github.com/instabledesign/logger"

type DefaultFormatter struct{}

func (n *DefaultFormatter) Format(e logger.Entry) string {
	return e.Level.String() + " " + e.Message
}

func NewDefaultFormatter() *DefaultFormatter {
	return &DefaultFormatter{}
}
