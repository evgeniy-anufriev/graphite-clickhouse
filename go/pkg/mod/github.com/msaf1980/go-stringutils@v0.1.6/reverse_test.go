package stringutils

import "testing"

func TestReverse(t *testing.T) {
	for _, c := range []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"Hello, мир", "рим ,olleH"},
		{"", ""},
	} {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestReverseSegments(t *testing.T) {
	for _, c := range []struct {
		in, want, delim string
	}{
		{"Hello, world", "world, Hello", ", "},
		{"Hello\t世界", "世界\tHello", "\t"},
		{"Hello мир", "мир Hello", " "},
		{"Hello мир", "Hello мир", ""},
		{"", "", "\t"},
	} {
		got := ReverseSegments(c.in, c.delim)
		if got != c.want {
			t.Errorf("ReverseSegments(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
