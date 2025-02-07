package duration

import (
	"testing"
	"time"
)

func TestString(t *testing.T) {
	tests := []struct {
		d    time.Duration
		want string
	}{
		{
			d:    time.Hour,
			want: "1h",
		},
		{
			d:    12*time.Hour + time.Second,
			want: "12h1s",
		},
		{
			d:    12*time.Hour + 31*time.Minute + 7*time.Second,
			want: "12h31m7s",
		},
		{
			d:    12*time.Hour + 31*time.Minute + 7*time.Second + 100*time.Millisecond + time.Nanosecond,
			want: "12h31m7.100000001s",
		},
		{
			d:    12*time.Hour + 31*time.Minute + 7*time.Second + 10*time.Millisecond + time.Nanosecond,
			want: "12h31m7.010000001s",
		},
		{
			d:    12*time.Hour + 31*time.Minute + 7*time.Second + time.Microsecond + 10*time.Nanosecond,
			want: "12h31m7.00000101s",
		},
		{
			d:    12*time.Hour + 31*time.Minute + 10*time.Nanosecond,
			want: "12h31m0.00000001s",
		},
		{
			d:    10 * time.Nanosecond,
			want: "10ns",
		},
		{
			d:    2*time.Microsecond + 10*time.Nanosecond,
			want: "2.01µs",
		},
		{
			d:    4*time.Millisecond + 2*time.Microsecond + 10*time.Nanosecond,
			want: "4.00201ms",
		},
	}
	for _, tt := range tests {
		t.Run(tt.d.String(), func(t *testing.T) {
			if got := String(tt.d); got != tt.want {
				t.Errorf("String() = %v, want %v (%s)", got, tt.want, tt.d.String())
			}
		})
	}
}
