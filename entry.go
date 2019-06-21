package logger

import "strings"

// Entry represents a log in its entirety
// it is composed of a level, a message and context
type Entry struct {
	Message string
	Level   Level
	Context *Context
}

// String will return Entry as string
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
