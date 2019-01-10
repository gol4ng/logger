package formatter

import (
	"encoding/json"

	"github.com/instabledesign/logger"
)

type Json struct {
}

func (j *Json) Format(e logger.Entry) interface{} {
	// TODO USE ENCODER
	b, _ := json.Marshal(e)

	return string(b)
}

func NewJson() *Json {
	return &Json{}
}
