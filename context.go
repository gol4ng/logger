package logger

import (
	"strings"
)

// Context contain all contextual log data
// we advise you to choose your data wisely
// You should keep a reasonable quantity of data
type Context map[string]Field

// Merge and overwrite context data with the given one
func (c *Context) Merge(context Context) *Context {
	for _, field := range context {
		c.SetField(field)
	}
	return c
}

// SetField will add a new context field
func (c *Context) SetField(fields ...Field) *Context {
	for _, field := range fields {
		(*c)[field.Name] = field
	}
	return c
}

// Has will checks if the current context has a field name
func (c *Context) Has(name string) bool {
	_, ok := (*c)[name]
	return ok
}

// Get will return field with given name or the default value
func (c *Context) Get(name string, defaultField *Field) Field {
	if c.Has(name) {
		return (*c)[name]
	}
	return *defaultField
}

// Add will guess and add value to the context
func (c *Context) Add(name string, value interface{}) *Context {
	return c.SetField(Any(name, value))
}

// Skip will add skip field to context
func (c *Context) Skip(name string, value string) *Context {
	return c.SetField(Skip(name, value))
}

// Binary will add binary field to context
func (c *Context) Binary(name string, value []byte) *Context {
	return c.SetField(Binary(name, value))
}

// ByteString will add byteString field to context
func (c *Context) ByteString(name string, value []byte) *Context {
	return c.SetField(ByteString(name, value))
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

// Slice will return fields as *Fields
func (c *Context) Slice() *Fields {
	slice := Fields{}
	for _, field := range *c {
		slice = append(slice, field)
	}
	return &slice
}

// String will return fields as string
func (c *Context) String() string {
	if c == nil || len(*c) == 0 {
		return "<nil>"
	}
	builder := &strings.Builder{}
	c.stringTo(builder)
	return builder.String()
}

// GoString was called by fmt.Printf("%#v", context)
// fmt GoStringer interface
func (c *Context) GoString() string {
	builder := &strings.Builder{}
	builder.WriteString("logger.context[")
	c.stringTo(builder)
	builder.WriteString("]")
	return builder.String()
}

// Ctx will create a new context with given value
func Ctx(name string, value interface{}) *Context {
	return NewContext().Add(name, value)
}

// NewContext will create a new context with optional fields
func NewContext(fields ...Field) *Context {
	return (&Context{}).SetField(fields...)
}
