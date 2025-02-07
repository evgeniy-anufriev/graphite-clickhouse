package stringutils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TrimRight(t *testing.T) {
	t.Parallel()
	res := TrimRight("/test//////", '/')
	assert.Equal(t, "/test", res)

	res = TrimRight("/test", '/')
	assert.Equal(t, "/test", res)

	res = TrimRight(" ", ' ')
	assert.Equal(t, "", res)

	res = TrimRight("  ", ' ')
	assert.Equal(t, "", res)

	res = TrimRight("", ' ')
	assert.Equal(t, "", res)
}

func Benchmark_TrimRight(b *testing.B) {
	var res string

	b.Run("stringutils", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = TrimRight("foobar  ", ' ')
		}
		assert.Equal(b, "foobar", res)
	})
	b.Run("stdlib", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = strings.TrimRight("foobar  ", " ")
		}
		assert.Equal(b, "foobar", res)
	})
}

func Test_TrimLeft(t *testing.T) {
	t.Parallel()
	res := TrimLeft("////test/", '/')
	assert.Equal(t, "test/", res)

	res = TrimLeft("test/", '/')
	assert.Equal(t, "test/", res)

	res = TrimLeft(" ", ' ')
	assert.Equal(t, "", res)

	res = TrimLeft("  ", ' ')
	assert.Equal(t, "", res)

	res = TrimLeft("", ' ')
	assert.Equal(t, "", res)
}

func Benchmark_TrimLeft(b *testing.B) {
	var res string

	b.Run("stringutils", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = TrimLeft("  foobar", ' ')
		}
		assert.Equal(b, "foobar", res)
	})
	b.Run("stdlib", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = strings.TrimLeft("  foobar", " ")
		}
		assert.Equal(b, "foobar", res)
	})
}

func Test_Trim(t *testing.T) {
	t.Parallel()
	res := Trim("   test  ", ' ')
	assert.Equal(t, "test", res)

	res = Trim("test", ' ')
	assert.Equal(t, "test", res)

	res = Trim(".test", '.')
	assert.Equal(t, "test", res)

	res = Trim(" ", ' ')
	assert.Equal(t, "", res)

	res = Trim("  ", ' ')
	assert.Equal(t, "", res)

	res = Trim("", ' ')
	assert.Equal(t, "", res)
}

func Benchmark_Trim(b *testing.B) {
	var res string

	b.Run("stringutils", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = Trim("  foobar   ", ' ')
		}
		assert.Equal(b, "foobar", res)
	})
	b.Run("stdlib", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = strings.Trim("  foobar   ", " ")
		}
		assert.Equal(b, "foobar", res)
	})
	b.Run("stdlib.trimspace", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = strings.TrimSpace("  foobar   ")
		}
		assert.Equal(b, "foobar", res)
	})
}
