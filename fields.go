package logger

import (
	"github.com/valyala/bytebufferpool"
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

func (f *Fields) StringTo(byteBuffer *bytebufferpool.ByteBuffer) {
	if len(*f) == 0 {
		byteBuffer.WriteString(" ")
		return
	}
	first := true
	for _, field := range *f {
		if !first {
			byteBuffer.WriteString(" ")
		}
		byteBuffer.WriteString("<")
		byteBuffer.WriteString(field.Name)
		byteBuffer.WriteString(":")
		byteBuffer.WriteString(field.String())
		byteBuffer.WriteString(">")
		first = false
	}
	return
}

// String will return Fields as string
func (f *Fields) String() string {
	byteBuffer := bytebufferpool.Get()
	defer bytebufferpool.Put(byteBuffer)

	f.StringTo(byteBuffer)
	return byteBuffer.String()
}

// GoString was called by fmt.Printf("%#v", Fields)
// fmt GoStringer interface
func (f *Fields) GoString() string {
	byteBuffer := bytebufferpool.Get()
	defer bytebufferpool.Put(byteBuffer)

	byteBuffer.WriteString("logger.Fields[")
	f.StringTo(byteBuffer)
	byteBuffer.WriteString("]")
	return byteBuffer.String()
}

// NewFields will create a Fields collection with initial value
func NewFields(name string, value interface{}) *Fields {
	return (&Fields{}).Add(name, value)
}
