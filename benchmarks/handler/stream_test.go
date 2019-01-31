package benchmark_handler_test

import (
	"os"
	"testing"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/handler"
)

func BenchmarkStdoutStreamHandler(b *testing.B) {
	b.ReportAllocs()

	streamHandler := handler.NewStream(os.Stdout, logger.NewNopFormatter())

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		streamHandler.Handle(logger.Entry{Message: "This log message go anywhere.", Level: logger.InfoLevel})
	}
}
