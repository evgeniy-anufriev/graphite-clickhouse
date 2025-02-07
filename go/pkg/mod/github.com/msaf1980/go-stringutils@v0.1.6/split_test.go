package stringutils

import (
	"reflect"
	"strings"
	"testing"
)

func Test_Split2(t *testing.T) {
	tests := []struct {
		s           string
		sep         string
		want0       string
		want1       string
		wantStrings int
	}{
		{"", "", "", "", 1},
		{"", "&", "", "", 1},
		{"test", "&", "test", "", 1},
		{"test&", "&", "test", "", 2},
		{"test&after", "&", "test", "after", 2},
		{"test=&after", "=&", "test", "after", 2},
		{"тестПроверкАпосле", "ПроверкА", "тест", "после", 2},
	}
	for _, tt := range tests {
		t.Run(tt.s+" -> "+tt.sep, func(t *testing.T) {
			s0, s1, n := Split2(tt.s, tt.sep)
			if s0 != tt.want0 {
				t.Errorf("Split2() s[0] = %v, want %v", s0, tt.want0)
			}
			if s1 != tt.want1 {
				t.Errorf("Split2() s[1] = %v, want %v", s1, tt.want1)
			}
			if n != tt.wantStrings {
				t.Errorf("Split2() count = %v, want %v", n, tt.wantStrings)
			}
		})
	}
}

// for compare with Benchmark_Split2
func Benchmark_Strings_Split2(b *testing.B) {
	s := "teststring&where"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = strings.SplitN(s, "&", 1)
	}
}

func Benchmark_Split2(b *testing.B) {
	s := "teststring&where"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _, _ = Split2(s, "&")
	}
}

func TestSplit(t *testing.T) {
	buf := make([]string, 4)
	tests := []struct {
		s    string
		sep  string
		want []string
	}{
		{"", ",", []string{""}},
		{"test", ",", []string{"test"}},
		{"test1,2", ",", []string{"test1", "2"}},
		{"test1.2.test3.", ".", []string{"test1", "2", "test3", ""}},
		{"test1=.2.test3.", "=.", []string{"test1", "2.test3."}},
		{"test1.2.test3.4", ".", []string{"test1", "2", "test3", "4"}},
		{"test1.2.test3.4.", ".", []string{"test1", "2", "test3", "4", ""}},
		{"test1.2.test3.4.5", ".", []string{"test1", "2", "test3", "4", "5"}},
		{"тестПА1ПА2Пк3ПА4ПА5", "ПА", []string{"тест", "1", "2Пк3", "4", "5"}},
	}
	for _, tt := range tests {
		t.Run(tt.s+" -> "+tt.sep, func(t *testing.T) {
			got := Split(tt.s, tt.sep, buf)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitN() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestSplitByte(t *testing.T) {
	buf := make([]string, 4)
	tests := []struct {
		s    string
		sep  byte
		want []string
	}{
		{"", ',', []string{""}},
		{"test", ',', []string{"test"}},
		{"test1,2", ',', []string{"test1", "2"}},
		{"test1.2.test3.", '.', []string{"test1", "2", "test3", ""}},
		{"test1.2.test3.4", '.', []string{"test1", "2", "test3", "4"}},
		{"test1.2.test3.4.", '.', []string{"test1", "2", "test3", "4", ""}},
		{"test1.2.test3.4.5", '.', []string{"test1", "2", "test3", "4", "5"}},
	}
	for _, tt := range tests {
		t.Run(tt.s+" -> "+string(tt.sep), func(t *testing.T) {
			got := SplitByte(tt.s, tt.sep, buf)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitN() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestSplitRune(t *testing.T) {
	buf := make([]string, 4)
	tests := []struct {
		s    string
		sep  rune
		want []string
	}{
		{"", ',', []string{""}},
		{"test", ',', []string{"test"}},
		{"test1,2", ',', []string{"test1", "2"}},
		{"test1.2.test3.", '.', []string{"test1", "2", "test3", ""}},
		{"test1.2.test3.4", '.', []string{"test1", "2", "test3", "4"}},
		{"test1.2.test3.4.", '.', []string{"test1", "2", "test3", "4", ""}},
		{"test1.2.test3.4.5", '.', []string{"test1", "2", "test3", "4", "5"}},
		{"тестПА1ПА2Пк3ПА4ПА5", 'П', []string{"тест", "А1", "А2", "к3", "А4", "А5"}},
	}
	for _, tt := range tests {
		t.Run(tt.s+" -> "+string(tt.sep), func(t *testing.T) {
			got := SplitRune(tt.s, tt.sep, buf)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitN() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

// degraded, reallocated buffer not reused, use strings.Split for this case
func Benchmark_SplitRealloc(b *testing.B) {
	buf := make([]string, 1)
	s := "test1.2.test3.4.5"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Split(s, ".", buf)
	}
}

// reuse reallocated buffer
func Benchmark_SplitReuse(b *testing.B) {
	buf := make([]string, 1)
	s := "test1.2.test3.4.5"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf = Split(s, ".", buf)
	}
}

func Benchmark_SplitWReuse(b *testing.B) {
	buf := make([]string, 1)
	s := "тестПА1ПА2Пк3ПА4ПА5"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf = Split(s, "ПА", buf)
	}
}

func Benchmark_Split_Check(b *testing.B) {
	buf := make([]string, 1)
	s := "test1.2.test3.4.5"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		count := strings.Count(s, ".")
		if len(buf) < count {
			buf = make([]string, count+1)
		}
		_ = Split(s, ".", buf)
	}
}

func Benchmark_Strings_Split(b *testing.B) {
	s := "test1.2.test3.4.5"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ss := strings.Split(s, ".")
		_ = ss
	}
}

// reuse reallocated buffer
func Benchmark_SplitByteReuse(b *testing.B) {
	buf := make([]string, 1)
	s := "test1.2.test3.4.5"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf = SplitByte(s, '.', buf)
	}
}

// reuse reallocated buffer
func Benchmark_SplitRuneReuse(b *testing.B) {
	buf := make([]string, 1)
	s := "test1.2.test3.4.5"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf = SplitRune(s, '.', buf)
	}
}

// reuse reallocated buffer, unicode rune
func Benchmark_SplitRuneWReuse(b *testing.B) {
	buf := make([]string, 1)
	s := "тестПА1ПА2Пк3ПА4ПА5"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf = SplitRune(s, 'П', buf)
	}
}
