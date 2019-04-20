package formatter

import (
	"encoding/json"
	"github.com/gol4ng/logger"
)

type Json struct {
	encode func(v interface{}) ([]byte, error)
}

func (j *Json) Format(e logger.Entry) string {
	b, _ := j.encode(e)

	return string(b)
}

func NewJsonMarshall() *Json {
	return &Json{encode: json.Marshal}
}

func NewJson(encode func(v interface{}) (bytes []byte, e error)) *Json {
	return &Json{encode: encode}
}
