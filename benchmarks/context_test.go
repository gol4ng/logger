package benchmarks_test

import (
	"github.com/gol4ng/logger"
	"testing"
)

func Benchmark_Context_Typed(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		logger.NewContext().
			Binary("my binary", []byte{1, 2, 3}). // 2 allocs
			ByteString("my byte string", []byte{1, 2, 3}) // 2 allocs
	}
}

func Benchmark_Context_Add(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		logger.NewContext().
			Add("my binary", []byte{1, 2, 3}). // 3 allocs
			Add("my byte string", []byte{1, 2, 3}) // 3 allocs
	}
}
