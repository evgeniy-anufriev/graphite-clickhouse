package stringutils

import (
	"reflect"
	"testing"
)

func TestUnsafeString(t *testing.T) {
	s := "test"
	bytes := []byte(s)
	if got := UnsafeString(bytes); got != s {
		t.Errorf("UnsafeString() = '%s', want '%s'", got, s)
	}
}

func TestString(t *testing.T) {
	s := "test"
	bytes := []byte(s)
	if got := String(&bytes[0], len(bytes)); got != s {
		t.Errorf("String()) = '%s', want '%s'", got, s)
	}
}

func TestUnsafeStringBytes(t *testing.T) {
	bytes := []byte("test")
	s := string(bytes)
	if got := UnsafeStringBytes(&s); !reflect.DeepEqual(got, bytes) {
		t.Errorf("UnsafeStringBytes() = '%v', want '%v'", got, bytes)
	}
}

func TestStringData(t *testing.T) {
	bytes := []byte("test")
	s := string(bytes)
	length := len(s)
	got := StringData(s)
	gotBytes := UnsafeBytes(got, length, length)
	if !reflect.DeepEqual(gotBytes, bytes) {
		t.Errorf("UnsafeStringBytes() = '%v' (%q), want '%v' (%q)", gotBytes, string(gotBytes), bytes, string(bytes))
	}
}

func Benchmark_UnsafeString(b *testing.B) {
	bytes := []byte("test1.2.test3.4.5")
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = UnsafeString(bytes)
	}
}

func Benchmark_String(b *testing.B) {
	s := "test1.2.test3.4.5"
	bytes := []byte(s)
	ptr := &bytes[0]

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = String(ptr, len(bytes))
	}
}

func Benchmark_UnsafeStringBytes(b *testing.B) {
	s := "test1.2.test3.4.5"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = UnsafeStringBytes(&s)
	}
}

func Benchmark_UnsafeStringBytePtr(b *testing.B) {
	s := "test1.2.test3.4.5"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = UnsafeStringBytePtr(s)
	}
}
