package launch

import (
	"fmt"
	"time"

	str2duration "github.com/xhit/go-str2duration/v2"
)

func (la Launch) Run() {
	for _, op := range la.Conf.Content.Operations {
		now := time.Now()
		fit := false
		at, err := la.Calc.GetTime(op.At)
		if err == nil {
			rangeDuration, err := str2dur(op.Range)
			fit = timesFit(now, at, rangeDuration)
			if err == nil {
			} else {
				la.Lg.Warn(err)
			}
		}
		fmt.Printf("%+v\n", fit)
	}
}

func timesFit(t1, t2 time.Time, rg time.Duration) (b bool) {
	diff := t2.Sub(t1)
	if diff >= 0 {
		b = diff.Seconds() <= rg.Seconds()
	}
	fmt.Printf(
		"%+v == %+v, %v, %v\n", printTime(t1), printTime(t2), diff, b,
	)
	return
}

func str2dur(s string) (dur time.Duration, err error) {
	dur, err = str2duration.ParseDuration(s)
	return
}

func printTime(t time.Time) (s string) {
	return t.Format("2022-03-23 15:04:05 MST")
}
