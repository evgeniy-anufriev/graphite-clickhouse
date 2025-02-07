package stringutils

import (
	"math"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestBuilder_Grow(t *testing.T) {
	tests := []struct {
		capacity int
	}{
		{1},
		{2},
		{100},
	}
	var sb Builder
	for _, tt := range tests {
		t.Run("Grow", func(t *testing.T) {
			sb.Grow(tt.capacity)
			if sb.Cap() != tt.capacity {
				t.Errorf("Grow(%d) capacity = %d', want %d", tt.capacity, sb.Cap(), tt.capacity)
			}
			if sb.Len() != 0 {
				t.Errorf("Len() = %d', want 0", sb.Len())
			}
		})
	}
}

func TestBuilder_Len(t *testing.T) {
	var sb Builder
	b := make([]byte, 0)

	if sb.Cap() != 0 {
		t.Errorf("Cap() = %d', want 0", sb.Cap())
	}
	if sb.Len() != 0 {
		t.Errorf("Length() = %d', want 0", sb.Len())
	}
	if sb.Bytes() != nil && !reflect.DeepEqual(sb.Bytes(), b) {
		t.Errorf("Bytes() = '%+v', want nil", sb.Bytes())
	}

	sb.Grow(10)
	if sb.Cap() != 10 {
		t.Errorf("Cap() = %d', want 10", sb.Cap())
	}
	if sb.Len() != 0 {
		t.Errorf("Length() = %d', want 0", sb.Len())
	}
	if sb.Bytes() != nil && !reflect.DeepEqual(sb.Bytes(), b) {
		t.Errorf("Bytes() = '%+v', want nil", sb.Bytes())
	}

	sb.Reset()
	if sb.Cap() != 10 {
		t.Errorf("Cap() = %d', want 10", sb.Cap())
	}
	if sb.Len() != 0 {
		t.Errorf("Length() = %d', want 0", sb.Len())
	}
	if sb.Bytes() != nil && !reflect.DeepEqual(sb.Bytes(), b) {
		t.Errorf("Bytes() = '%+v', want nil", sb.Bytes())
	}

	sb.Release()
	if sb.Cap() != 0 {
		t.Errorf("Cap() = %d', want 0", sb.Cap())
	}
	if sb.Len() != 0 {
		t.Errorf("Length() = %d', want 0", sb.Len())
	}
	if sb.Bytes() != nil && !reflect.DeepEqual(sb.Bytes(), b) {
		t.Errorf("Bytes() = '%+v', want nil", sb.Bytes())
	}
}

func TestBuilder_WriteString(t *testing.T) {
	tests := []struct {
		s       string
		reset   bool
		release bool
		want    string
	}{
		{"one.", false, false, "one."},
		{"twotwo.", false, false, "one.twotwo."},
		{"", true, false, ""},
		{"three.", true, false, "three."},
		{"", false, true, ""},
		{"four.", false, false, "four."},
	}
	var sb Builder
	for id, tt := range tests {
		t.Run("Test #"+strconv.Itoa(id), func(t *testing.T) {
			if tt.reset {
				sb.Reset()
				if sb.Len() != 0 {
					t.Errorf("Length() = %d', want 0", sb.Len())
				}
			}
			if tt.release {
				sb.Release()
				if sb.Cap() != 0 {
					t.Errorf("Cap() = %d', want 0", sb.Cap())
				}
				if sb.Len() != 0 {
					t.Errorf("Length() = %d', want 0", sb.Len())
				}
			}
			sb.WriteString(tt.s)
			if sb.String() != tt.want {
				t.Errorf("String() = '%s', want '%s'", sb.String(), tt.want)
			}
		})
	}
}

func TestBuilder_Write(t *testing.T) {
	tests := []struct {
		s       string
		reset   bool
		release bool
		want    string
	}{
		{"one.", false, false, "one."},
		{"twotwo.", false, false, "one.twotwo."},
		{"", true, false, ""},
		{"three.", false, false, "three."},
		{"", false, true, ""},
		{"four.", false, false, "four."},
	}
	var sb Builder
	for id, tt := range tests {
		t.Run("Test #"+strconv.Itoa(id), func(t *testing.T) {
			if tt.reset {
				sb.Reset()
				if sb.Len() != 0 {
					t.Errorf("Length() = %d', want 0", sb.Len())
				}
			}
			if tt.release {
				sb.Release()
				if sb.Cap() != 0 {
					t.Errorf("Cap() = %d', want 0", sb.Cap())
				}
				if sb.Len() != 0 {
					t.Errorf("Length() = %d', want 0", sb.Len())
				}
			}
			sb.Write([]byte(tt.s))
			if sb.String() != tt.want {
				t.Errorf("String() = '%s', want '%s'", sb.String(), tt.want)
			}
		})
	}
}

func TestBuilder_WriteByte(t *testing.T) {
	tests := []struct {
		b    byte
		want string
	}{
		{byte('c'), "c"},
		{byte('d'), "cd"},
	}
	var sb Builder
	for id, tt := range tests {
		t.Run("Test #"+strconv.Itoa(id), func(t *testing.T) {
			sb.WriteByte(tt.b)
			if sb.String() != tt.want {
				t.Errorf("String() = '%s', want '%s'", sb.String(), tt.want)
			}
		})
	}
}

func TestBuilder_WriteRune(t *testing.T) {
	const s0 = "hello 世界"

	tests := []struct {
		name string
		fn   func(b *Builder)
		want string
	}{
		{
			"Write",
			func(sb *Builder) { sb.Write([]byte(s0)) },
			s0,
		},
		{
			"WriteRune",
			func(sb *Builder) { sb.WriteRune('a') },
			"a",
		},
		{
			"WriteRuneWide",
			func(sb *Builder) { sb.WriteRune('世') },
			"世",
		},
		{
			"WriteString",
			func(sb *Builder) { sb.WriteString(s0) },
			s0,
		},
	}

	for id, tt := range tests {
		t.Run("Test #"+strconv.Itoa(id), func(t *testing.T) {
			var sb Builder
			tt.fn(&sb)
			if sb.String() != tt.want {
				t.Errorf("String() = '%s', want '%s'", sb.String(), tt.want)
			}
		})
	}
}

func TestBuilder_WriteInt(t *testing.T) {
	tests := []struct {
		n           int64
		appendSpace bool
		want        string
	}{
		{n: 10, want: "10"},
		{n: -123, want: "10-123"},
		{n: -428009, want: "10-123-428009"},
		{n: 328007, want: "10-123-428009328007"},
		{n: math.MaxInt64, appendSpace: true, want: "10-123-428009328007 " + strconv.FormatUint(math.MaxInt64, 10)},
	}
	var sb Builder
	sb.Grow(128)
	for id, tt := range tests {
		t.Run("Test #"+strconv.Itoa(id), func(t *testing.T) {
			if tt.appendSpace {
				sb.WriteRune(' ')
			}
			sb.WriteInt(tt.n, 10)
			if sb.String() != tt.want {
				t.Errorf("String()\n'%s'\nwant\n'%s'", sb.String(), tt.want)
			}
		})
	}
}

func TestBuilder_WriteUint(t *testing.T) {
	tests := []struct {
		n           uint64
		appendSpace bool
		want        string
	}{
		{n: 10, want: "10"},
		{n: 123, want: "10123"},
		{n: 428009, want: "10123428009"},
		{n: 328007, want: "10123428009328007"},
		{n: math.MaxUint64, appendSpace: true, want: "10123428009328007 " + strconv.FormatUint(math.MaxUint64, 10)},
	}
	var sb Builder
	sb.Grow(128)
	for id, tt := range tests {
		t.Run("Test #"+strconv.Itoa(id), func(t *testing.T) {
			if tt.appendSpace {
				sb.WriteString(" ")
			}
			sb.WriteUint(tt.n, 10)
			if sb.String() != tt.want {
				t.Errorf("String()\n'%s'\nwant\n'%s'", sb.String(), tt.want)
			}
		})
	}
}

func TestBuilder_WriteFloat(t *testing.T) {
	tests := []struct {
		n           float64
		appendSpace bool
		want        string
	}{
		{n: 10.0, want: "10"},
		{n: -123.0, want: "10-123"},
		{n: -428009.45678444, want: "10-123-428009.45678444"},
		{n: math.MaxFloat64, appendSpace: true, want: "10-123-428009.45678444 " + strconv.FormatFloat(math.MaxFloat64, 'f', -1, 64)},
	}
	var sb Builder
	sb.Grow(128)
	for id, tt := range tests {
		t.Run("Test #"+strconv.Itoa(id), func(t *testing.T) {
			if tt.appendSpace {
				sb.WriteRune(' ')
			}
			sb.WriteFloat(tt.n, 'f', -1, 64)
			if sb.String() != tt.want {
				t.Errorf("String()\n'%s'\nwant\n'%s'", sb.String(), tt.want)
			}
		})
	}
}

func TestBuilder(t *testing.T) {
	const s0 = "hello 世界"

	tests := []struct {
		name    string
		reset   bool
		release bool
		fn      func(b *Builder)
		want    string
	}{
		{
			name: "Write",
			fn:   func(sb *Builder) { sb.Write([]byte(s0)) },
			want: s0,
		},
		{
			name: "WriteRune",
			fn:   func(sb *Builder) { sb.WriteRune('a') },
			want: s0 + "a",
		},
		{
			name: "WriteRuneWide",
			fn:   func(sb *Builder) { sb.WriteRune('世') },
			want: s0 + "a世",
		},
		{
			name: "WriteInt",
			fn:   func(sb *Builder) { sb.WriteInt(math.MaxInt64, 10) },
			want: s0 + "a世" + strconv.FormatUint(math.MaxInt64, 10),
		},
		{
			name: "WriteString #1",
			fn:   func(sb *Builder) { sb.WriteString(s0) },
			want: s0 + "a世" + strconv.FormatUint(math.MaxInt64, 10) + s0,
		},
		{
			name:  "WriteString #2",
			reset: true,
			fn:    func(sb *Builder) { sb.WriteString(s0) },
			want:  s0,
		},
		{
			name:    "WriteQuote",
			release: true,
			fn:      func(sb *Builder) { sb.WriteQuote(s0) },
			want:    "\"" + s0 + "\"",
		},
		{
			name: "WriteQuoteRune",
			fn:   func(sb *Builder) { sb.WriteQuoteRune(' ') },
			want: "\"" + s0 + "\"" + "' '",
		},
		{
			name: "WriteQuoteRuneToASCII space",
			fn:   func(sb *Builder) { sb.WriteQuoteRuneToASCII(' ') },
			want: "\"" + s0 + "\"" + "' '' '",
		},
		{
			name:    "WriteQuoteRuneToASCII ascii",
			release: true,
			fn:      func(sb *Builder) { sb.WriteQuoteRuneToASCII('a') },
			want:    "'a'",
		},
		{
			name:  "WriteQuoteRuneToASCII wide",
			reset: true,
			fn:    func(sb *Builder) { sb.WriteQuoteRuneToASCII('世') },
			want:  `'\u4e16'`,
		},
		{
			name:  "WriteQuoteRuneToGraphic",
			reset: true,
			fn:    func(sb *Builder) { sb.WriteQuoteRuneToGraphic('\u4e16') },
			want:  "'世'",
		},
		{
			name: "WriteQuoteToGraphic",
			fn:   func(sb *Builder) { sb.WriteQuoteToGraphic("\u4e16") },
			want: "'世'\"世\"",
		},
	}

	var sb Builder
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.reset {
				sb.Reset()
				if sb.Len() != 0 {
					t.Errorf("Length() = %d', want 0", sb.Len())
				}
			}
			if tt.release {
				sb.Release()
				if sb.Cap() != 0 {
					t.Errorf("Cap() = %d', want 0", sb.Cap())
				}
				if sb.Len() != 0 {
					t.Errorf("Length() = %d', want 0", sb.Len())
				}
			}
			tt.fn(&sb)
			if sb.String() != tt.want {
				t.Fatalf("String()\n'%s'\nwant\n'%s'", sb.String(), tt.want)
			}
		})
	}
}

