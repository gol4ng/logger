package benchmark_formatter_test

import (
	"testing"

	"github.com/instabledesign/logger/formatter"

	"github.com/instabledesign/logger"
)

func BenchmarkJsonFormatter(b *testing.B) {
	b.ReportAllocs()

	jsonFormatter := formatter.NewJson()

	for n := 0; n < b.N; n++ {
		jsonFormatter.Format(logger.Entry{Message: "This log message is really logged.", Level: logger.InfoLevel})
	}
}
