package logger_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/logger"
)

func TestFields_Add(t *testing.T) {
	fields := logger.Fields{}
	fields.
		Add("my_key", "my_value").
		Add("my_key", int8(2))

	assert.Equal(t, "my_value", fields[0].Value)
	assert.Equal(t, int8(2), fields[1].Value)
}

func TestFields_Skip(t *testing.T) {
	fields := logger.Fields{}
	fields.Skip("my_key", "my_value")
	assert.Equal(t, "my_value", fields[0].Value)
	assert.Equal(t, logger.SkipType, fields[0].Type)
}

func TestFields_Binary(t *testing.T) {
	fields := logger.Fields{}
	fields.Binary("my_key", []byte{1, 2, 3})
	assert.Equal(t, []byte{1, 2, 3}, fields[0].Value)
	assert.Equal(t, logger.BinaryType, fields[0].Type)
}

func TestFields_ByteString(t *testing.T) {
	fields := logger.Fields{}
	fields.ByteString("my_key", []byte("my_value"))
	assert.Equal(t, []byte("my_value"), fields[0].Value)
	assert.Equal(t, logger.ByteStringType, fields[0].Type)
}

func TestFields_String(t *testing.T) {
	tests := []struct {
		name   string
		fields *logger.Fields
		string string
	}{
		{name: "Empty fields", fields: &logger.Fields{}, string: " "},
		{name: "Simple fields", fields: logger.NewFields("my_key", "my_value"), string: "<my_key:my_value>"},
		{name: "Simple fields", fields: logger.NewFields("my_key", "my_value").Add("my_int_val", 3), string: "<my_key:my_value> <my_int_val:3>"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.string, tt.fields.String())
		})
	}
}

func TestFields_GoString(t *testing.T) {
	tests := []struct {
		name     string
		fields   *logger.Fields
		goString string
	}{
		{name: "Empty fields", fields: &logger.Fields{}, goString: "logger.Fields[ ]"},
		{name: "Simple fields", fields: logger.NewFields("my_key", "my_value"), goString: "logger.Fields[<my_key:my_value>]"},
		{name: "Simple fields", fields: logger.NewFields("my_key", "my_value").Add("my_int_val", 3), goString: "logger.Fields[<my_key:my_value> <my_int_val:3>]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.goString, tt.fields.GoString())
		})
	}
}
