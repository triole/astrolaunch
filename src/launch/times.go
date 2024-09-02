package launch

import (
	"time"

	str2duration "github.com/xhit/go-str2duration/v2"
)

func rangeFits(t1, t2 time.Time, rg time.Duration) (b bool) {
	b = false
	diff := t2.Sub(t1)
	if diff >= 0 {
		b = diff.Seconds() <= rg.Seconds()
	}
	return
}

func str2dur(s string) (dur time.Duration, err error) {
	dur, err = str2duration.ParseDuration(s)
	return
}

func printTime(t time.Time) (s string) {
	return t.Format("2022-03-23 15:04:05 MST")
}
