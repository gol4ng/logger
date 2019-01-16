package benchmark_formatter_test

import (
	"testing"

	"github.com/instabledesign/logger/formatter"

	"github.com/instabledesign/logger"
)

func BenchmarkLineFormatter(b *testing.B) {
	b.ReportAllocs()

	lineFormatter := formatter.NewLine("%s %s %s")

	for n := 0; n < b.N; n++ {
		lineFormatter.Format(logger.Entry{Message: "This log message is really logged.", Level: logger.InfoLevel})
	}
}
