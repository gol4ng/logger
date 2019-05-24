package benchmark_handler

import (
	"testing"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/handler"
)

func BenchmarkGroupHandler(b *testing.B) {
	b.ReportAllocs()

	groupHandler := handler.Group(
		logger.NopHandler,
		logger.NopHandler,
	)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		groupHandler(logger.Entry{Message: "This log message goes nowhere.", Level: logger.InfoLevel})
	}
}
