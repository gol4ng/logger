package logger

import (
	"github.com/valyala/bytebufferpool"
)

// Entry represents a log in its entirety
// it is composed of a level, a message and a context
type Entry struct {
	Message string
	Level   Level
	Context *Context
}

// String will return Entry as string
func (e *Entry) String() string {
	byteBuffer := bytebufferpool.Get()
	defer bytebufferpool.Put(byteBuffer)

	EntryToString(*e, byteBuffer)
	return byteBuffer.String()
}

// EntryToString will write entry as string in byteBuffer
func EntryToString(entry Entry, byteBuffer *bytebufferpool.ByteBuffer) {
	byteBuffer.WriteString(`<`)
	entry.Level.StringTo(byteBuffer)
	byteBuffer.WriteString(`> `)
	byteBuffer.WriteString(entry.Message)
	if entry.Context != nil {
		byteBuffer.WriteString(` [ `)
		entry.Context.StringTo(byteBuffer)
		byteBuffer.WriteString(` ]`)
	}
}
