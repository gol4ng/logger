package benchmarks_test

import (
	"os"
	"testing"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
	"github.com/gol4ng/logger/handler"
)

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

	lineLogHandler := handler.NewStream(os.Stdout, formatter.NewLine("%[2]s | %[1]s"))
	myLogger := logger.NewLogger(lineLogHandler)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		myLogger.Info("Log example", nil)
	}
}

func BenchmarkLoggerJsonFormatter(b *testing.B) {
	b.ReportAllocs()

	lineLogHandler := handler.NewStream(os.Stdout, formatter.NewJsonEncoder())
	myLogger := logger.NewLogger(lineLogHandler)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		myLogger.Info("Log example", nil)
	}
}

func BenchmarkLoggerMinLevelFilterHandler(b *testing.B) {
	b.ReportAllocs()

	lineLogHandler := handler.NewStream(os.Stdout, &logger.NopFormatter{})
	filterLogHandler := handler.NewMinLevelFilter(lineLogHandler, logger.InfoLevel)
	myLogger := logger.NewLogger(filterLogHandler)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		myLogger.Debug("Log example", nil)
		myLogger.Info("Log example", nil)
	}
}

func BenchmarkLoggerGroupHandler(b *testing.B) {
	b.ReportAllocs()

	jsonLogHandler := handler.NewStream(os.Stdout, &formatter.Json{})
	lineLogHandler := handler.NewStream(os.Stdout, &logger.NopFormatter{})
	groupLogHandler := handler.NewGroup(jsonLogHandler, lineLogHandler)
	myLogger := logger.NewLogger(groupLogHandler)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		myLogger.Debug("Log example", nil)
	}
}