func Benchmark_String_RawCopy(b *testing.B) {
	buf := make([]byte, 1000000)
	pos := 0
	s := "asdfghjklqwertyuiopzxcvbnm1234567890"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if pos+len(s) > len(buf) {
			pos = 0
		}
		copy(buf[pos:], s)
		pos += len(s)
	}
}

func BenchmarkStd_StringsBuilder_WriteString(b *testing.B) {
	var sb strings.Builder
	sb.Grow(1000000)
	sb.Reset()
	s := "asdfghjklqwertyuiopzxcvbnm1234567890"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if sb.Len()+len(s) > sb.Cap() {
			sb.Reset()
		}
		sb.WriteString(s)
	}
}

func BenchmarkThis_Builder_WriteString(b *testing.B) {
	var sb Builder
	sb.Grow(1000000)
	sb.Reset()
	s := "asdfghjklqwertyuiopzxcvbnm1234567890"

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if sb.Len()+len(s) > sb.Cap() {
			sb.Reset()
		}
		sb.WriteString(s)
	}
}

func BenchmarkStd_StringsBuilder_Write(b *testing.B) {
	var sb strings.Builder
	sb.Grow(1000000)
	sb.Reset()
	bytes := []byte("asdfghjklqwertyuiopzxcvbnm1234567890")

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if sb.Len()+len(bytes) > sb.Cap() {
			sb.Reset()
		}
		sb.Write(bytes)
	}
}

