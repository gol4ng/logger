package logger_test

import (
	"errors"
	"runtime"
	"testing"
	"time"

	"github.com/gol4ng/logger"
	"github.com/stretchr/testify/assert"
)

// Create a stringer object
// StringerMock with mockery does'nt work properly
type MyStringer struct{}

func (s MyStringer) String() string {
	return "my_stringer"
}

func TestField(t *testing.T) {
	now := time.Now()
	error := errors.New("my_value")

	tests := []struct {
		name          string
		field         logger.Field
		expectedValue interface{}
		expectedType  logger.FieldType
	}{
		{name: "Skip field", field: logger.Skip("my_value"), expectedValue: "my_value", expectedType: logger.SkipType},
		{name: "Bool field", field: logger.Bool(true), expectedValue: true, expectedType: logger.BoolType},
		{name: "Bool field", field: logger.Bool(false), expectedValue: false, expectedType: logger.BoolType},
		{name: "Int8 field", field: logger.Int8(123), expectedValue: int8(123), expectedType: logger.Int8Type},
		{name: "Int16 field", field: logger.Int16(123), expectedValue: int16(123), expectedType: logger.Int16Type},
		{name: "Int32 field", field: logger.Int32(123), expectedValue: int32(123), expectedType: logger.Int32Type},
		{name: "Int64 field", field: logger.Int64(123), expectedValue: int64(123), expectedType: logger.Int64Type},
		{name: "Uint8 field", field: logger.Uint8(123), expectedValue: uint8(123), expectedType: logger.Uint8Type},
		{name: "Uint16 field", field: logger.Uint16(123), expectedValue: uint16(123), expectedType: logger.Uint16Type},
		{name: "Uint32 field", field: logger.Uint32(123), expectedValue: uint32(123), expectedType: logger.Uint32Type},
		{name: "Uint64 field", field: logger.Uint64(123), expectedValue: uint64(123), expectedType: logger.Uint64Type},
		{name: "Uintptr field", field: logger.Uintptr(123), expectedValue: uintptr(123), expectedType: logger.UintptrType},
		{name: "Float32 field", field: logger.Float32(123), expectedValue: float32(123), expectedType: logger.Float32Type},
		{name: "Float64 field", field: logger.Float64(123), expectedValue: float64(123), expectedType: logger.Float64Type},
		{name: "Complex64 field", field: logger.Complex64(123), expectedValue: complex64(123), expectedType: logger.Complex64Type},
		{name: "Complex128 field", field: logger.Complex128(123), expectedValue: complex128(123), expectedType: logger.Complex128Type},
		{name: "String field", field: logger.String("my_value"), expectedValue: "my_value", expectedType: logger.StringType},
		{name: "Binary field", field: logger.Binary([]byte("my_value")), expectedValue: []byte("my_value"), expectedType: logger.BinaryType},
		{name: "ByteString field", field: logger.ByteString([]byte("my_value")), expectedValue: []byte("my_value"), expectedType: logger.ByteStringType},
		{name: "Error field", field: logger.Error(error), expectedValue: error, expectedType: logger.ErrorType},
		{name: "Time field", field: logger.Time(now), expectedValue: now, expectedType: logger.TimeType},
		{name: "Duration field", field: logger.Duration(5 * time.Second), expectedValue: 5 * time.Second, expectedType: logger.DurationType},
		{name: "Stringer field", field: logger.Stringer(MyStringer{}), expectedValue: MyStringer{}, expectedType: logger.StringerType},
		{name: "Reflect field", field: logger.Reflect(struct{}{}), expectedValue: struct{}{}, expectedType: logger.ReflectType},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
		expectedMallocs    uint64
		expectedTotalAlloc uint64
	}{
		{name: "my bool", value: true, expectedType: logger.BoolType, expectedMallocs: 0, expectedTotalAlloc: 0},
		{name: "my bool", value: false, expectedType: logger.BoolType, expectedMallocs: 0, expectedTotalAlloc: 0},
		{name: "my int", value: 123, expectedType: logger.Int64Type, expectedMallocs: 1, expectedTotalAlloc: 16},
		{name: "my int8", value: int8(123), expectedType: logger.Int8Type, expectedMallocs: 0, expectedTotalAlloc: 0},
		{name: "my int16", value: int16(123), expectedType: logger.Int16Type, expectedMallocs: 1, expectedTotalAlloc: 16},
		{name: "my int32", value: int32(123), expectedType: logger.Int32Type, expectedMallocs: 1, expectedTotalAlloc: 16},
		{name: "my int64", value: int64(123), expectedType: logger.Int64Type, expectedMallocs: 1, expectedTotalAlloc: 16},
		{name: "my uint8", value: uint8(123), expectedType: logger.Uint8Type, expectedMallocs: 0, expectedTotalAlloc: 0},
		{name: "my uint16", value: uint16(123), expectedType: logger.Uint16Type, expectedMallocs: 1, expectedTotalAlloc: 16},
		{name: "my uint32", value: uint32(123), expectedType: logger.Uint32Type, expectedMallocs: 1, expectedTotalAlloc: 16},
		{name: "my uint64", value: uint64(123), expectedType: logger.Uint64Type, expectedMallocs: 1, expectedTotalAlloc: 16},
		{name: "my uintptr", value: uintptr(123), expectedType: logger.UintptrType, expectedMallocs: 1, expectedTotalAlloc: 16},
		{name: "my float32", value: float32(123), expectedType: logger.Float32Type, expectedMallocs: 1, expectedTotalAlloc: 16},
		{name: "my float64", value: float64(123), expectedType: logger.Float64Type, expectedMallocs: 1, expectedTotalAlloc: 16},
		{name: "my complex64", value: complex64(123), expectedType: logger.Complex64Type, expectedMallocs: 1, expectedTotalAlloc: 16},
		{name: "my complex128", value: complex128(123), expectedType: logger.Complex128Type, expectedMallocs: 1, expectedTotalAlloc: 16},
		{name: "my strings", value: "my strings", expectedType: logger.StringType, expectedMallocs: 1, expectedTotalAlloc: 16},
		{name: "my binary", value: []byte{1, 2, 3}, expectedType: logger.BinaryType, expectedMallocs: 1, expectedTotalAlloc: 32},
		{name: "my error", value: errors.New("my error message"), expectedType: logger.ErrorType, expectedMallocs: 0, expectedTotalAlloc: 0},
		{name: "my time", value: time.Now(), expectedType: logger.TimeType, expectedMallocs: 1, expectedTotalAlloc: 32},
		{name: "my duration", value: time.Second, expectedType: logger.DurationType, expectedMallocs: 1, expectedTotalAlloc: 16},
		{name: "my stringer", value: MyStringer{}, expectedType: logger.StringerType, expectedMallocs: 0, expectedTotalAlloc: 0},
		{name: "my reflect", value: struct{}{}, expectedType: logger.ReflectType, expectedMallocs: 0, expectedTotalAlloc: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var memStatsStart runtime.MemStats
			runtime.ReadMemStats(&memStatsStart)

			field := logger.Any(tt.value)

			var memStatsEnd runtime.MemStats
			runtime.ReadMemStats(&memStatsEnd)
			assert.Equal(t, tt.expectedMallocs, memStatsEnd.Mallocs-memStatsStart.Mallocs)
			// This assertion was to variable
			//if tt.expectedTotalAlloc != memStatsEnd.TotalAlloc-memStatsStart.TotalAlloc {
			//	t.Logf(
			//		"Test \"%s\" totalAlloc:%d expected %d",
			//		tt.name,
			//		memStatsEnd.TotalAlloc-memStatsStart.TotalAlloc,
			//		tt.expectedTotalAlloc,
			//	)
			//}
			assert.Equal(t, tt.expectedType, field.Type)
		})
	}
}

