package formatter_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
)

func TestDefaultFormatter_Format(t *testing.T) {
	tests := []struct {
		name      string
		formatter *formatter.DefaultFormatter
		entry     logger.Entry
		expected  string
	}{
		{name: "test default formatter struct", formatter: &formatter.DefaultFormatter{}, entry: logger.Entry{}, expected: "<emergency>"},
		{name: "test NewDefaultFormatter()", formatter: formatter.NewDefaultFormatter(), entry: logger.Entry{}, expected: "<emergency>"},
		{
			name:      "test NewDefaultFormatter(formatter.WithContext)",
			formatter: formatter.NewDefaultFormatter(formatter.WithContext),
			entry:     logger.Entry{Message: "my message", Level: logger.DebugLevel, Context: logger.Ctx("my_key", "my_value")},
			expected:  "<debug> my message {\"my_key\":\"my_value\"}",
		},
		{
			name:      "test formatter.NewDefaultFormatter(formatter.WithColor)",
			formatter: formatter.NewDefaultFormatter(formatter.WithColor),
			entry:     logger.Entry{Message: "my message", Level: logger.DebugLevel, Context: logger.Ctx("my_key", "my_value")},
			expected:  "\x1b[1;36m<debug>\x1b[m my message",
		},
		{
			name:      "test formatter.NewDefaultFormatter(formatter.WithColor, formatter.WithContext)",
			formatter: formatter.NewDefaultFormatter(formatter.WithColor, formatter.WithContext),
			entry:     logger.Entry{Message: "my message", Level: logger.DebugLevel, Context: logger.Ctx("my_key", "my_value")},
			expected:  "\x1b[1;36m<debug>\x1b[m my message {\"my_key\":\"my_value\"}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.formatter.Format(tt.entry))
		})
	}
}

// =====================================================================================================================
// ================================================= EXAMPLES ==========================================================
// =====================================================================================================================

func ExampleDefaultFormatter() {
	defaultFormatter := formatter.NewDefaultFormatter(formatter.WithContext)

	fmt.Println(defaultFormatter.Format(
		logger.Entry{
			Message: "My log message",
			Level:   logger.InfoLevel,
			Context: logger.NewContext().Add("my_key", "my_value"),
		},
	))

	//Output:
	//<info> My log message {"my_key":"my_value"}
}
