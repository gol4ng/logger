package formatter

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/gol4ng/logger"
)

// Json formatter will transform a logger entry into JSON
// it takes an encode function that allows you to encode the data
//
// the encode function is useful if you do not use the default provided logger implementation
type Json struct {
	encode func(interface{}) ([]byte, error)
}

// Format will return Entry as json
func (j *Json) Format(entry logger.Entry) string {
	b, _ := j.encode(entry)
	return string(b)
}

// MarshalContextTo will marshall the logger context into json
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

// MarshalEntryTo will marshall the logger Entry into json
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

// NewJsonEncoder will create a new Json with default json encoder function
func NewJsonEncoder() *Json {
	return NewJson(entryJsonEncoder)
}

// NewJson will create a new Json with given json encoder
// it allow you tu use your own json encoder
func NewJson(encode func(interface{}) ([]byte, error)) *Json {
	return &Json{encode: encode}
}
