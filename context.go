package logger

import (
	"strings"
)

// represents a log context
type Context map[string]Field

// merge a context with another one
func (c *Context) Merge(context Context) *Context {
	for name, field := range context {
		c.Set(name, field)
	}
	return c
}

// add a new context field
func (c *Context) Set(name string, value Field) *Context {
	(*c)[name] = value
	return c
}

// checks if the current context has a context entry having a given name
func (c *Context) Has(name string) bool {
	_, ok := (*c)[name]
	return ok
}

// FIXME useless : remove field param
func (c *Context) Get(name string, field *Field) Field {
	if c.Has(name) {
		return (*c)[name]
	}
	return *field
}

// helper that adds a context entry without specifying field type
func (c *Context) Add(name string, value interface{}) *Context {
	return c.Set(name, Any(value))
}

func (c *Context) Skip(name string, value string) *Context {
	return c.Set(name, Skip(value))
}

func (c *Context) Binary(name string, value []byte) *Context {
	return c.Set(name, Binary(value))
}

func (c *Context) ByteString(name string, value []byte) *Context {
	return c.Set(name, ByteString(value))
}

func (c *Context) stringTo(builder *strings.Builder) *Context {
	if len(*c) == 0 {
		builder.WriteString("nil")
		return c
	}
	i := 0
	for name, field := range *c {
		if i != 0 {
			builder.WriteString(" ")
		}
		builder.WriteString("<")
		builder.WriteString(name)
		builder.WriteString(":")
		builder.WriteString(field.String())
		builder.WriteString(">")
		i++
	}
	return c
}

// stringify a context
func (c *Context) String() string {
	if len(*c) == 0 {
		return "<nil>"
	}
	builder := &strings.Builder{}
	c.stringTo(builder)
	return builder.String()
}

// TODO MOVE context serialization
// fmt GoStringer
// usefull when you fmt.Printf("%#v", GoStringer)
func (c *Context) GoString() string {
	builder := &strings.Builder{}
	builder.WriteString("logger.context[")
	c.stringTo(builder)
	builder.WriteString("]")
	return builder.String()
}

// create a new context with some context entry
func Ctx(name string, value interface{}) *Context {
	return NewContext().Add(name, value)
}

// create a new context
func NewContext() *Context {
	return &Context{}
}
