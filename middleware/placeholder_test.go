package middleware_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/middleware"
)

func TestPlaceholder(t *testing.T) {
	tests := []struct {
		entry           logger.Entry
		expectedMessage string
	}{
		{expectedMessage: "begin %placeholder% end", entry: newEntry("begin %placeholder% end", nil)},
		{expectedMessage: "begin false end", entry: newEntry("begin %placeholder% end", logger.Ctx("placeholder", false))},
		{expectedMessage: "begin 8 end", entry: newEntry("begin %placeholder% end", logger.Ctx("placeholder", 8))},
		{expectedMessage: "begin my_string end", entry: newEntry("begin %placeholder% end", logger.Ctx("placeholder", "my_string"))},
		{expectedMessage: "begin 10s end", entry: newEntry("begin %placeholder% end", logger.Ctx("placeholder", time.Second*10))},
		{expectedMessage: "begin my_string and my_stringer end", entry: newEntry("begin %placeholder1% and %placeholder2% end", logger.Ctx("placeholder1", "my_string").Add("placeholder2", MyStringer{}))},
		{expectedMessage: "begin my_string and {} end", entry: newEntry("begin %placeholder1% and %placeholder2% end", logger.Ctx("placeholder1", "my_string").Add("placeholder2", struct{}{}))},
		{expectedMessage: "begin my_string and {attrValue} end", entry: newEntry("begin %placeholder1% and %placeholder2% end", logger.Ctx("placeholder1", "my_string").Add("placeholder2", struct{ attr string }{attr: "attrValue"}))},

		{expectedMessage: "begin my_string and again my_string end", entry: newEntry("begin %placeholder% and again %placeholder% end", logger.Ctx("placeholder", "my_string"))},
	}

	for _, tt := range tests {
		t.Run(tt.expectedMessage, func(t *testing.T) {
			mockHandler := func(entry logger.Entry) error {
				assert.Equal(t, tt.expectedMessage, entry.Message)
				return nil
			}
			assert.Nil(t, middleware.Placeholder()(mockHandler)(tt.entry))
		})
	}
}

type MyStringer struct{}

func (s MyStringer) String() string {
	return "my_stringer"
}

func newEntry(message string, context *logger.Context) logger.Entry {
	return logger.Entry{
		Message: message,
		Context: context,
	}
}
