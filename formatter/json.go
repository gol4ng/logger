package formatter

import (
	"github.com/gol4ng/logger"
	"github.com/valyala/bytebufferpool"
	"strconv"
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
func ContextToJSON(context *logger.Context, byteBuffer *bytebufferpool.ByteBuffer) {
	if context == nil || len(*context) == 0 {
		byteBuffer.WriteString("null")
	} else {
		byteBuffer.WriteString(`{`)
		first := true
		for name, field := range *context {
			if !first {
				byteBuffer.WriteString(`,`)
			}
			byteBuffer.WriteString(`"`)
			byteBuffer.WriteString(name)
			byteBuffer.WriteString(`":`)
			d, _ := field.MarshalJSON()
			byteBuffer.Write(d)
			first = false
		}
		byteBuffer.WriteString(`}`)
	}
}

// EntryToJSON will marshall the logger Entry into json
func EntryToJSON(entry logger.Entry, byteBuffer *bytebufferpool.ByteBuffer) {
	byteBuffer.WriteString(`{"Message":"`)
	byteBuffer.WriteString(entry.Message)
	byteBuffer.WriteString(`","Level":`)
	byteBuffer.WriteString(strconv.Itoa(int(entry.Level)))
	byteBuffer.WriteString(`,"Context":`)
	ContextToJSON(entry.Context, byteBuffer)
	byteBuffer.WriteString(`}`)
}

// JSONEncoder will return Entry to json string
func JSONEncoder(entry logger.Entry) ([]byte, error) {
	byteBuffer := bytebufferpool.Get()
	defer bytebufferpool.Put(byteBuffer)

	EntryToJSON(entry, byteBuffer)

	return byteBuffer.Bytes(), nil
}
