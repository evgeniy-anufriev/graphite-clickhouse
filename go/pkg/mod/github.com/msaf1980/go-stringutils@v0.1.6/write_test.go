package stringutils

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteString(t *testing.T) {
	var b bytes.Buffer
	var w io.Writer = &b

	WriteString(w, "hello world\n")
	assert.Equal(t, b.String(), "hello world\n")

	WriteString(w, "again, hello")
	assert.Equal(t, b.String(), "hello world\nagain, hello")
}

func Benchmark_io_WriteString(b *testing.B) {
	s := "test1.2.test3.4.5"

	var buf bytes.Buffer
	var w io.Writer = &buf
	buf.Grow(len(s))

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()

		_, _ = io.WriteString(w, s)
		_ = buf.String()
	}
}

func Benchmark_io_WriteStringDirect(b *testing.B) {
	s := "test1.2.test3.4.5"

	var buf bytes.Buffer
	buf.Grow(len(s))

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()

		_, _ = io.WriteString(&buf, s)
		_ = buf.String()
	}
}

func Benchmark_WriteString(b *testing.B) {
	s := "test1.2.test3.4.5"

	var buf bytes.Buffer
	var w io.Writer = &buf
	buf.Grow(len(s))

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()

		_, _ = WriteString(w, s)
		_ = buf.String()
	}
}
