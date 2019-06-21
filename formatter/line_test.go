package formatter_test

import (
	"fmt"
	"testing"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/stretchr/testify/assert"
)

func TestLine_Format(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		expected string
	}{
		{
			name:     "test simple format 1",
			format:   "%s %s %s",
			expected: "test message warning <my_key:my_value>",
		},
		{
			name:     "test simple format 2",
			format:   "%s %d %s",
			expected: "test message 4 <my_key:my_value>",
		},
		{
			name:     "test simple format",
			format:   "%[2]s %[1]s %[3]s",
			expected: "warning test message <my_key:my_value>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := formatter.NewLine(tt.format)

			assert.Equal(t, tt.expected, f.Format(logger.Entry{Message: "test message", Level: logger.WarningLevel, Context: logger.Ctx("my_key", "my_value")}))
		})
	}
}

/////////////////////
// Examples
/////////////////////

func ExampleLineFormatter() {
	lineFormatter := formatter.NewLine("%s %s %s")

	fmt.Println(lineFormatter.Format(
		logger.Entry{
			Message: "My log message",
			Level: logger.InfoLevel,
			Context: logger.NewContext().Add("my_key", "my_value"),
		},
	))

	//Output:
	//My log message info <my_key:my_value>
}
