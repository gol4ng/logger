package formatter

import (
	"strconv"
	"strings"

	"github.com/gol4ng/logger"
)

const (
	// JsonMessageKey represents the reserved "message" key of a JSON log entry.
	JsonMessageKey = "message"

	// JsonLevelKey represents the reserved "level" key of a JSON log entry.
	JsonLevelKey = "level"
)

type jsonEncodeFunc func(logger.Entry, *JSONOptions) ([]byte, error)

// JSONOptions stores all the options used by the JSON formatter.
type JSONOptions struct {
	flatten       bool
	levelAsString bool
}

// JSONOption represents a JSON option function available for this formatter.
type JSONOption func(*JSONOptions)

// WithJSONFlattenContext will flatten the context fields in the root JSON object.
func WithJSONFlattenContext() JSONOption {
	return func(o *JSONOptions) {
		o.flatten = true
	}
}

// WithJSONLevelAsString will render the level field as a string ("notice", "debug", ...)
// instead of integer value.
func WithJSONLevelAsString() JSONOption {
	return func(o *JSONOptions) {
		o.levelAsString = true
	}
}

// JSON formatter will transform a logger entry into JSON
// it takes an encode function that allows you to encode the data
//
// the encode function is useful if you do not use the default provided logger implementation
type JSON struct {
	encode  jsonEncodeFunc
	options *JSONOptions
}

// Format will return Entry as json
func (j *JSON) Format(entry logger.Entry) string {
	b, _ := j.encode(entry, j.options)
	return string(b)
}

// NewJSONEncoder will create a new JSON with default json encoder function
func NewJSONEncoder(options ...JSONOption) *JSON {
	return NewJSON(JSONEncoder, options...)
}

// NewJSON will create a new JSON with given json encoder
// it allow you tu use your own json encoder
func NewJSON(encode jsonEncodeFunc, options ...JSONOption) *JSON {
	opts := &JSONOptions{}

	for _, option := range options {
		option(opts)
	}

	return &JSON{
		encode:  encode,
		options: opts,
	}
}

// ContextToJSON will marshall the logger context into json
func ContextToJSON(context *logger.Context, builder *strings.Builder, options *JSONOptions) {
	if context == nil || len(*context) == 0 {
		return
	}

	if options.flatten {
		builder.WriteRune(',')
	} else {
		builder.WriteString(`,"Context":{`)
	}

	i := 0
	for name, field := range *context {
		if options.flatten {
			// Avoid collisions on reserved keys when flattening.
			switch name {
			case JsonLevelKey, JsonMessageKey:
				continue
			}
		}

		if i != 0 {
			builder.WriteRune(',')
		}
		builder.WriteRune('"')
		builder.WriteString(name)
		builder.WriteString(`":`)
		d, _ := field.MarshalJSON()
		builder.Write(d)
		i++
	}

	if !options.flatten {
		builder.WriteString(`}`)
	}
}

// EntryToJSON will marshall the logger Entry into json
func EntryToJSON(entry logger.Entry, builder *strings.Builder, options *JSONOptions) {
	builder.WriteRune('{')

	builder.WriteString(`"message":"`)
	builder.WriteString(entry.Message)
	builder.WriteRune('"')

	builder.WriteRune(',')
	builder.WriteString(`"level":"`)

	if options.levelAsString {
		builder.WriteString(entry.Level.String())
	} else {
		builder.WriteString(strconv.Itoa(int(entry.Level)))
	}

	builder.WriteRune('"')

	ContextToJSON(entry.Context, builder, options)

	builder.WriteRune('}')
}

// JSONEncoder will return Entry to json string
func JSONEncoder(entry logger.Entry, options *JSONOptions) ([]byte, error) {
	builder := &strings.Builder{}
	EntryToJSON(entry, builder, options)

	return []byte(builder.String()), nil
}
