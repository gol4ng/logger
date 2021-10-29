package logger_test

import (
	"github.com/valyala/bytebufferpool"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/logger"
)

func TestEntry_String(t *testing.T) {
	tests := []struct {
		name    string
		entry   logger.Entry
		strings []string
	}{
		{name: "Empty context", entry: logger.Entry{Message: "log message", Level: logger.DebugLevel, Context: logger.NewContext()}, strings: []string{"<debug> log message [ <nil> ]"}},
		{name: "Simple entry", entry: logger.Entry{Message: "log message", Level: logger.DebugLevel, Context: logger.Ctx("my_key", "my_value")}, strings: []string{"<debug> log message [ <my_key:my_value> ]"}},
		{name: "Simple entry", entry: logger.Entry{Message: "log message", Level: logger.DebugLevel, Context: logger.Ctx("my_key", "my_value").Add("my_int_val", 3)}, strings: []string{"<debug> log message [", "<my_key:my_value>", "<my_int_val:3>"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := tt.entry.String()
			for _, s := range tt.strings {
				assert.Contains(t, str, s)
			}
		})
	}
}

func TestEntryToString(t *testing.T) {
	tests := []struct {
		name    string
		entry   logger.Entry
		strings []string
	}{
		{name: "Empty context", entry: logger.Entry{Message: "log message", Level: logger.DebugLevel, Context: logger.NewContext()}, strings: []string{"<debug> log message [ <nil> ]"}},
		{name: "Simple entry", entry: logger.Entry{Message: "log message", Level: logger.DebugLevel, Context: logger.Ctx("my_key", "my_value")}, strings: []string{"<debug> log message [ <my_key:my_value> ]"}},
		{name: "Simple entry", entry: logger.Entry{Message: "log message", Level: logger.DebugLevel, Context: logger.Ctx("my_key", "my_value").Add("my_int_val", 3)}, strings: []string{"<debug> log message [", "<my_key:my_value>", "<my_int_val:3>"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			byteBuffer := bytebufferpool.Get()
			defer bytebufferpool.Put(byteBuffer)

			logger.EntryToString(tt.entry, byteBuffer)
			for _, s := range tt.strings {
				assert.Contains(t, byteBuffer.String(), s)
			}
		})
	}
}
