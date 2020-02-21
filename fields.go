package logger

import (
	"strings"
)

// Fields was slice format of Fields
type Fields []Field

// SetField will add a new fields field
func (f *Fields) SetField(fields ...Field) *Fields {
	*f = append(*f, fields...)
	return f
}

// Add will guess and add value to the fields
func (f *Fields) Add(name string, value interface{}) *Fields {
	return f.SetField(Any(name, value))
}

// Skip will add skip field to fields
func (f *Fields) Skip(name string, value string) *Fields {
	return f.SetField(Skip(name, value))
}

// Binary will add binary field to fields
func (f *Fields) Binary(name string, value []byte) *Fields {
	return f.SetField(Binary(name, value))
}

// ByteString will add byteString field to fields
func (f *Fields) ByteString(name string, value []byte) *Fields {
	return f.SetField(ByteString(name, value))
}

func (f *Fields) stringTo(builder *strings.Builder) *Fields {
	if len(*f) == 0 {
		builder.WriteString(" ")
		return f
	}
	i := 0
	for _, field := range *f {
		if i != 0 {
			builder.WriteString(" ")
		}
		builder.WriteString("<")
		builder.WriteString(field.Name)
		builder.WriteString(":")
		builder.WriteString(field.String())
		builder.WriteString(">")
		i++
	}
	return f

}

// String will return Fields as string
func (f *Fields) String() string {
	builder := &strings.Builder{}
	f.stringTo(builder)
	return builder.String()
}

// GoString was called by fmt.Printf("%#v", Fields)
// fmt GoStringer interface
func (f *Fields) GoString() string {
	builder := &strings.Builder{}
	builder.WriteString("logger.Fields[")
	f.stringTo(builder)
	builder.WriteString("]")
	return builder.String()
}

// NewFields will create a Fields collection with initial value
func NewFields(name string, value interface{}) *Fields {
	return (&Fields{}).Add(name, value)
}
