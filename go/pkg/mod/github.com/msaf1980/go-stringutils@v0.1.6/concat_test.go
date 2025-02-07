package stringutils

import (
	"strconv"
	"testing"
)

func Benchmark_Concat3(b *testing.B) {
	i1 := 123456789
	b2 := []byte("valuestring")

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s := strconv.Itoa(i1) + ":" + UnsafeString(b2)
		_ = s
	}
}
