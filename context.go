package logger

import (
	"fmt"
	"time"
)

type Context map[string]Field

func (c Context) Merge(context Context) Context {
	for name, field := range context {
		c.Set(name, field)
	}
	return c
}

func (c Context) Set(name string, value Field) Context {
	c[name] = value
	return c
}

func (c Context) Has(name string) bool {
	_, ok := c[name]
	return ok
}

func (c Context) Get(name string, field *Field) Field {
	if c.Has(name) {
		return c[name]
	}

	return *field
}

func (c Context) Skip(name string, value string) Context {
	return c.Set(name, Skip(value))
}
func (c Context) Bool(name string, value bool) Context {
	return c.Set(name, Bool(value))
}
func (c Context) Int8(name string, value int8) Context {
	return c.Set(name, Int8(value))
}
func (c Context) Int16(name string, value int16) Context {
	return c.Set(name, Int16(value))
}
func (c Context) Int32(name string, value int32) Context {
	return c.Set(name, Int32(value))
}
func (c Context) Int64(name string, value int64) Context {
	return c.Set(name, Int64(value))
}
func (c Context) Uint8(name string, value uint8) Context {
	return c.Set(name, Uint8(value))
}
func (c Context) Uint16(name string, value uint16) Context {
	return c.Set(name, Uint16(value))
}
func (c Context) Uint32(name string, value uint32) Context {
	return c.Set(name, Uint32(value))
}
func (c Context) Uint64(name string, value uint64) Context {
	return c.Set(name, Uint64(value))
}
func (c Context) Uintptr(name string, value uintptr) Context {
	return c.Set(name, Uintptr(value))
}
func (c Context) Float32(name string, value float32) Context {
	return c.Set(name, Float32(value))
}
func (c Context) Float64(name string, value float64) Context {
	return c.Set(name, Float64(value))
}
func (c Context) Complex64(name string, value complex64) Context {
	return c.Set(name, Complex64(value))
}
func (c Context) Complex128(name string, value complex128) Context {
	return c.Set(name, Complex128(value))
}
func (c Context) String(name string, value string) Context {
	return c.Set(name, String(value))
}
func (c Context) Binary(name string, value []byte) Context {
	return c.Set(name, Binary(value))
}
func (c Context) ByteString(name string, value []byte) Context {
	return c.Set(name, ByteString(value))
}
func (c Context) Error(name string, value error) Context {
	return c.Set(name, Error(value))
}
func (c Context) Stringer(name string, value fmt.Stringer) Context {
	return c.Set(name, Stringer(value))
}
func (c Context) Time(name string, value time.Time) Context {
	return c.Set(name, Time(value))
}
func (c Context) Duration(name string, value time.Duration) Context {
	return c.Set(name, Duration(value))
}
func (c Context) Reflect(name string, value interface{}) Context {
	return c.Set(name, Reflect(value))
}
func (c Context) Any(name string, value interface{}) Context {
	return c.Set(name, Any(value))
}

func NewContext() Context {
	return Context{}
}
