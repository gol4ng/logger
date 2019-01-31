package benchmark_handler

import (
	"testing"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/handler"
)

func BenchmarkGroupHandler(b *testing.B) {
	b.ReportAllocs()

	nopHandler := logger.NewNopHandler()

	groupHandler := handler.NewGroup(nopHandler, nopHandler)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		groupHandler.Handle(logger.Entry{Message: "This log message go anywhere.", Level: logger.InfoLevel})
	}
}
