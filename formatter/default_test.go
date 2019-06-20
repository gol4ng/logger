package formatter_test

import (
	"fmt"
	"testing"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/stretchr/testify/assert"
)

func TestDefaultFormatter_Format(t *testing.T) {
	tests := []struct {
		name      string
		formatter *formatter.DefaultFormatter
		entry     logger.Entry
		expected  string
	}{
		{name: "test default formatter struct", formatter: &formatter.DefaultFormatter{}, entry: logger.Entry{}, expected: "<emergency> "},
		{name: "test NewDefaultFormatter()", formatter: formatter.NewDefaultFormatter(), entry: logger.Entry{}, expected: "<emergency> "},
		{name: "test NewDefaultFormatter()", formatter: formatter.NewDefaultFormatter(), entry: logger.Entry{Message: "my message", Level: logger.DebugLevel, Context: logger.Ctx("my_key", "my_value")}, expected: "<debug> my message {\"my_key\":\"my_value\"}"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.formatter.Format(tt.entry))
		})
	}
}

/////////////////////
// Examples
/////////////////////

func ExampleDefaultFormatter() {
	defaultFormatter := formatter.NewDefaultFormatter()

	fmt.Println(defaultFormatter.Format(
		logger.Entry{
			Message: "My log message",
			Level: logger.InfoLevel,
			Context: logger.NewContext().Add("my_key", "my_value"),
		},
	))

	//Output:
	//<info> My log message {"my_key":"my_value"}
}
