package formatter

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gol4ng/logger"
)

const (
	// Version is the current Gelf version
	Version string = "1.1"
)

// Gelf formatter will transform Entry into Gelf format
type Gelf struct {
	hostname string
	version  string
}

// Format will return Entry as gelf
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

	builder.WriteString("}")

	return builder.String()
}

// NewGelf will create a new Gelf formatter
func NewGelf() *Gelf {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	return &Gelf{hostname: hostname, version: Version}
}

// GelfTCPFormatter will decorate Gelf formatter in order to add null byte between each entry
// http://docs.graylog.org/en/3.0/pages/gelf.html#gelf-via-tcp
type GelfTCPFormatter struct {
	*Gelf
}

// Format will return Entry as gelf TCP
func (g *GelfTCPFormatter) Format(entry logger.Entry) string {
	return g.Gelf.Format(entry) + "\x00"
}

// NewGelfTCP will create a new GelfTCPFormatter
func NewGelfTCP() *GelfTCPFormatter {
	return &GelfTCPFormatter{Gelf: NewGelf()}
}
