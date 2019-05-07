package benchmarks_test

import (
	"github.com/gol4ng/logger/middleware"
	"testing"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/gol4ng/logger/handler"
)

type NopWriter struct{}

func (w *NopWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

func BenchmarkNopLogger(b *testing.B) {
	b.ReportAllocs()

	myLogger := logger.NewNopLogger()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		myLogger.Info("This log message go anywhere.", nil)
	}
}

func BenchmarkLoggerLineFormatter(b *testing.B) {
	b.ReportAllocs()

	myLogger := logger.NewLogger(
		handler.Stream(&NopWriter{}, formatter.NewLine("%[2]s | %[1]s")),
	)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		myLogger.Info("Log example", nil)
	}
}

func BenchmarkLoggerJsonFormatter(b *testing.B) {
	b.ReportAllocs()

	myLogger := logger.NewLogger(
		handler.Stream(&NopWriter{}, formatter.NewJsonEncoder()),
	)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		myLogger.Info("Log example", nil)
	}
}

func BenchmarkLoggerMinLevelFilterHandler(b *testing.B) {
	b.ReportAllocs()

	myLogger := logger.NewLogger(
		middleware.MinLevelFilter(logger.InfoLevel)(logger.NopHandler),
	)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		myLogger.Debug("Log example", nil)
		myLogger.Info("Log example", nil)
	}
}

func BenchmarkLoggerGroupHandler(b *testing.B) {
	b.ReportAllocs()

	myLogger := logger.NewLogger(handler.Group(
		handler.Stream(&NopWriter{}, &logger.NopFormatter{}),
		handler.Stream(&NopWriter{}, &logger.NopFormatter{}),
	))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		myLogger.Debug("Log example", nil)
	}
}