func BenchmarkThis_Builder_Write(b *testing.B) {
	var sb Builder
	sb.Grow(1000000)
	sb.Reset()
	bytes := []byte("asdfghjklqwertyuiopzxcvbnm1234567890")

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if sb.Len()+len(bytes) > sb.Cap() {
			sb.Reset()
		}
		sb.Write(bytes)
	}
}

func BenchmarkThis_Builder_WriteBytes(b *testing.B) {
	var sb Builder
	sb.Grow(1000000)
	sb.Reset()
	bytes := []byte("asdfghjklqwertyuiopzxcvbnm1234567890")

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if sb.Len()+len(bytes) > sb.Cap() {
			sb.Reset()
		}
		sb.WriteBytes(bytes)
	}
}

func BenchmarkThis_StringsBuilderString(b *testing.B) {
	var sb Builder
	sb.Grow(1000000)
	sb.Reset()
	s := "asdfghjklqwertyuiopzxcvbnm1234567890"
	sb.WriteString(s)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = sb.String()
	}
}

func BenchmarkThis_Builder_WriteIntSmall(b *testing.B) {
	var sb Builder
	sb.Grow(1000000)
	sb.Reset()
	var n int64 = 10

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if sb.Len()+1024 > sb.Cap() {
			sb.Reset()
		}
		sb.WriteInt(n, 10)
	}
}

