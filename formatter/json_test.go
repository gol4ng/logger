package formatter_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/stretchr/testify/assert"
)

func TestJson_Format(t *testing.T) {
	tests := []struct {
		name     string
		entry    logger.Entry
		options  []formatter.JSONOption
		expected map[string]interface{}
	}{
		{
			name:    "test simple message without context",
			entry:   logger.Entry{Message: "test message", Level: logger.DebugLevel, Context: nil},
			options: []formatter.JSONOption{},
			expected: map[string]interface{}{
				"message": "test message",
				"level":   "7",
			},
		},
		{
			name:  "test simple message without context and levelAsString option",
			entry: logger.Entry{Message: "test message", Level: logger.DebugLevel, Context: nil},
			options: []formatter.JSONOption{
				formatter.WithJSONLevelAsString(),
			},
			expected: map[string]interface{}{
				"message": "test message",
				"level":   "debug",
			},
		},
		{
			name:    "test simple message with context",
			entry:   logger.Entry{Message: "test message", Level: logger.WarningLevel, Context: logger.NewContext().Add("my_key", "my_value")},
			options: []formatter.JSONOption{},
			expected: map[string]interface{}{
				"message": "test message",
				"level":   "4",
				"Context": map[string]interface{}{
					"my_key": "my_value",
				},
			},
		},
		{
			name:  "test simple message with context and flatten and levelAsString options",
			entry: logger.Entry{Message: "test message", Level: logger.WarningLevel, Context: logger.NewContext().Add("my_key", "my_value")},
			options: []formatter.JSONOption{
				formatter.WithJSONFlattenContext(),
				formatter.WithJSONLevelAsString(),
			},
			expected: map[string]interface{}{
				"message": "test message",
				"level":   "warning",
				"my_key":  "my_value",
			},
		},
		{
			name: "test simple message with context and multiple fields",
			entry: logger.Entry{
				Message: "test message",
				Level:   logger.ErrorLevel,
				Context: logger.NewContext().
					Add("my_key_1", "my_value_1").
					Add("my_key_2", "my_value_2"),
			},
			options: []formatter.JSONOption{},
			expected: map[string]interface{}{
				"message": "test message",
				"level":   "3",
				"Context": map[string]interface{}{
					"my_key_1": "my_value_1",
					"my_key_2": "my_value_2",
				},
			},
		},
		{
			name: "test simple message with context and multiple fields and flatten and levelAsString options",
			entry: logger.Entry{
				Message: "test message",
				Level:   logger.ErrorLevel,
				Context: logger.NewContext().
					Add("my_key_1", "my_value_1").
					Add("my_key_2", "my_value_2"),
			},
			options: []formatter.JSONOption{
				formatter.WithJSONFlattenContext(),
				formatter.WithJSONLevelAsString(),
			},
			expected: map[string]interface{}{
				"message":  "test message",
				"level":    "error",
				"my_key_1": "my_value_1",
				"my_key_2": "my_value_2",
			},
		},
		{
			name: "test simple message with context and multiple fields and duplicated field",
			entry: logger.Entry{
				Message: "test message",
				Level:   logger.ErrorLevel,
				Context: logger.NewContext().
					Add("my_key_1", "my_value_1").
					Add("my_key_1", "my_value_1").
					Add("my_key_2", "my_value_2"),
			},
			options: []formatter.JSONOption{},
			expected: map[string]interface{}{
				"message": "test message",
				"level":   "3",
				"Context": map[string]interface{}{
					"my_key_1": "my_value_1",
					"my_key_2": "my_value_2",
				},
			},
		},
		{
			name: "test simple message with context and with usage of reserved keys (should appear)",
			entry: logger.Entry{
				Message: "test message",
				Level:   logger.ErrorLevel,
				Context: logger.NewContext().
					Add("message", "my test message").
					Add("level", "a level entry"),
			},
			options: []formatter.JSONOption{},
			expected: map[string]interface{}{
				"message": "test message",
				"level":   "3",
				"Context": map[string]interface{}{
					"message": "my test message",
					"level":   "a level entry",
				},
			},
		},
		{
			name: "test simple message with context, flatten and with reserved keys (should not appear)",
			entry: logger.Entry{
				Message: "test message",
				Level:   logger.ErrorLevel,
				Context: logger.NewContext().
					Add("my_key_1", "my_value_1").
					Add("message", "my test message that should not appear").
					Add("level", "a level entry that should not appear").
					Add("my_key_2", "my_value_2"),
			},
			options: []formatter.JSONOption{
				formatter.WithJSONFlattenContext(),
			},
			expected: map[string]interface{}{
				"message":  "test message",
				"level":    "3",
				"my_key_1": "my_value_1",
				"my_key_2": "my_value_2",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := formatter.NewJSONEncoder(tt.options...)
			result := f.Format(tt.entry)

			got := map[string]interface{}{}
			if err := json.Unmarshal([]byte(result), &got); err != nil {
				assert.Fail(t, "unable to unmarshal result", result)
			}

			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestMarshalContextTo(t *testing.T) {
	tests := []struct {
		name            string
		context         *logger.Context
		options         *formatter.JSONOptions
		expectedStrings []string
	}{
		{
			name:            "test simple message without context",
			context:         nil,
			options:         &formatter.JSONOptions{},
			expectedStrings: []string{""},
		},
		{
			name:            "test simple message with context",
			context:         logger.NewContext().Add("my_key", "my_value"),
			options:         &formatter.JSONOptions{},
			expectedStrings: []string{`"my_key":"my_value"`},
		},
		{
			name:            "test multiple message with context",
			context:         logger.NewContext().Add("my_key", "my_value").Add("my_key2", "my_value2"),
			options:         &formatter.JSONOptions{},
			expectedStrings: []string{`"my_key":"my_value"`, `"my_key2":"my_value2"`},
		},
		{
			name:            "test time message with context",
			context:         logger.NewContext().Add("my_key", time.Date(2020, 1, 2, 3, 4, 5, 6, time.UTC)),
			options:         &formatter.JSONOptions{},
			expectedStrings: []string{`my_key":"2020-01-02T03:04:05.000000006Z"`},
		},
		{
			name:            "test time message with context",
			context:         logger.NewContext().Add("my_key", errors.New("my error message")),
			options:         &formatter.JSONOptions{},
			expectedStrings: []string{`my_key":"my error message"`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := &strings.Builder{}
			formatter.ContextToJSON(tt.context, builder, tt.options)
			str := builder.String()
			for _, s := range tt.expectedStrings {
				assert.Contains(t, str, s)
			}
		})
	}
}

// =====================================================================================================================
// ================================================= EXAMPLES ==========================================================
// =====================================================================================================================

func ExampleJsonFormatter() {
	jsonFormatter := formatter.NewJSONEncoder()

	fmt.Println(jsonFormatter.Format(
		logger.Entry{
			Message: "My log message",
			Level:   logger.InfoLevel,
			Context: logger.NewContext().Add("my_key", "my_value"),
		},
	))
	//Output:
	//{"message":"My log message","level":"6","Context":{"my_key":"my_value"}}
}
