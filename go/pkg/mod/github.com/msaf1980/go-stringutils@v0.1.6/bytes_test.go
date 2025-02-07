package stringutils

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ToLowerBytes(t *testing.T) {
	t.Parallel()
	res := ToLowerBytes([]byte("/MY/NAME/IS/:PARAM/*"))
	assert.Equal(t, true, bytes.Equal([]byte("/my/name/is/:param/*"), res))
	res = ToLowerBytes([]byte("/MY1/NAME/IS/:PARAM/*"))
	assert.Equal(t, true, bytes.Equal([]byte("/my1/name/is/:param/*"), res))
	res = ToLowerBytes([]byte("/MY2/NAME/IS/:PARAM/*"))
	assert.Equal(t, true, bytes.Equal([]byte("/my2/name/is/:param/*"), res))
	res = ToLowerBytes([]byte("/MY3/NAME/IS/:PARAM/*"))
	assert.Equal(t, true, bytes.Equal([]byte("/my3/name/is/:param/*"), res))
	res = ToLowerBytes([]byte("/MY4/NAME/IS/:PARAM/*"))
	assert.Equal(t, true, bytes.Equal([]byte("/my4/name/is/:param/*"), res))
}

func Benchmark_ToLowerBytes(b *testing.B) {
	path := []byte(largeStr)
	want := []byte(lowerStr)
	var res []byte
	b.Run("fiber", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = ToLowerBytes(path)
		}
		assert.Equal(b, bytes.Equal(want, res), true)
	})
	b.Run("default", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = bytes.ToLower(path)
		}
		assert.Equal(b, bytes.Equal(want, res), true)
	})
}

func Test_ToUpperBytes(t *testing.T) {
	t.Parallel()
	res := ToUpperBytes([]byte("/my/name/is/:param/*"))
	assert.Equal(t, true, bytes.Equal([]byte("/MY/NAME/IS/:PARAM/*"), res))
	res = ToUpperBytes([]byte("/my1/name/is/:param/*"))
	assert.Equal(t, true, bytes.Equal([]byte("/MY1/NAME/IS/:PARAM/*"), res))
	res = ToUpperBytes([]byte("/my2/name/is/:param/*"))
	assert.Equal(t, true, bytes.Equal([]byte("/MY2/NAME/IS/:PARAM/*"), res))
	res = ToUpperBytes([]byte("/my3/name/is/:param/*"))
	assert.Equal(t, true, bytes.Equal([]byte("/MY3/NAME/IS/:PARAM/*"), res))
	res = ToUpperBytes([]byte("/my4/name/is/:param/*"))
	assert.Equal(t, true, bytes.Equal([]byte("/MY4/NAME/IS/:PARAM/*"), res))
}

func Benchmark_ToUpperBytes(b *testing.B) {
	path := []byte(largeStr)
	want := []byte(upperStr)
	var res []byte
	b.Run("fiber", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = ToUpperBytes(path)
		}
		assert.Equal(b, bytes.Equal(want, res), true)
	})
	b.Run("default", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = bytes.ToUpper(path)
		}
		assert.Equal(b, bytes.Equal(want, res), true)
	})
}

func Test_TrimRightBytes(t *testing.T) {
	t.Parallel()
	res := TrimRightBytes([]byte("/test//////"), '/')
	assert.Equal(t, []byte("/test"), res)

	res = TrimRightBytes([]byte("/test"), '/')
	assert.Equal(t, []byte("/test"), res)

	res = TrimRightBytes([]byte(" "), ' ')
	assert.Equal(t, 0, len(res))

	res = TrimRightBytes([]byte("  "), ' ')
	assert.Equal(t, 0, len(res))

	res = TrimRightBytes([]byte(""), ' ')
	assert.Equal(t, 0, len(res))
}

