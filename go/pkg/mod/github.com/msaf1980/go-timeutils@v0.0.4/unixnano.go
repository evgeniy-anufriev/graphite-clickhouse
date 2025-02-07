package timeutils

import (
	"strconv"
	"time"
)

// UnixNano create unix timestamp with nanoseconds timestamp (from epoch)
func UnixNano(ts int64) time.Time {
	s := ts / 1e9
	ns := ts % 1e9
	return time.Unix(s, ns)
}

// ParseUnixNano parse string and create unix timestamp with nanoseconds timestamp (from epoch)
func ParseUnixNano(s string) (t time.Time, err error) {
	var ts int64
	ts, err = strconv.ParseInt(s, 10, 64)
	if err != nil {
		return
	}
	t = UnixNano(ts)
	return
}
