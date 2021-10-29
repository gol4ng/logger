package formatter

import (
	"encoding/json"
	"github.com/valyala/bytebufferpool"
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
	byteBuffer := bytebufferpool.Get()
	defer bytebufferpool.Put(byteBuffer)

	byteBuffer.WriteString(`{"version":"`)
	byteBuffer.WriteString(g.version)

	byteBuffer.WriteString(`","host":"`)
	byteBuffer.WriteString(g.hostname)

	byteBuffer.WriteString(`","level":`)
	byteBuffer.WriteString(strconv.Itoa(int(entry.Level)))

	byteBuffer.WriteString(`,"timestamp":`)
	byteBuffer.WriteString(strconv.FormatFloat(float64(time.Now().UnixNano())/1e9, 'f', 3, 64))

	byteBuffer.WriteString(`,"short_message":"`)
	byteBuffer.WriteString(entry.Message)

	byteBuffer.WriteString(`","full_message":"`)
	logger.EntryToString(entry, byteBuffer)
	byteBuffer.WriteString(`"`)

	if entry.Context != nil {
		for name, field := range *entry.Context {
			byteBuffer.WriteString(`,"_`)
			byteBuffer.WriteString(strings.Replace(name, " ", "_", -1))
			byteBuffer.WriteString(`":`)
			d, _ := json.Marshal(field.Value)
			byteBuffer.Write(d)
		}
	}

	byteBuffer.WriteString("}")

	return byteBuffer.String()
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
