package formatter

import (
	"encoding/json"
	"os"
	"time"

	"github.com/gol4ng/logger"
)

const (
	Version string = "1.1"
)

type Gelf struct {
	hostname string
}

func (j *Gelf) Convert(e logger.Entry) *Message {
	return &Message{
		Version:  Version,
		Host:     j.hostname,
		Short:    e.Message,
		Full:     "TODO",
		TimeUnix: float64(time.Now().Unix()),
		Level:    e.Level,
		//Extra:    *e.Context,
		//Facility: "", this field is deprecated
	}
}

func (j *Gelf) Format(e logger.Entry) string {
	// TODO CHECK json ENCODER
	b, _ := json.Marshal(j.Convert(e))

	return string(b)
}

func NewGelf() (*Gelf, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	return &Gelf{
		hostname: hostname,
	}, nil
}

type Message struct {
	Version  string                 `json:"version"`
	Host     string                 `json:"host"`
	Short    string                 `json:"short_message"`
	Full     string                 `json:"full_message,omitempty"`
	TimeUnix float64                `json:"timestamp"`
	Level    logger.Level           `json:"level,omitempty"`
	Extra    map[string]interface{} `json:"-"`
	//Facility string                 `json:"facility,omitempty"`
}
