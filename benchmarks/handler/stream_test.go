package benchmark_handler_test

import (
	"os"
	"testing"

	"github.com/instabledesign/logger"
	"github.com/instabledesign/logger/handler"
)

func BenchmarkNopStreamHandler(b *testing.B) {
	b.ReportAllocs()

	nopStreamHandler := handler.NewNopStream()

	for n := 0; n < b.N; n++ {
		nopStreamHandler.Handle(logger.Entry{Message: "This log message go anywhere.", Level: logger.InfoLevel})
	}
}

func BenchmarkStdoutStreamHandler(b *testing.B) {
	b.ReportAllocs()

	streamHandler := handler.NewStream(os.Stdout, logger.NewNopFormatter())

	for n := 0; n < b.N; n++ {
		streamHandler.Handle(logger.Entry{Message: "This log message go anywhere.", Level: logger.InfoLevel})
	}
}
