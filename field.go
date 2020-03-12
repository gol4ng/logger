package logger

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
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
	// @TODO ArrayMarshallerType
	// @TODO ObjectMarshallerType
)

// Field represents a contextual information
// this data was carry by Context struct
type Field struct {
	Name  string
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
		return strconv.FormatInt(int64(f.Value.(int8)), 10)
	case Int16Type:
		return strconv.FormatInt(int64(f.Value.(int16)), 10)
	case Int32Type:
		return strconv.FormatInt(int64(f.Value.(int32)), 10)
	case Int64Type:
		return strconv.FormatInt(f.Value.(int64), 10)
	case Uint8Type:
		return strconv.FormatUint(uint64(f.Value.(uint8)), 10)
	case Uint16Type:
		return strconv.FormatUint(uint64(f.Value.(uint16)), 10)
	case Uint32Type:
		return strconv.FormatUint(uint64(f.Value.(uint32)), 10)
	case Uint64Type:
		return strconv.FormatUint(f.Value.(uint64), 10)
	case UintptrType:
		return strconv.FormatUint(uint64(f.Value.(uintptr)), 10)
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

// MarshalJSON was called by json.Marshal(field)
// json Marshaller interface
func (f *Field) MarshalJSON() ([]byte, error) {
	if marshallable, ok := f.Value.(json.Marshaler); ok {
		return marshallable.MarshalJSON()
	}
	switch f.Type {
	case BoolType, Int8Type, Int16Type, Int32Type, Int64Type, Uint8Type, Uint16Type, Uint32Type, Uint64Type, UintptrType, Float32Type, Float64Type:
		return []byte(f.String()), nil
	case SkipType, Complex64Type, Complex128Type, StringType, BinaryType, ByteStringType, ErrorType, TimeType, DurationType, StringerType:
		return strconv.AppendQuote([]byte{}, f.String()), nil
	default:
		return json.Marshal(f.Value)
	}
}

// GoString was called by fmt.Printf("%#v", Fields)
// fmt GoStringer interface
func (f *Field) GoString() string {
	builder := &strings.Builder{}
	builder.WriteString("logger.Field{Name: ")
	builder.WriteString(f.Name)
	builder.WriteString(", Value: ")
	builder.WriteString(f.String())
	builder.WriteString(", Type: ")
	builder.WriteString(strconv.FormatUint(uint64(f.Type), 10))
	builder.WriteString("}")
	return builder.String()
}

// Skip will create Skip Field
func Skip(name string, value string) Field {
	return Field{Name: name, Type: SkipType, Value: value}
}

// Bool will create Bool Field
func Bool(name string, value bool) Field {
	return Field{Name: name, Type: BoolType, Value: value}
}

// Int8 will create Int8 Field
func Int8(name string, value int8) Field {
	return Field{Name: name, Type: Int8Type, Value: value}
}

// Int16 will create Int16 Field
func Int16(name string, value int16) Field {
	return Field{Name: name, Type: Int16Type, Value: value}
}

// Int32 will create Int32 Field
func Int32(name string, value int32) Field {
	return Field{Name: name, Type: Int32Type, Value: value}
}

// Int64 will create Int64 Field
func Int64(name string, value int64) Field {
	return Field{Name: name, Type: Int64Type, Value: value}
}

// Uint8 will create Uint8 Field
func Uint8(name string, value uint8) Field {
	return Field{Name: name, Type: Uint8Type, Value: value}
}

// Uint16 will create Uint16 Field
func Uint16(name string, value uint16) Field {
	return Field{Name: name, Type: Uint16Type, Value: value}
}

// Uint32 will create Uint32 Field
func Uint32(name string, value uint32) Field {
	return Field{Name: name, Type: Uint32Type, Value: value}
}

// Uint64 will create Uint64 Field
func Uint64(name string, value uint64) Field {
	return Field{Name: name, Type: Uint64Type, Value: value}
}

// Uintptr will create Uintptr Field
func Uintptr(name string, value uintptr) Field {
	return Field{Name: name, Type: UintptrType, Value: value}
}

// Float32 will create Float32 Field
func Float32(name string, value float32) Field {
	return Field{Name: name, Type: Float32Type, Value: value}
}

// Float64 will create Float64 Field
func Float64(name string, value float64) Field {
	return Field{Name: name, Type: Float64Type, Value: value}
}

// Complex64 will create Complex64 Field
func Complex64(name string, value complex64) Field {
	return Field{Name: name, Type: Complex64Type, Value: value}
}

// Complex128 will create Complex128 Field
func Complex128(name string, value complex128) Field {
	return Field{Name: name, Type: Complex128Type, Value: value}
}

// String will create String Field
func String(name string, value string) Field {
	return Field{Name: name, Type: StringType, Value: value}
}

// Binary will create Binary Field
func Binary(name string, value []byte) Field {
	return Field{Name: name, Type: BinaryType, Value: value}
}

// ByteString will create ByteString Field
func ByteString(name string, value []byte) Field {
	return Field{Name: name, Type: ByteStringType, Value: value}
}

// Error will create Error Field
func Error(name string, value error) Field {
	return Field{Name: name, Type: ErrorType, Value: value}
}

// Stringer will create Stringer Field
func Stringer(name string, value fmt.Stringer) Field {
	return Field{Name: name, Type: StringerType, Value: value}
}

// Time will create Time Field
func Time(name string, value time.Time) Field {
	return Field{Name: name, Type: TimeType, Value: value}
}

// Duration will create Duration Field
func Duration(name string, value time.Duration) Field {
	return Field{Name: name, Type: DurationType, Value: value}
}

// Reflect will create Reflect Field
func Reflect(name string, value interface{}) Field {
	return Field{Name: name, Type: ReflectType, Value: value}
}

// Any will guess and create Field for given value
func Any(name string, value interface{}) Field {
	switch val := value.(type) {
	case bool:
		return Bool(name, val)
	case int:
		return Int64(name, int64(val))
	case int8:
		return Int8(name, val)
	case int16:
		return Int16(name, val)
	case int32:
		return Int32(name, val)
	case int64:
		return Int64(name, val)
	case uint8:
		return Uint8(name, val)
	case uint16:
		return Uint16(name, val)
	case uint32:
		return Uint32(name, val)
	case uint64:
		return Uint64(name, val)
	case uintptr:
		return Uintptr(name, val)
	case float32:
		return Float32(name, val)
	case float64:
		return Float64(name, val)
	case complex64:
		return Complex64(name, val)
	case complex128:
		return Complex128(name, val)
	case string:
		return String(name, val)
	case []byte:
		return Binary(name, val)
	case error:
		return Error(name, val)
	case time.Time:
		return Time(name, val)
	case time.Duration:
		return Duration(name, val)
	case fmt.Stringer:
		return Stringer(name, val)
	default:
		return Reflect(name, val)
	}
}
