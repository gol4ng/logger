package logger

import "strings"

// logger.Entry represents a log in its entirety
// it is composed of a level, a message and context
type Entry struct {
	Message string
	Level   Level
	Context *Context
}

// stringify a log entry
// typically used when logging with line formatter
func (e *Entry) String() string {
	builder := &strings.Builder{}
	builder.WriteString("<")
	builder.WriteString(e.Level.String())
	builder.WriteString("> ")
	builder.WriteString(e.Message)
	if e.Context != nil {
		builder.WriteString(" [ ")
		builder.WriteString(e.Context.String())
		builder.WriteString(" ]")
	}
	return builder.String()
}
