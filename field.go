package logger

import (
	"fmt"
	"strconv"
	"time"
)

type FieldType uint8

const (
	// UnknownType is the default field type. Attempting to add it to an encoder will panic.
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
	StringerType
	TimeType
	DurationType
	//ArrayMarshalerType
	//ObjectMarshalerType
	// ReflectType indicates that the field carries an interface{}, which should be serialized using reflection.
	ReflectType
)

type Field struct {
	Type  FieldType
	Value interface{}
}

func (f *Field) String() string {
	switch f.Type {
	case UnknownType:
		return "@TODO UnknownType"
	case SkipType:
		return "@TODO SkipType"
	case BoolType:
		return "@TODO BoolType"
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
		return "@TODO Float32Type"
	case Float64Type:
		return "@TODO Float64Type"
	case Complex64Type:
		return "@TODO Complex64Type"
	case Complex128Type:
		return "@TODO Complex128Type"
	case StringType:
		return f.Value.(string)
	case BinaryType:
		return "@TODO BinaryType"
	case ByteStringType:
		return "@TODO ByteStringType"
	case ErrorType:
		return "@TODO ErrorType"
	case StringerType:
		return "@TODO StringerType"
	case TimeType:
		return "@TODO TimeType"
	case DurationType:
		return "@TODO DurationType"
	case ReflectType:
		return "@TODO ReflectType"
	default:
		return "unknown field type"
	}
}

//func (f *Field) GoString() string {}

func Skip(value string) Field {
	return Field{Type: SkipType, Value: value}
}
func Bool(value bool) Field {
	return Field{Type: BoolType, Value: value}
}
func Int8(value int8) Field {
	return Field{Type: Int8Type, Value: value}
}
func Int16(value int16) Field {
	return Field{Type: Int16Type, Value: value}
}
func Int32(value int32) Field {
	return Field{Type: Int32Type, Value: value}
}
func Int64(value int64) Field {
	return Field{Type: Int64Type, Value: value}
}
func Uint8(value uint8) Field {
	return Field{Type: Uint8Type, Value: value}
}
func Uint16(value uint16) Field {
	return Field{Type: Uint16Type, Value: value}
}
func Uint32(value uint32) Field {
	return Field{Type: Uint32Type, Value: value}
}
func Uint64(value uint64) Field {
	return Field{Type: Uint64Type, Value: value}
}
func Uintptr(value uintptr) Field {
	return Field{Type: UintptrType, Value: value}
}
func Float32(value float32) Field {
	return Field{Type: Float32Type, Value: value}
}
func Float64(value float64) Field {
	return Field{Type: Float64Type, Value: value}
}
func Complex64(value complex64) Field {
	return Field{Type: Complex64Type, Value: value}
}
func Complex128(value complex128) Field {
	return Field{Type: Complex128Type, Value: value}
}
func String(value string) Field {
	return Field{Type: StringType, Value: value}
}
func Binary(value []byte) Field {
	return Field{Type: BinaryType, Value: value}
}
func ByteString(value []byte) Field {
	return Field{Type: ByteStringType, Value: value}
}
func Error(value error) Field {
	return Field{Type: ErrorType, Value: value}
}
func Stringer(value fmt.Stringer) Field {
	return Field{Type: StringerType, Value: value}
}
func Time(value time.Time) Field {
	return Field{Type: TimeType, Value: value}
}
func Duration(value time.Duration) Field {
	return Field{Type: DurationType, Value: value}
}
func Reflect(value interface{}) Field {
	return Field{Type: ReflectType, Value: value}
}

func Any(value interface{}) Field {
	switch val := value.(type) {
	case bool:
		return Bool(val)
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
	case fmt.Stringer:
		return Stringer(val)
	case time.Time:
		return Time(val)
	case time.Duration:
		return Duration(val)
	default:
		return Reflect(val)
	}
}
