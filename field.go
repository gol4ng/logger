package logger

import (
	"fmt"
	"strconv"
	"time"
)

// FieldType represents the type of a logger context field
type FieldType uint8

// list of available field types
const (
	// UnknownType is the default field type.
	UnknownType FieldType = iota
	// SkipType indicates that the field is a no-op.
	SkipType
	BoolType
	Int8Type
	Int16Type
	Int32Type
	Int64Type
	Uint8Type
	Uint16Type
	Uint32Type
	Uint64Type
	UintptrType
	Float32Type
	Float64Type
	Complex64Type
	Complex128Type
	StringType
	BinaryType
	ByteStringType
	ErrorType
	TimeType
	DurationType
	StringerType
	// ReflectType indicates that the field carries an interface{}, which should be serialized using reflection.
	ReflectType
	// @TODO ArrayMarshalerType
	// @TODO ObjectMarshalerType
)

// Field represents a contextual information
// this data was carry by Context struct
type Field struct {
	Type  FieldType
	Value interface{}
}

// String will return Field as string
func (f *Field) String() string {
	switch f.Type {
	case SkipType:
		return "<skipped>"
	case BoolType:
		if f.Value.(bool) {
			return "true"
		}
		return "false"
	case Int8Type:
		return strconv.Itoa(int(f.Value.(int8)))
	case Int16Type:
		return strconv.Itoa(int(f.Value.(int16)))
	case Int32Type:
		return strconv.Itoa(int(f.Value.(int32)))
	case Int64Type:
		return strconv.Itoa(int(f.Value.(int64)))
	case Uint8Type:
		return strconv.Itoa(int(f.Value.(uint8)))
	case Uint16Type:
		return strconv.Itoa(int(f.Value.(uint16)))
	case Uint32Type:
		return strconv.Itoa(int(f.Value.(uint32)))
	case Uint64Type:
		return strconv.Itoa(int(f.Value.(uint64)))
	case UintptrType:
		return strconv.Itoa(int(f.Value.(uintptr)))
	case Float32Type:
		return strconv.FormatFloat(float64(f.Value.(float32)), 'g', 10, 64)
	case Float64Type:
		return strconv.FormatFloat(f.Value.(float64), 'g', 10, 64)
	case Complex64Type:
		return fmt.Sprintf("%v", f.Value.(complex64))
	case Complex128Type:
		return fmt.Sprintf("%v", f.Value.(complex128))
	case StringType:
		return f.Value.(string)
	case BinaryType:
		return string(f.Value.([]byte)[:])
	case ByteStringType:
		return string(f.Value.([]byte)[:])
	case ErrorType:
		return f.Value.(error).Error()
	case TimeType:
		return f.Value.(time.Time).String()
	case DurationType:
		return f.Value.(time.Duration).String()
	case StringerType:
		return f.Value.(fmt.Stringer).String()
	default:
		return fmt.Sprintf("%v", f.Value)
	}
}

// Skip will create Skip Field
func Skip(value string) Field {
	return Field{Type: SkipType, Value: value}
}

// Bool will create Bool Field
func Bool(value bool) Field {
	return Field{Type: BoolType, Value: value}
}

// Int8 will create Int8 Field
func Int8(value int8) Field {
	return Field{Type: Int8Type, Value: value}
}

// Int16 will create Int16 Field
func Int16(value int16) Field {
	return Field{Type: Int16Type, Value: value}
}

// Int32 will create Int32 Field
func Int32(value int32) Field {
	return Field{Type: Int32Type, Value: value}
}

// Int64 will create Int64 Field
func Int64(value int64) Field {
	return Field{Type: Int64Type, Value: value}
}

// Uint8 will create Uint8 Field
func Uint8(value uint8) Field {
	return Field{Type: Uint8Type, Value: value}
}

// Uint16 will create Uint16 Field
func Uint16(value uint16) Field {
	return Field{Type: Uint16Type, Value: value}
}

// Uint32 will create Uint32 Field
func Uint32(value uint32) Field {
	return Field{Type: Uint32Type, Value: value}
}

// Uint64 will create Uint64 Field
func Uint64(value uint64) Field {
	return Field{Type: Uint64Type, Value: value}
}

// Uintptr will create Uintptr Field
func Uintptr(value uintptr) Field {
	return Field{Type: UintptrType, Value: value}
}

// Float32 will create Float32 Field
func Float32(value float32) Field {
	return Field{Type: Float32Type, Value: value}
}

// Float64 will create Float64 Field
func Float64(value float64) Field {
	return Field{Type: Float64Type, Value: value}
}

// Complex64 will create Complex64 Field
func Complex64(value complex64) Field {
	return Field{Type: Complex64Type, Value: value}
}

// Complex128 will create Complex128 Field
func Complex128(value complex128) Field {
	return Field{Type: Complex128Type, Value: value}
}

// String will create String Field
func String(value string) Field {
	return Field{Type: StringType, Value: value}
}

// Binary will create Binary Field
func Binary(value []byte) Field {
	return Field{Type: BinaryType, Value: value}
}

// ByteString will create ByteString Field
func ByteString(value []byte) Field {
	return Field{Type: ByteStringType, Value: value}
}

// Error will create Error Field
func Error(value error) Field {
	return Field{Type: ErrorType, Value: value}
}

// Stringer will create Stringer Field
func Stringer(value fmt.Stringer) Field {
	return Field{Type: StringerType, Value: value}
}

// Time will create Time Field
func Time(value time.Time) Field {
	return Field{Type: TimeType, Value: value}
}

// Duration will create Duration Field
func Duration(value time.Duration) Field {
	return Field{Type: DurationType, Value: value}
}

// Reflect will create Reflect Field
func Reflect(value interface{}) Field {
	return Field{Type: ReflectType, Value: value}
}

// Any will guess and create Field with givan value
func Any(value interface{}) Field {
	switch val := value.(type) {
	case bool:
		return Bool(val)
	case int:
		return Int64(int64(val))
	case int8:
		return Int8(val)
	case int16:
		return Int16(val)
	case int32:
		return Int32(val)
	case int64:
		return Int64(val)
	case uint8:
		return Uint8(val)
	case uint16:
		return Uint16(val)
	case uint32:
		return Uint32(val)
	case uint64:
		return Uint64(val)
	case uintptr:
		return Uintptr(val)
	case float32:
		return Float32(val)
	case float64:
		return Float64(val)
	case complex64:
		return Complex64(val)
	case complex128:
		return Complex128(val)
	case string:
		return String(val)
	case []byte:
		return Binary(val)
	case error:
		return Error(val)
	case time.Time:
		return Time(val)
	case time.Duration:
		return Duration(val)
	case fmt.Stringer:
		return Stringer(val)
	default:
		return Reflect(val)
	}
}
