package logger_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/gol4ng/logger"
)

func TestContext_Merge(t *testing.T) {
	context1 := logger.Context(map[string]logger.Field{
		"my_key":                  logger.Field{Value: "my_value"},
		"key_gonna_be_overwrited": logger.Field{Value: "my_initial_value"},
	})
	context2 := logger.Context(map[string]logger.Field{
		"key_gonna_be_overwrited": logger.Field{Value: 2},
		"new_key":                 logger.Field{Value: "my_new_value"},
	})

	(&context1).Merge(context2)

	assert.Equal(t, "my_value", context1["my_key"].Value)
	assert.Equal(t, 2, context1["key_gonna_be_overwrited"].Value)
	assert.Equal(t, "my_new_value", context1["new_key"].Value)
}

func TestContext_Set(t *testing.T) {
	context1 := logger.Context(map[string]logger.Field{
		"my_key": logger.Field{Value: "my_value"},
	})

	context1.Set("my_key", logger.Field{Value: 2})

	assert.Equal(t, 2, context1["my_key"].Value)
}

func TestContext_Has(t *testing.T) {
	context1 := logger.Context(map[string]logger.Field{
		"my_key": logger.Field{Value: "my_value"},
	})

	assert.True(t, context1.Has("my_key"))
	assert.False(t, context1.Has("unknow_key"))
}

func TestContext_Get(t *testing.T) {
	context := logger.Context(map[string]logger.Field{
		"my_key": logger.Field{Value: "my_value"},
	})

	defaultValue := &logger.Field{}
	assert.Equal(t, "my_value", context.Get("my_key", nil).Value)
	assert.Equal(t, *defaultValue, context.Get("unknow_key", defaultValue))
}

func TestContext_Add(t *testing.T) {
	context := logger.Context(map[string]logger.Field{})
	context.
		Add("my_key", "my_value").
		Add("my_key", int8(2))

	assert.Equal(t, int8(2), context["my_key"].Value)
}

func TestContext_Skip(t *testing.T) {
	context := logger.Context(map[string]logger.Field{})
	context.Skip("my_key", "my_value")
	field := context.Get("my_key", nil)
	assert.Equal(t, "my_value", field.Value)
	assert.Equal(t, logger.SkipType, field.Type)
}

func TestContext_Binary(t *testing.T) {
	context := logger.Context(map[string]logger.Field{})
	context.Binary("my_key", []byte{1, 2, 3})
	field := context.Get("my_key", nil)
	assert.Equal(t, []byte{1, 2, 3}, field.Value)
	assert.Equal(t, logger.BinaryType, field.Type)
}

func TestContext_ByteString(t *testing.T) {
	context := logger.Context(map[string]logger.Field{})
	context.ByteString("my_key", []byte("my_value"))
	field := context.Get("my_key", nil)
	assert.Equal(t, []byte("my_value"), field.Value)
	assert.Equal(t, logger.ByteStringType, field.Type)
}

func TestContext_Ctx(t *testing.T) {
	context := logger.Ctx("my_key", "my_value")

	assert.Equal(t, "my_value", (*context)["my_key"].Value)
}

func TestContext_String(t *testing.T) {
	tests := []struct {
		name    string
		context *logger.Context
		strings []string
	}{
		{name: "Empty context", context: logger.NewContext(), strings: []string{"<nil>"}},
		{name: "Simple context", context: logger.Ctx("my_key", "my_value"), strings: []string{"my_key:my_value"}},
		{name: "Simple context", context: logger.Ctx("my_key", "my_value").Add("my_int_val", 3), strings: []string{"my_key:my_value", "my_int_val:3"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := tt.context.String()
			for _, s := range tt.strings {
				assert.Contains(t, str, s)
			}
		})
	}
}

func TestContext_GoString(t *testing.T) {
	tests := []struct {
		name      string
		context   *logger.Context
		goStrings []string
	}{
		{name: "Empty context", context: logger.NewContext(), goStrings: []string{"<nil>"}},
		{name: "Simple context", context: logger.Ctx("my_key", "my_value"), goStrings: []string{"my_key:my_value"}},
		{name: "Simple context", context: logger.Ctx("my_key", "my_value").Add("my_int_val", 3), goStrings: []string{"my_key:my_value", "my_int_val:3"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := tt.context.GoString()
			assert.Contains(t, str, "logger.context<")
			for _, s := range tt.goStrings {
				assert.Contains(t, str, s)
			}
		})
	}
}
