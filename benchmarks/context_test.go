package benchmarks_test

import (
	"testing"

	"github.com/gol4ng/logger"
)

func Benchmark_Context_Typed(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		logger.NewContext().
			Binary("my binary", []byte{1, 2, 3}). // 2 allocations
			ByteString("my byte string", []byte{1, 2, 3}) // 2 allocations
	}
}

func Benchmark_Context_Add(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		logger.NewContext().
			Add("my binary", []byte{1, 2, 3}). // 3 allocations
			Add("my byte string", []byte{1, 2, 3}) // 3 allocations
	}
}
