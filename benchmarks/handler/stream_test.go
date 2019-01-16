package benchmark_handler_test

import (
	"os"
	"testing"

	"github.com/instabledesign/logger"
	"github.com/instabledesign/logger/handler"
)

func BenchmarkNilStreamHandler(b *testing.B) {
	b.ReportAllocs()

	nilStreamHandler := handler.NewNilStream()

	for n := 0; n < b.N; n++ {
		nilStreamHandler.Handle(logger.Entry{Message: "This log message go anywhere.", Level: logger.InfoLevel})
	}
}

func BenchmarkStdoutStreamHandler(b *testing.B) {
	b.ReportAllocs()

	streamHandler := handler.NewStream(os.Stdout, logger.NewNilFormatter())

	for n := 0; n < b.N; n++ {
		streamHandler.Handle(logger.Entry{Message: "This log message go anywhere.", Level: logger.InfoLevel})
	}
}
