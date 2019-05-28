package formatter_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/stretchr/testify/assert"
)

func TestJson_Format(t *testing.T) {
	tests := []struct {
		name     string
		entry    logger.Entry
		expected string
	}{
		{
			name:     "test simple message without context",
			entry:    logger.Entry{Message: "test message", Level: logger.DebugLevel, Context: nil},
			expected: `{"Message":"test message","Level":7,"Context":null}`,
		},
		{
			name:     "test simple message with context",
			entry:    logger.Entry{Message: "test message", Level: logger.WarningLevel, Context: logger.NewContext().Add("my_key", "my_value")},
			expected: `{"Message":"test message","Level":4,"Context":{"my_key":"my_value"}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := formatter.NewJsonEncoder()

			assert.Equal(t, tt.expected, f.Format(tt.entry))
		})
	}
}

func TestMarshalContextTo(t *testing.T) {
	tests := []struct {
		name            string
		context         *logger.Context
		expectedStrings []string
	}{
		{
			name:            "test simple message without context",
			context:         nil,
			expectedStrings: []string{"null"},
		},
		{
			name:            "test simple message with context",
			context:         logger.NewContext().Add("my_key", "my_value"),
			expectedStrings: []string{`{"my_key":"my_value"}`},
		},
		{
			name:            "test multiple message with context",
			context:         logger.NewContext().Add("my_key", "my_value").Add("my_key2", "my_value2"),
			expectedStrings: []string{`my_key":"my_value"`, `"my_key2":"my_value2"`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := &strings.Builder{}
			formatter.ContextToJson(tt.context, builder)
			str := builder.String()
			for _, s := range tt.expectedStrings {
				assert.Contains(t, str, s)
			}
		})
	}
}

/////////////////////
// Examples
/////////////////////

func ExampleJsonFormatter() {
	jsonFormatter := formatter.NewJsonEncoder()

	fmt.Println(jsonFormatter.Format(
		logger.Entry{
			Message: "My log message",
			Level: logger.InfoLevel,
			Context: logger.NewContext().Add("my_key", "my_value"),
		},
	))
	//Output:
	//{"Message":"My log message","Level":6,"Context":{"my_key":"my_value"}}
}
