package formatter

import (
	"strings"

	"github.com/gol4ng/logger"
)

// DefaultFormatter is the default Entry formatter
type DefaultFormatter struct{}

// Format will return Entry as string
func (n *DefaultFormatter) Format(entry logger.Entry) string {
	builder := &strings.Builder{}
	builder.WriteString("<")
	builder.WriteString(entry.Level.String())
	builder.WriteString("> ")
	builder.WriteString(entry.Message)
	if entry.Context != nil {
		builder.WriteString(" ")
		ContextToJSON(entry.Context, builder)
	}

	return builder.String()
}

// NewDefaultFormatter will create a new DefaultFormatter
func NewDefaultFormatter() *DefaultFormatter {
	return &DefaultFormatter{}
}
