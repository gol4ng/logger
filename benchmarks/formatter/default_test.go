package benchmark_formatter_test

import (
	"testing"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
)

func BenchmarkDefaultFormatter(b *testing.B) {
	b.ReportAllocs()

	jsonFormatter := formatter.NewJSONEncoder()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		jsonFormatter.Format(logger.Entry{Message: "This log message is really logged.", Level: logger.InfoLevel})
	}
}
