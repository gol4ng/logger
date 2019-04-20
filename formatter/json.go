package formatter

import (
	"encoding/json"
	"github.com/gol4ng/logger"
	"strconv"
	"strings"
)

type Json struct {
	encode func(v interface{}) ([]byte, error)
}

func (j *Json) Format(e logger.Entry) string {
	b, _ := j.encode(e)

	return string(b)
}

func entryJsonEncoder(v interface{}) ([]byte, error) {
	entry := v.(logger.Entry)
	data := strings.Builder{}
	data.WriteRune('{')

	data.WriteString("\"Message\":\"")
	data.WriteString(entry.Message)
	data.WriteString("\"")

	data.WriteRune(',')
	data.WriteString("\"Level\":")
	data.WriteString(strconv.Itoa(int(entry.Level)))

	data.WriteRune(',')
	data.WriteString("\"Context\":")

	if entry.Context == nil || len(*(entry.Context)) == 0 {
		data.WriteString("null")
	} else {
		data.WriteString("{")
		i := 0
		for name, field := range *entry.Context {
			if i != 0 {
				data.WriteRune(',')
			}
			data.WriteRune('"')
			data.WriteString(name)
			data.WriteString("\":")
			d, _ := json.Marshal(field.Value)
			data.WriteString(string(d))
			i++
		}
		data.WriteString("}")
	}

	data.WriteRune('}')

	return []byte(data.String()), nil
}

func NewJsonEncoder() *Json {
	return NewJson(entryJsonEncoder)
}

func NewJson(encode func(v interface{}) (bytes []byte, e error)) *Json {
	return &Json{encode: encode}
}
