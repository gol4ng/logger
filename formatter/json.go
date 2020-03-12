package formatter

import (
	"strconv"
	"strings"

	"github.com/gol4ng/logger"
)

// JSON formatter will transform a logger entry into JSON
// it takes an encode function that allows you to encode the data
//
// the encode function is useful if you do not use the default provided logger implementation
type JSON struct {
	encode func(logger.Entry) ([]byte, error)
}

// Format will return Entry as json
func (j *JSON) Format(entry logger.Entry) string {
	b, _ := j.encode(entry)
	return string(b)
}

// NewJSONEncoder will create a new JSON with default json encoder function
func NewJSONEncoder() *JSON {
	return NewJSON(JSONEncoder)
}

// NewJSON will create a new JSON with given json encoder
// it allow you tu use your own json encoder
func NewJSON(encode func(logger.Entry) ([]byte, error)) *JSON {
	return &JSON{encode: encode}
}

// ContextToJSON will marshall the logger context into json
func ContextToJSON(context *logger.Context, builder *strings.Builder) {
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
			d, _ := field.MarshalJSON()
			builder.Write(d)
			i++
		}
		builder.WriteString("}")
	}
}

// EntryToJSON will marshall the logger Entry into json
func EntryToJSON(entry logger.Entry, builder *strings.Builder) {
	builder.WriteRune('{')

	builder.WriteString("\"Message\":\"")
	builder.WriteString(entry.Message)
	builder.WriteString("\"")

	builder.WriteRune(',')
	builder.WriteString("\"Level\":")
	builder.WriteString(strconv.Itoa(int(entry.Level)))

	builder.WriteRune(',')
	builder.WriteString("\"Context\":")

	ContextToJSON(entry.Context, builder)

	builder.WriteRune('}')
}

// JSONEncoder will return Entry to json string
func JSONEncoder(entry logger.Entry) ([]byte, error) {
	builder := &strings.Builder{}
	EntryToJSON(entry, builder)

	return []byte(builder.String()), nil
}
