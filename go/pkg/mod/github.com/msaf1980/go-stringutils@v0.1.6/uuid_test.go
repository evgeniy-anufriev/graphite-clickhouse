package stringutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UUID(t *testing.T) {
	t.Parallel()
	res := UUID()
	assert.Equal(t, 36, len(res))
	assert.Equal(t, true, res != emptyUUID)
}

func Test_UUID_Concurrency(t *testing.T) {
	t.Parallel()
	iterations := 1000
	var res string
	ch := make(chan string, iterations)
	results := make(map[string]string)
	for i := 0; i < iterations; i++ {
		go func() {
			ch <- UUID()
		}()
	}
	for i := 0; i < iterations; i++ {
		res = <-ch
		results[res] = res
	}
	assert.Equal(t, iterations, len(results))
}

func Benchmark_YourFunc(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = UUID()
		}
	})
}
