package formatter

import (
	"encoding/json"

	"github.com/gol4ng/logger"
)

type Json struct {
}

func (j *Json) Format(e logger.Entry) string {
	// TODO USE ENCODER
	b, _ := json.Marshal(e)

	return string(b)
}

func NewJson() *Json {
	return &Json{}
}