func TestField_String(t *testing.T) {
	now := time.Now()
	error := errors.New("my_error_value")

	tests := []struct {
		name           string
		field          logger.Field
		expectedString string
	}{
		{name: "Skip field", field: logger.Skip("my_value"), expectedString: "<skipped>"},
		{name: "Bool field", field: logger.Bool(true), expectedString: "true"},
		{name: "Bool field", field: logger.Bool(false), expectedString: "false"},
		{name: "Int8 field", field: logger.Int8(123), expectedString: "123"},
		{name: "Int16 field", field: logger.Int16(123), expectedString: "123"},
		{name: "Int32 field", field: logger.Int32(123), expectedString: "123"},
		{name: "Int64 field", field: logger.Int64(123), expectedString: "123"},
		{name: "Uint8 field", field: logger.Uint8(123), expectedString: "123"},
		{name: "Uint16 field", field: logger.Uint16(123), expectedString: "123"},
		{name: "Uint32 field", field: logger.Uint32(123), expectedString: "123"},
		{name: "Uint64 field", field: logger.Uint64(123), expectedString: "123"},
		{name: "Uintptr field", field: logger.Uintptr(123), expectedString: "123"},
		{name: "Float32 field", field: logger.Float32(1.23456789), expectedString: "1.234567881"},
		{name: "Float64 field", field: logger.Float64(123.4567891011), expectedString: "123.4567891"},
		{name: "Complex64 field", field: logger.Complex64(6 + 7i), expectedString: "(6+7i)"},
		{name: "Complex128 field", field: logger.Complex128(6 + 7i), expectedString: "(6+7i)"},
		{name: "String field", field: logger.String("my_value"), expectedString: "my_value"},
		{name: "Binary field", field: logger.Binary([]byte{1, 2, 3}), expectedString: "\x01\x02\x03"},
		{name: "ByteString field", field: logger.ByteString([]byte("my_value")), expectedString: "my_value"},
		{name: "Error field", field: logger.Error(error), expectedString: "my_error_value"},
		{name: "Time field", field: logger.Time(now), expectedString: now.String()},
		{name: "Duration field", field: logger.Duration(5 * time.Second), expectedString: "5s"},
		{name: "Stringer field", field: logger.Stringer(MyStringer{}), expectedString: "my_stringer"},
		{name: "Reflect field", field: logger.Reflect(struct{}{}), expectedString: "{}"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectedString, tt.field.String())
		})
	}
}
