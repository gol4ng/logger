package benchmark_formatter_test

import (
	"testing"

	"github.com/gol4ng/logger/formatter"

	"github.com/gol4ng/logger"
)

func BenchmarkGelfFormatter(b *testing.B) {
	b.ReportAllocs()

	gelfFormatter, _ := formatter.NewGelf()
	e := logger.Entry{Message: "This log message is really logged.", Level: logger.InfoLevel}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		gelfFormatter.Format(e)
	}
}
