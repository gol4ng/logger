package formatter

import (
	"strings"

	"github.com/gol4ng/logger"
)

type DefaultFormatter struct{}
//format a logger entry as a string with a json representation of the context
func (n *DefaultFormatter) Format(entry logger.Entry) string {
	builder := &strings.Builder{}
	builder.WriteString("<")
	builder.WriteString(entry.Level.String())
	builder.WriteString("> ")
	builder.WriteString(entry.Message)
	if entry.Context != nil {
		builder.WriteString(" ")
		MarshalContextTo(entry.Context, builder)
	}

	return builder.String()
}
//constructor
func NewDefaultFormatter() *DefaultFormatter {
	return &DefaultFormatter{}
}
