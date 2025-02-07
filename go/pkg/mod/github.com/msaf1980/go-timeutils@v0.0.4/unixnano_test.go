package timeutils

import (
	"reflect"
	"testing"
	"time"
)

type TimeCmp struct {
	year  int
	month time.Month
	day   int
	hour  int
	min   int
	sec   int
	nsec  int
	loc   string
}

func NewTimeCmp(t time.Time) TimeCmp {
	return TimeCmp{
		year:  t.Year(),
		month: t.Month(),
		day:   t.Day(),
		hour:  t.Hour(),
		min:   t.Minute(),
		sec:   t.Second(),
		nsec:  t.Nanosecond(),
		loc:   t.Location().String(),
	}
}

func TestUnixNanoUTC(t *testing.T) {
	nanoutctests := []struct {
		ts   int64
		want time.Time
	}{
		{0, time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)},
		{1671344766000000000, time.Date(2022, time.December, 18, 6, 26, 6, 0, time.UTC)},
		{1671344766000000010, time.Date(2022, time.December, 18, 6, 26, 6, 10, time.UTC)},
		{1671344766999999999, time.Date(2022, time.December, 18, 6, 26, 6, 999999999, time.UTC)},
	}
	for _, tt := range nanoutctests {
		got := UnixNano(tt.ts).UTC()

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("UnixNano(%d).UTC()\ngot\n%+v\nwant\n%+v", tt.ts, NewTimeCmp(got), NewTimeCmp(tt.want))
		}
	}
}