func BenchmarkStd_strconv_FormatIntSmall(b *testing.B) {
	var sb Builder
	sb.Grow(1000000)
	sb.Reset()
	var n int64 = 10

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if sb.Len()+128 > sb.Cap() {
			sb.Reset()
		}
		s := strconv.FormatInt(n, 10)
		sb.WriteString(s)
	}
}

func BenchmarkThis_Builder_WriteIntLarge(b *testing.B) {
	var sb Builder
	sb.Grow(1000000)
	sb.Reset()
	var n int64 = 102400

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if sb.Len()+128 > sb.Cap() {
			sb.Reset()
		}
		sb.WriteInt(n, 10)
	}
}

func BenchmarkStd_strconv_FormatIntLarge(b *testing.B) {
	var sb Builder
	sb.Grow(1000000)
	sb.Reset()
	var n int64 = 102400

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if sb.Len()+128 > sb.Cap() {
			sb.Reset()
		}
		s := strconv.FormatInt(n, 10)
		sb.WriteString(s)
	}
}

func BenchmarkThis_Builder_WriteFloat(b *testing.B) {
	var sb Builder
	sb.Grow(1000000)
	sb.Reset()
	var n float64 = 102400.12345667

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if sb.Len()+128 > sb.Cap() {
			sb.Reset()
		}
		sb.WriteFloat(n, 'f', -1, 64)
	}
}

func BenchmarkStd_strconv_FormatFloat(b *testing.B) {
	var sb Builder
	sb.Grow(1000000)
	sb.Reset()
	var n float64 = 102400.12345667

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if sb.Len()+128 > sb.Cap() {
			sb.Reset()
		}
		s := strconv.FormatFloat(n, 'f', -1, 64)
		sb.WriteString(s)
	}
}
