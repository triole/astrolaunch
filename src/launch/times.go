package launch

import (
	"time"

	"github.com/triole/logseal"
	str2duration "github.com/xhit/go-str2duration/v2"
)

func (la Launch) calcRangeDiff(t1, t2 time.Time, preRange, postRange time.Duration) (diff time.Duration, b bool) {
	diff = t2.Sub(t1)
	diffSec := diff.Seconds()
	b = diffSec >= -(preRange.Seconds()) && diffSec <= postRange.Seconds()
	la.Lg.Debug(
		"calc range diff",
		logseal.F{"t1": t1, "t2": t2, "rpre": preRange, "rpost": postRange, "diff_in_s": diffSec},
	)
	return
}

func (la Launch) str2dur(s string) (dur time.Duration, err error) {
	dur, err = str2duration.ParseDuration(s)
	return
}

func (la Launch) printTime(t time.Time) (s string) {
	return t.Format("2006-01-02 15:04:05 MST")
}
