package logger_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/logger"
)

// Create a stringer object
// StringerMock with mockery does'nt work properly
type MyStringer struct{}

func (s MyStringer) String() string {
	return "my_stringer"
}

func TestField(t *testing.T) {
	now := time.Now()
	err := errors.New("my_value")

	tests := []struct {
		name          string
		field         logger.Field
		expectedValue interface{}
		expectedType  logger.FieldType
	}{
		{field: logger.Skip("Skip field", "my_value"), expectedValue: "my_value", expectedType: logger.SkipType},
		{field: logger.Bool("Bool field", true), expectedValue: true, expectedType: logger.BoolType},
		{field: logger.Bool("Bool field", false), expectedValue: false, expectedType: logger.BoolType},
		{field: logger.Int8("Int8 field", 123), expectedValue: int8(123), expectedType: logger.Int8Type},
		{field: logger.Int16("Int16 field", 123), expectedValue: int16(123), expectedType: logger.Int16Type},
		{field: logger.Int32("Int32 field", 123), expectedValue: int32(123), expectedType: logger.Int32Type},
		{field: logger.Int64("Int64 field", 123), expectedValue: int64(123), expectedType: logger.Int64Type},
		{field: logger.Uint8("Uint8 field", 123), expectedValue: uint8(123), expectedType: logger.Uint8Type},
		{field: logger.Uint16("Uint16 field", 123), expectedValue: uint16(123), expectedType: logger.Uint16Type},
		{field: logger.Uint32("Uint32 field", 123), expectedValue: uint32(123), expectedType: logger.Uint32Type},
		{field: logger.Uint64("Uint64 field", 123), expectedValue: uint64(123), expectedType: logger.Uint64Type},
		{field: logger.Uintptr("Uintptr field", 123), expectedValue: uintptr(123), expectedType: logger.UintptrType},
		{field: logger.Float32("Float32 field", 123), expectedValue: float32(123), expectedType: logger.Float32Type},
		{field: logger.Float64("Float64 field", 123), expectedValue: float64(123), expectedType: logger.Float64Type},
		{field: logger.Complex64("Complex64 field", 123), expectedValue: complex64(123), expectedType: logger.Complex64Type},
		{field: logger.Complex128("Complex128 field", 123), expectedValue: complex128(123), expectedType: logger.Complex128Type},
		{field: logger.String("String field", "my_value"), expectedValue: "my_value", expectedType: logger.StringType},
		{field: logger.Binary("Binary field", []byte("my_value")), expectedValue: []byte("my_value"), expectedType: logger.BinaryType},
		{field: logger.ByteString("ByteString field", []byte("my_value")), expectedValue: []byte("my_value"), expectedType: logger.ByteStringType},
		{field: logger.Error("Error field", err), expectedValue: err, expectedType: logger.ErrorType},
		{field: logger.Time("Time field", now), expectedValue: now, expectedType: logger.TimeType},
		{field: logger.Duration("Duration field", 5 * time.Second), expectedValue: 5 * time.Second, expectedType: logger.DurationType},
		{field: logger.Stringer("Stringer field", MyStringer{}), expectedValue: MyStringer{}, expectedType: logger.StringerType},
		{field: logger.Reflect("Reflect field", struct{}{}), expectedValue: struct{}{}, expectedType: logger.ReflectType},
	}

	for _, tt := range tests {
		t.Run(tt.field.Name, func(t *testing.T) {
			assert.Equal(t, tt.expectedType, tt.field.Type)
			assert.Equal(t, tt.expectedValue, tt.field.Value)
		})
	}
}

