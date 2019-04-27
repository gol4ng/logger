package logger

import "strings"

type Entry struct {
	Message string
	Level   Level
	Context *Context
}

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
