package logger

import "strings"

// Entry represents a log in its entirety
// it is composed of a level, a message and a context
type Entry struct {
	Message string
	Level   Level
	Context *Context
}

// String will return Entry as string
func (e *Entry) String() string {
	builder := &strings.Builder{}
	EntryToString(*e, builder)
	return builder.String()
}

// EntryToString will write entry as string in builder
func EntryToString(entry Entry, builder *strings.Builder) {
	builder.WriteString("<")
	builder.WriteString(entry.Level.String())
	builder.WriteString("> ")
	builder.WriteString(entry.Message)
	if entry.Context != nil {
		builder.WriteString(" [ ")
		builder.WriteString(entry.Context.String())
		builder.WriteString(" ]")
	}
}
