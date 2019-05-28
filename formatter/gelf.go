package formatter

import (
	"encoding/json"
	"github.com/gol4ng/logger"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	Version string = "1.1"
)

type GelfEncoder struct {
	Json
	hostname string
	version  string
}

func (g *GelfEncoder) encode(entry logger.Entry) ([]byte, error) {
	builder := &strings.Builder{}

	builder.WriteString("{\"version\":\"")
	builder.WriteString(g.version)

	builder.WriteString("\",\"host\":\"")
	builder.WriteString(g.hostname)

	builder.WriteString("\",\"level\":")
	builder.WriteString(strconv.Itoa(int(entry.Level)))

	builder.WriteString(",\"timestamp\":")
	builder.WriteString(strconv.FormatFloat(float64(time.Now().UnixNano())/1e9, 'f', 3, 64))

	builder.WriteString(",\"short_message\":\"")
	builder.WriteString(entry.Message)

	builder.WriteString("\",\"full_message\":\"")
	logger.EntryToString(entry, builder)
	builder.WriteString("\"")

	for name, field := range *entry.Context {
		builder.WriteString(",\"_")
		builder.WriteString(strings.Replace(name, " ", "_", -1))
		builder.WriteString("\":")
		d, _ := json.Marshal(field.Value)
		builder.Write(d)
	}

	builder.WriteString("}\n")

	return []byte(builder.String()), nil
}

func NewGelfEncoder() *Json {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	e := GelfEncoder{hostname: hostname, version: Version}
	return NewJson(e.encode)
}
