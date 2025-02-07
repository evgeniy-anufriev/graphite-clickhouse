package stringutils

import (
	"strings"
	"testing"
)

var ReplaceTests = []struct {
	in       string
	old, new string
	n        int
	out      string
	changed  bool
}{
	{"hello", "l", "L", 0, "hello", false},
	{"hello", "l", "L", -1, "heLLo", true},
	{"hello", "x", "X", -1, "hello", false},
	{"", "x", "X", -1, "", false},
	{"radar", "r", "<r>", -1, "<r>ada<r>", true},
	{"", "", "<>", -1, "<>", true},
	{"banana", "a", "<>", -1, "b<>n<>n<>", true},
	{"banana", "a", "<>", 1, "b<>nana", true},
	{"banana", "a", "<>", 1000, "b<>n<>n<>", true},
	{"banana", "an", "<>", -1, "b<><>a", true},
	{"banana", "ana", "<>", -1, "b<>na", true},
	{"banana", "", "<>", -1, "<>b<>a<>n<>a<>n<>a<>", true},
	{"banana", "", "<>", 10, "<>b<>a<>n<>a<>n<>a<>", true},
	{"banana", "", "<>", 6, "<>b<>a<>n<>a<>n<>a", true},
	{"banana", "", "<>", 5, "<>b<>a<>n<>a<>na", true},
	{"banana", "", "<>", 1, "<>banana", true},
	{"banana", "a", "a", -1, "banana", false},
	{"banana", "a", "a", 1, "banana", false},
	{"☺☻☹", "", "<>", -1, "<>☺<>☻<>☹<>", true},
}

func TestReplace(t *testing.T) {
	for _, tt := range ReplaceTests {
		if s, changed := Replace(tt.in, tt.old, tt.new, tt.n); s != tt.out {
			t.Errorf("Replace(%q, %q, %q, %d).value = %q, want %q", tt.in, tt.old, tt.new, tt.n, s, tt.out)
		} else if changed != tt.changed {
			t.Errorf("Replace(%q, %q, %q).changed = %v, want %v", tt.in, tt.old, tt.new, changed, tt.changed)
		}
		if tt.n == -1 {
			s, changed := ReplaceAll(tt.in, tt.old, tt.new)
			if s != tt.out {
				t.Errorf("ReplaceAll(%q, %q, %q).value = %q, want %q", tt.in, tt.old, tt.new, s, tt.out)
			} else if changed != tt.changed {
				t.Errorf("ReplaceAll(%q, %q, %q).changed = %v, want %v", tt.in, tt.old, tt.new, changed, tt.changed)
			}
		}
	}
}

func Benchmark_strings_ReplaceAll(b *testing.B) {
	s := "test1.2.test3.4.5"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ss := strings.ReplaceAll(s, ".", " ")
		_ = ss
	}
}

func Benchmark_ReplaceAll(b *testing.B) {
	s := "test1.2.test3.4.5"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ss, changed := ReplaceAll(s, ".", " ")
		_ = ss
		_ = changed
	}
}
