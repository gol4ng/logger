package formatter

import (
	"github.com/gol4ng/logger"
	"strings"
)

type DefaultFormatter struct{}

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

func NewDefaultFormatter() *DefaultFormatter {
	return &DefaultFormatter{}
}