func TestField_Any(t *testing.T) {
	tests := []struct {
		name               string
		value              interface{}
		expectedType       logger.FieldType
	}{
		{name: "my bool", value: true, expectedType: logger.BoolType},
		{name: "my bool", value: false, expectedType: logger.BoolType},
		{name: "my int", value: 123, expectedType: logger.Int64Type},
		{name: "my int8", value: int8(123), expectedType: logger.Int8Type},
		{name: "my int16", value: int16(123), expectedType: logger.Int16Type},
		{name: "my int32", value: int32(123), expectedType: logger.Int32Type},
		{name: "my int64", value: int64(123), expectedType: logger.Int64Type},
		{name: "my uint8", value: uint8(123), expectedType: logger.Uint8Type},
		{name: "my uint16", value: uint16(123), expectedType: logger.Uint16Type},
		{name: "my uint32", value: uint32(123), expectedType: logger.Uint32Type},
		{name: "my uint64", value: uint64(123), expectedType: logger.Uint64Type},
		{name: "my uintptr", value: uintptr(123), expectedType: logger.UintptrType},
		{name: "my float32", value: float32(123), expectedType: logger.Float32Type},
		{name: "my float64", value: float64(123), expectedType: logger.Float64Type},
		{name: "my complex64", value: complex64(123), expectedType: logger.Complex64Type},
		{name: "my complex128", value: complex128(123), expectedType: logger.Complex128Type},
		{name: "my strings", value: "my strings", expectedType: logger.StringType},
		{name: "my binary", value: []byte{1, 2, 3}, expectedType: logger.BinaryType},
		{name: "my error", value: errors.New("my error message"), expectedType: logger.ErrorType},
		{name: "my time", value: time.Now(), expectedType: logger.TimeType},
		{name: "my duration", value: time.Second, expectedType: logger.DurationType},
		{name: "my stringer", value: MyStringer{}, expectedType: logger.StringerType},
		{name: "my reflect", value: struct{}{}, expectedType: logger.ReflectType},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := logger.Any(tt.name, tt.value)
			assert.Equal(t, tt.expectedType, field.Type)
		})
	}
}

func TestField_String(t *testing.T) {
	now := time.Now()
	err := errors.New("my_error_value")

	tests := []struct {
		name           string
		field          logger.Field
		expectedString string
	}{
		{field: logger.Skip("Skip field", "my_value"), expectedString: "<skipped>"},
		{field: logger.Bool("Bool field", true), expectedString: "true"},
		{field: logger.Bool("Bool field", false), expectedString: "false"},
		{field: logger.Int8("Int8 field", 123), expectedString: "123"},
		{field: logger.Int16("Int16 field", 123), expectedString: "123"},
		{field: logger.Int32("Int32 field", 123), expectedString: "123"},
		{field: logger.Int64("Int64 field", 123), expectedString: "123"},
		{field: logger.Uint8("Uint8 field", 123), expectedString: "123"},
		{field: logger.Uint16("Uint16 field", 123), expectedString: "123"},
		{field: logger.Uint32("Uint32 field", 123), expectedString: "123"},
		{field: logger.Uint64("Uint64 field", 123), expectedString: "123"},
		{field: logger.Uintptr("Uintptr field", 123), expectedString: "123"},
		{field: logger.Float32("Float32 field", 1.23456789), expectedString: "1.234567881"},
		{field: logger.Float64("Float64 field", 123.4567891011), expectedString: "123.4567891"},
		{field: logger.Complex64("Complex64 field", 6 + 7i), expectedString: "(6+7i)"},
		{field: logger.Complex128("Complex128 field", 6 + 7i), expectedString: "(6+7i)"},
		{field: logger.String("String field", "my_value"), expectedString: "my_value"},
		{field: logger.Binary("Binary field", []byte{1, 2, 3}), expectedString: "\x01\x02\x03"},
		{field: logger.ByteString("ByteString field", []byte("my_value")), expectedString: "my_value"},
		{field: logger.Error("Error field", err), expectedString: "my_error_value"},
		{field: logger.Time("Time field", now), expectedString: now.String()},
		{field: logger.Duration("Duration field", 5 * time.Second), expectedString: "5s"},
		{field: logger.Stringer("Stringer field", MyStringer{}), expectedString: "my_stringer"},
		{field: logger.Reflect("Reflect field", struct{}{}), expectedString: "{}"},
	}

	for _, tt := range tests {
		t.Run(tt.field.Name, func(t *testing.T) {
			assert.Equal(t, tt.expectedString, tt.field.String())
		})
	}
}