func Benchmark_TrimRightBytes(b *testing.B) {
	var res []byte

	b.Run("fiber", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = TrimRightBytes([]byte("foobar  "), ' ')
		}
		assert.Equal(b, []byte("foobar"), res)
	})
	b.Run("default", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = bytes.TrimRight([]byte("foobar  "), " ")
		}
		assert.Equal(b, []byte("foobar"), res)
	})
}

func Test_TrimLeftBytes(t *testing.T) {
	t.Parallel()
	res := TrimLeftBytes([]byte("////test/"), '/')
	assert.Equal(t, []byte("test/"), res)

	res = TrimLeftBytes([]byte("test/"), '/')
	assert.Equal(t, []byte("test/"), res)

	res = TrimLeftBytes([]byte(" "), ' ')
	assert.Equal(t, 0, len(res))

	res = TrimLeftBytes([]byte("  "), ' ')
	assert.Equal(t, 0, len(res))

	res = TrimLeftBytes([]byte(""), ' ')
	assert.Equal(t, 0, len(res))
}

func Benchmark_TrimLeftBytes(b *testing.B) {
	var res []byte

	b.Run("fiber", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = TrimLeftBytes([]byte("  foobar"), ' ')
		}
		assert.Equal(b, []byte("foobar"), res)
	})
	b.Run("default", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = bytes.TrimLeft([]byte("  foobar"), " ")
		}
		assert.Equal(b, []byte("foobar"), res)
	})
}

func Test_TrimBytes(t *testing.T) {
	t.Parallel()
	res := TrimBytes([]byte("   test  "), ' ')
	assert.Equal(t, []byte("test"), res)

	res = TrimBytes([]byte("test"), ' ')
	assert.Equal(t, []byte("test"), res)

	res = TrimBytes([]byte(".test"), '.')
	assert.Equal(t, []byte("test"), res)

	res = TrimBytes([]byte(" "), ' ')
	assert.Equal(t, 0, len(res))

	res = TrimBytes([]byte("  "), ' ')
	assert.Equal(t, 0, len(res))

	res = TrimBytes([]byte(""), ' ')
	assert.Equal(t, 0, len(res))
}

func Benchmark_TrimBytes(b *testing.B) {
	var res []byte

	b.Run("fiber", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = TrimBytes([]byte("  foobar   "), ' ')
		}
		assert.Equal(b, []byte("foobar"), res)
	})
	b.Run("default", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = bytes.Trim([]byte("  foobar   "), " ")
		}
		assert.Equal(b, []byte("foobar"), res)
	})
}

func Benchmark_EqualFoldBytes(b *testing.B) {
	left := []byte(upperStr)
	right := []byte(lowerStr)
	var res bool
	b.Run("fiber", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = EqualFoldBytes(left, right)
		}
		assert.Equal(b, true, res)
	})
	b.Run("default", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = bytes.EqualFold(left, right)
		}
		assert.Equal(b, true, res)
	})
}

func Test_EqualFoldBytes(t *testing.T) {
	t.Parallel()
	res := EqualFoldBytes([]byte("/MY/NAME/IS/:PARAM/*"), []byte("/my/name/is/:param/*"))
	assert.Equal(t, true, res)
	res = EqualFoldBytes([]byte("/MY1/NAME/IS/:PARAM/*"), []byte("/MY1/NAME/IS/:PARAM/*"))
	assert.Equal(t, true, res)
	res = EqualFoldBytes([]byte("/my2/name/is/:param/*"), []byte("/my2/name"))
	assert.Equal(t, false, res)
	res = EqualFoldBytes([]byte("/dddddd"), []byte("eeeeee"))
	assert.Equal(t, false, res)
	res = EqualFoldBytes([]byte("\na"), []byte("*A"))
	assert.Equal(t, false, res)
	res = EqualFoldBytes([]byte("/MY3/NAME/IS/:PARAM/*"), []byte("/my3/name/is/:param/*"))
	assert.Equal(t, true, res)
	res = EqualFoldBytes([]byte("/MY4/NAME/IS/:PARAM/*"), []byte("/my4/nAME/IS/:param/*"))
	assert.Equal(t, true, res)
}
