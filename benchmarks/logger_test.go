package benchmarks_test

import (
	"testing"

	"github.com/instabledesign/logger"
)

func BenchmarkNilLogger(b *testing.B) {
	b.ReportAllocs()

	myLogger := logger.NewNilLogger()

	for n := 0; n < b.N; n++ {
		myLogger.Info("This log message go anywhere.", nil)
	}
}
