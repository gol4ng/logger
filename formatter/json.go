package formatter

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/gol4ng/logger"
)

// json formatter will transform a logger entry into JSON
// it takes an encode function that allows you to encode the data
//
// the encode function is useful if you do not use the default provided logger implementation
type Json struct {
	encode func(interface{}) ([]byte, error)
}

// transforms a log entry into json
func (j *Json) Format(entry logger.Entry) string {
	b, _ := j.encode(entry)
	return string(b)
}

// marshall the logger context into json
func MarshalContextTo(context *logger.Context, builder *strings.Builder) {
	if context == nil || len(*context) == 0 {
		builder.WriteString("null")
	} else {
		builder.WriteString("{")
		i := 0
		for name, field := range *context {
			if i != 0 {
				builder.WriteRune(',')
			}
			builder.WriteRune('"')
			builder.WriteString(name)
			builder.WriteString("\":")
			d, _ := json.Marshal(field.Value)
			builder.WriteString(string(d))
			i++
		}
		builder.WriteString("}")
	}
}

// marshall the logger entry into json
func MarshalEntryTo(entry logger.Entry, builder *strings.Builder) {
	builder.WriteRune('{')

	builder.WriteString("\"Message\":\"")
	builder.WriteString(entry.Message)
	builder.WriteString("\"")

	builder.WriteRune(',')
	builder.WriteString("\"Level\":")
	builder.WriteString(strconv.Itoa(int(entry.Level)))

	builder.WriteRune(',')
	builder.WriteString("\"Context\":")

	MarshalContextTo(entry.Context, builder)

	builder.WriteRune('}')
}

func entryJsonEncoder(value interface{}) ([]byte, error) {
	data := &strings.Builder{}
	MarshalEntryTo(value.(logger.Entry), data)

	return []byte(data.String()), nil
}

// json encoder that encodes a logger entry into json
func NewJsonEncoder() *Json {
	return NewJson(entryJsonEncoder)
}

// json formatter constructor
func NewJson(encode func(interface{}) ([]byte, error)) *Json {
	return &Json{encode: encode}
}
