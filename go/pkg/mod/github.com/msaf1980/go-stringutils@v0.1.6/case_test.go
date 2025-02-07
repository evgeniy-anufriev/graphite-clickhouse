package stringutils

import (
	"strings"
	"testing"
	"unicode"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
)

func tenRunes(ch rune) string {
	r := make([]rune, 10)
	for i := range r {
		r[i] = ch
	}

	return string(r)
}

// User-defined self-inverse mapping function

func rot13(r rune) rune {
	step := rune(13)
	if r >= 'a' && r <= 'z' {
		return ((r - 'a' + step) % 26) + 'a'
	}

	if r >= 'A' && r <= 'Z' {
		return ((r - 'A' + step) % 26) + 'A'
	}

	return r
}

func TestMap(t *testing.T) {
	// Run a couple of awful growth/shrinkage tests
	var sb Builder

	a := tenRunes('a')

	// 1.  Grow. This triggers two reallocations in Map.
	maxRune := func(rune) rune { return unicode.MaxRune }
	sb.Map(maxRune, a)

	expect := tenRunes(unicode.MaxRune)
	assert.Equalf(t, expect, sb.String(), "growing")

	// 2. Shrink
	sb.Release()
	minRune := func(rune) rune { return 'a' }
	sb.Map(minRune, tenRunes(unicode.MaxRune))

	expect = a
	assert.Equalf(t, expect, sb.String(), "shrinking")

	// 3. Rot13
	sb.Release()
	sb.Map(rot13, "a to zed")

	expect = "n gb mrq"

	assert.Equalf(t, expect, sb.String(), "rot13")

	// 4. Rot13^2
	m := sb.String()
	sb.Release()

	sb.Map(rot13, m)

	expect = "a to zed"
	assert.Equalf(t, expect, sb.String(), "rot13")

	// 5. Drop
	sb.Release()

	dropNotLatin := func(r rune) rune {
		if unicode.Is(unicode.Latin, r) {
			return r
		}

		return -1
	}

	sb.Map(dropNotLatin, "Hello, 세계")

	expect = "Hello"
	assert.Equalf(t, expect, sb.String(), "drop")

	// 6. Identity
	sb.Release()

	identity := func(r rune) rune {
		return r
	}

	orig := "Input string that we expect not to be copied."

	sb.Map(identity, orig)
	assert.Equalf(t, orig, sb.String(), "unexpected copy")

	// 7. Handle invalid UTF-8 sequence
	sb.Release()

	replaceNotLatin := func(r rune) rune {
		if unicode.Is(unicode.Latin, r) {
			return r
		}

		return utf8.RuneError
	}

	sb.Map(replaceNotLatin, "Hello\255World")

	expect = "Hello\uFFFDWorld"
	assert.Equalf(t, expect, sb.String(), "replace invalid sequence")

	// 8. Check utf8.RuneSelf and utf8.MaxRune encoding
	sb.Release()

	encode := func(r rune) rune {
		switch r {
		case utf8.RuneSelf:
			return unicode.MaxRune
		case unicode.MaxRune:
			return utf8.RuneSelf
		}

		return r
	}

	s := string(rune(utf8.RuneSelf)) + string(utf8.MaxRune)
	r := string(utf8.MaxRune) + string(rune(utf8.RuneSelf)) // reverse of s

	sb.Map(encode, s)
	assert.Equalf(t, r, sb.String(), "encoding not handled correctly")

	sb.Release()
	sb.Map(encode, r)

	assert.Equalf(t, s, sb.String(), "encoding not handled correctly")

	// 9. Check mapping occurs in the front, middle and back
	sb.Release()

	trimSpaces := func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}

		return r
	}

	sb.Map(trimSpaces, "   abc    123   ")

	expect = "abc123"
	assert.Equalf(t, expect, sb.String(), "trimSpaces")
}

// Test case for any function which accepts and returns a single string.
type StringTest struct {
	in, out string
}

type StringOper int16

const (
	StringUpper StringOper = iota
	StringLower
)

// Execute f on each test case.  funcName should be the name of f; it's used
// in failure reports.
func runBuilderTests(t *testing.T, stringOper StringOper, testCases []StringTest) {
	for _, tc := range testCases {
		t.Run(tc.in, func(t *testing.T) {
			var sb Builder
			var funcName string
			switch stringOper {
			case StringUpper:
				sb.WriteStringUpper(tc.in)
				funcName = "Builder.WriteStringUpper"
			case StringLower:
				sb.WriteStringLower(tc.in)
				funcName = "Builder.WriteStringLower"
			}
			assert.Equalf(t, tc.out, sb.String(), "%s", funcName)
		})
	}
}

