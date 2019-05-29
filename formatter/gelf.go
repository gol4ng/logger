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

type Gelf struct {
	hostname string
	version  string
}

func (g *Gelf) Format(entry logger.Entry) string {
	builder := &strings.Builder{}

	builder.WriteString(`{"version":"`)
	builder.WriteString(g.version)

	builder.WriteString(`","host":"`)
	builder.WriteString(g.hostname)

	builder.WriteString(`","level":`)
	builder.WriteString(strconv.Itoa(int(entry.Level)))

	builder.WriteString(`,"timestamp":`)
	builder.WriteString(strconv.FormatFloat(float64(time.Now().UnixNano())/1e9, 'f', 3, 64))

	builder.WriteString(`,"short_message":"`)
	builder.WriteString(entry.Message)

	builder.WriteString(`","full_message":"`)
	logger.EntryToString(entry, builder)
	builder.WriteString(`"`)

	if entry.Context != nil {
		for name, field := range *entry.Context {
			builder.WriteString(`,"_`)
			builder.WriteString(strings.Replace(name, " ", "_", -1))
			builder.WriteString(`":`)
			d, _ := json.Marshal(field.Value)
			builder.Write(d)
		}
	}

	builder.WriteString("}\n")

	return builder.String()
}

func NewGelf() *Gelf {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	return &Gelf{hostname: hostname, version: Version}
}
