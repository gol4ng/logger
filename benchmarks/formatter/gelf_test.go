package benchmark_formatter_test

import (
	"testing"

	"github.com/gol4ng/logger"
	"github.com/gol4ng/logger/formatter"
)

func BenchmarkGelfFormatter(b *testing.B) {
	b.ReportAllocs()

	gelfFormatter := formatter.NewGelf()
	e := logger.Entry{Message: "This log message is really logged.", Level: logger.InfoLevel}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		gelfFormatter.Format(e)
	}
}