var upperTests = []StringTest{
	{"", ""},
	{"ONLYUPPER", "ONLYUPPER"},
	{"abc", "ABC"},
	{"AbC123", "ABC123"},
	{"azAZ09_", "AZAZ09_"},
	{"longStrinGwitHmixofsmaLLandcAps", "LONGSTRINGWITHMIXOFSMALLANDCAPS"},
	{"long\u0250string\u0250with\u0250nonascii\u2C6Fchars", "LONG\u2C6FSTRING\u2C6FWITH\u2C6FNONASCII\u2C6FCHARS"},
	{"\u0250\u0250\u0250\u0250\u0250", "\u2C6F\u2C6F\u2C6F\u2C6F\u2C6F"}, // grows one byte per char
	{"a\u0080\U0010FFFF", "A\u0080\U0010FFFF"},                           // test utf8.RuneSelf and utf8.MaxRune
}

var lowerTests = []StringTest{
	{"", ""},
	{"abc", "abc"},
	{"AbC123", "abc123"},
	{"azAZ09_", "azaz09_"},
	{"longStrinGwitHmixofsmaLLandcAps", "longstringwithmixofsmallandcaps"},
	{"LONG\u2C6FSTRING\u2C6FWITH\u2C6FNONASCII\u2C6FCHARS", "long\u0250string\u0250with\u0250nonascii\u0250chars"},
	{"\u2C6D\u2C6D\u2C6D\u2C6D\u2C6D", "\u0251\u0251\u0251\u0251\u0251"}, // shrinks one byte per char
	{"A\u0080\U0010FFFF", "a\u0080\U0010FFFF"},                           // test utf8.RuneSelf and utf8.MaxRune
}

func TestWriteStringUpper(t *testing.T) { runBuilderTests(t, StringUpper, upperTests) }

func TestWriteStringToLower(t *testing.T) { runBuilderTests(t, StringLower, lowerTests) }

func Test_ToUpper(t *testing.T) {
	t.Parallel()
	res := ToUpper("/my/name/is/:param/*")
	assert.Equal(t, "/MY/NAME/IS/:PARAM/*", res)
}

const (
	largeStr = "/RePos/GoFiBer/FibEr/iSsues/187643/CoMmEnts/RePos/GoFiBer/FibEr/iSsues/CoMmEnts"
	upperStr = "/REPOS/GOFIBER/FIBER/ISSUES/187643/COMMENTS/REPOS/GOFIBER/FIBER/ISSUES/COMMENTS"
	lowerStr = "/repos/gofiber/fiber/issues/187643/comments/repos/gofiber/fiber/issues/comments"
)

func Benchmark_ToUpper(b *testing.B) {
	var res string
	b.Run("stringutils", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = ToUpper(largeStr)
		}
		assert.Equal(b, upperStr, res)
	})
	b.Run("stdlib", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = strings.ToUpper(largeStr)
		}
		assert.Equal(b, upperStr, res)
	})
}

func Test_ToLower(t *testing.T) {
	t.Parallel()
	res := ToLower("/MY/NAME/IS/:PARAM/*")
	assert.Equal(t, "/my/name/is/:param/*", res)
	res = ToLower("/MY1/NAME/IS/:PARAM/*")
	assert.Equal(t, "/my1/name/is/:param/*", res)
	res = ToLower("/MY2/NAME/IS/:PARAM/*")
	assert.Equal(t, "/my2/name/is/:param/*", res)
	res = ToLower("/MY3/NAME/IS/:PARAM/*")
	assert.Equal(t, "/my3/name/is/:param/*", res)
	res = ToLower("/MY4/NAME/IS/:PARAM/*")
	assert.Equal(t, "/my4/name/is/:param/*", res)
}

func Benchmark_ToLower(b *testing.B) {
	var res string
	b.Run("stringutils", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = ToLower(largeStr)
		}
		assert.Equal(b, lowerStr, res)
	})
	b.Run("stdlib", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = strings.ToLower(largeStr)
		}
		assert.Equal(b, lowerStr, res)
	})
}

func Test_EqualFold(t *testing.T) {
	t.Parallel()
	res := EqualFold("/MY/NAME/IS/:PARAM/*", "/my/name/is/:param/*")
	assert.Equal(t, true, res)
	res = EqualFold("/MY1/NAME/IS/:PARAM/*", "/MY1/NAME/IS/:PARAM/*")
	assert.Equal(t, true, res)
	res = EqualFold("/my2/name/is/:param/*", "/my2/name")
	assert.Equal(t, false, res)
	res = EqualFold("/dddddd", "eeeeee")
	assert.Equal(t, false, res)
	res = EqualFold("\na", "*A")
	assert.Equal(t, false, res)
	res = EqualFold("/MY3/NAME/IS/:PARAM/*", "/my3/name/is/:param/*")
	assert.Equal(t, true, res)
	res = EqualFold("/MY4/NAME/IS/:PARAM/*", "/my4/nAME/IS/:param/*")
	assert.Equal(t, true, res)
}

// go test -v -run=^$ -bench=Benchmark_EqualFold -benchmem -count=4
func Benchmark_EqualFold(b *testing.B) {
	var res bool
	b.Run("stringutils", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = EqualFold(upperStr, lowerStr)
		}
		assert.Equal(b, true, res)
	})
	b.Run("stdlib", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			res = strings.EqualFold(upperStr, lowerStr)
		}
		assert.Equal(b, true, res)
	})
}
