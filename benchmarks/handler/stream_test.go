package benchmark_handler_test

import (
	"testing"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/handler"
)

type NopWriter struct{}

func (w *NopWriter) Write(_ []byte) (n int, err error) {
	return 0, nil
}

func BenchmarkStdoutStreamHandler(b *testing.B) {
	b.ReportAllocs()

	streamHandler := handler.Stream(&NopWriter{}, logger.NewNopFormatter())

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = streamHandler(logger.Entry{Message: "This log message goes nowhere.", Level: logger.InfoLevel})
	}
}
