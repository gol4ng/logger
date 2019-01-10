package benchmarks_test

import (
	"os"
	"testing"

	"github.com/instabledesign/logger"
	"github.com/instabledesign/logger/handler"
)

func BenchmarkNilStreamHandler(b *testing.B) {
	b.ReportAllocs()

	myLogger := logger.NewLogger(handler.NewNilStream())

	for n := 0; n < b.N; n++ {
		myLogger.Info("This log message go anywhere.", nil)
	}
}

func BenchmarkStdoutStreamHandler(b *testing.B) {
	b.ReportAllocs()

	myHandler := handler.NewStream(os.Stdout, logger.NewNilFormatter())
	myLogger := logger.NewLogger(myHandler)

	for n := 0; n < b.N; n++ {
		myLogger.Info("This log message go anywhere.", nil)
	}
}
