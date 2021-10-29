package logger_test

import (
	"errors"
	"testing"
	"time"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/gol4ng/logger/handler"
	"github.com/gol4ng/logger/middleware"
)

type NopWriter struct{}

func (w *NopWriter) Write(_ []byte) (n int, err error) {
	return 0, nil
}

type loggerToTest struct {
	name   string
	logger logger.LoggerInterface
}

func loggersToBench() []loggerToTest {
	return []loggerToTest{
		{
			"nop",
			logger.NewNopLogger(),
		},
		{
			"stream line",
			logger.NewLogger(
				handler.Stream(&NopWriter{}, formatter.NewLine("%[2]s | %[1]s")),
			),
		},
		{
			"min level filter",
			logger.NewLogger(middleware.MinLevelFilter(logger.WarningLevel)(logger.NopHandler)),
		},
		{
			"max level filter",
			logger.NewLogger(middleware.MaxLevelFilter(logger.WarningLevel)(logger.NopHandler)),
		},
		{
			"range level filter",
			logger.NewLogger(middleware.RangeLevelFilter(logger.WarningLevel, logger.InfoLevel)(logger.NopHandler)),
		},
		{
			"placeholder",
			logger.NewLogger(middleware.Placeholder()(logger.NopHandler)),
		},
	}
}

func getEntries() []logger.Entry {
	return []logger.Entry{
		{
			Message: "test %my_key% message",
			Level:   logger.WarningLevel,
			Context: logger.NewContext().
				Add("my_key", "my_value").
				Add("my_key2", "my_value2").
				Add("my_key3", "my_value3").
				Add("my_key4", "my_value4").
				Add("my_key5", "my_value5"),
		},
	}
}

func BenchmarkMiddleware(b *testing.B) {
	for _, entry := range getEntries() {
		for _, m := range loggersToBench() {
			b.Run(m.name+"_"+entry.Message, func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					m.logger.Log(entry.Message, entry.Level, *(entry.Context.Slice())...)
				}
			})
		}
	}
}

func BenchmarkContext_Add(b *testing.B) {
	tests := []struct {
		name string
		data interface{}
	}{
		{name: "nil", data: nil},
		{name: "boolean", data: false},
		{name: "int", data: 1234},
		{name: "int8", data: int8(123)},
		{name: "int16", data: int16(1234)},
		{name: "int32", data: int32(1234)},
		{name: "int64", data: int64(1234)},
		{name: "uint8", data: uint8(123)},
		{name: "uint16", data: uint16(1234)},
		{name: "uint32", data: uint32(1234)},
		{name: "uint64", data: uint64(1234)},
		{name: "uintptr", data: uintptr(1234)},
		{name: "float32", data: float32(1234.56)},
		{name: "float64", data: float64(1234.56)},
		{name: "complex64", data: complex64(123)},
		{name: "complex128", data: complex128(123)},
		{name: "string", data: "example"},
		{name: "[]byte", data: []byte("example")},
		{name: "error", data: errors.New("example")},
		{name: "time", data: time.Now()},
		{name: "duration", data: time.Second},
		{name: "stringer", data: MyStringer{}},
		{name: "reflect", data: struct{}{}},
	}
	ctx := logger.NewContext()

	for _, tt := range tests {
		b.Run("context_add_"+tt.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ctx.Add("my_message", tt.data)
			}
		})
	}
}