func TestField_GoString(t *testing.T) {
	now := time.Now()
	error := errors.New("my_error_value")

	tests := []struct {
		name             string
		field            logger.Field
		expectedGoString string
	}{
		{field: logger.Skip("Skip field", "my_value"), expectedGoString: "logger.Field{Name: Skip field, Value: <skipped>, Type: 1}"},
		{field: logger.Bool("Bool field", true), expectedGoString: "logger.Field{Name: Bool field, Value: true, Type: 2}"},
		{field: logger.Bool("Bool field", false), expectedGoString: "logger.Field{Name: Bool field, Value: false, Type: 2}"},
		{field: logger.Int8("Int8 field", 123), expectedGoString: "logger.Field{Name: Int8 field, Value: 123, Type: 3}"},
		{field: logger.Int16("Int16 field", 123), expectedGoString: "logger.Field{Name: Int16 field, Value: 123, Type: 4}"},
		{field: logger.Int32("Int32 field", 123), expectedGoString: "logger.Field{Name: Int32 field, Value: 123, Type: 5}"},
		{field: logger.Int64("Int64 field", 123), expectedGoString: "logger.Field{Name: Int64 field, Value: 123, Type: 6}"},
		{field: logger.Uint8("Uint8 field", 123), expectedGoString: "logger.Field{Name: Uint8 field, Value: 123, Type: 7}"},
		{field: logger.Uint16("Uint16 field", 123), expectedGoString: "logger.Field{Name: Uint16 field, Value: 123, Type: 8}"},
		{field: logger.Uint32("Uint32 field", 123), expectedGoString: "logger.Field{Name: Uint32 field, Value: 123, Type: 9}"},
		{field: logger.Uint64("Uint64 field", 123), expectedGoString: "logger.Field{Name: Uint64 field, Value: 123, Type: 10}"},
		{field: logger.Uintptr("Uintptr field", 123), expectedGoString: "logger.Field{Name: Uintptr field, Value: 123, Type: 11}"},
		{field: logger.Float32("Float32 field", 1.23456789), expectedGoString: "logger.Field{Name: Float32 field, Value: 1.234567881, Type: 12}"},
		{field: logger.Float64("Float64 field", 123.4567891011), expectedGoString: "logger.Field{Name: Float64 field, Value: 123.4567891, Type: 13}"},
		{field: logger.Complex64("Complex64 field", 6 + 7i), expectedGoString: "logger.Field{Name: Complex64 field, Value: (6+7i), Type: 14}"},
		{field: logger.Complex128("Complex128 field", 6 + 7i), expectedGoString: "logger.Field{Name: Complex128 field, Value: (6+7i), Type: 15}"},
		{field: logger.String("String field", "my_value"), expectedGoString: "logger.Field{Name: String field, Value: my_value, Type: 16}"},
		{field: logger.Binary("Binary field", []byte{1, 2, 3}), expectedGoString: "logger.Field{Name: Binary field, Value: \x01\x02\x03, Type: 17}"},
		{field: logger.ByteString("ByteString field", []byte("my_value")), expectedGoString: "logger.Field{Name: ByteString field, Value: my_value, Type: 18}"},
		{field: logger.Error("Error field", error), expectedGoString: "logger.Field{Name: Error field, Value: my_error_value, Type: 19}"},
		{field: logger.Time("Time field", now), expectedGoString: "logger.Field{Name: Time field, Value: "+now.String()+", Type: 20}"},
		{field: logger.Duration("Duration field", 5 * time.Second), expectedGoString: "logger.Field{Name: Duration field, Value: 5s, Type: 21}"},
		{field: logger.Stringer("Stringer field", MyStringer{}), expectedGoString: "logger.Field{Name: Stringer field, Value: my_stringer, Type: 22}"},
		{field: logger.Reflect("Reflect field", struct{}{}), expectedGoString: "logger.Field{Name: Reflect field, Value: {}, Type: 23}"},
	}

	for _, tt := range tests {
		t.Run(tt.field.Name, func(t *testing.T) {
			assert.Equal(t, tt.expectedGoString, tt.field.GoString())
		})
	}
}
