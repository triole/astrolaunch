package launch

import (
	"time"

	"github.com/triole/logseal"
	str2duration "github.com/xhit/go-str2duration/v2"
)

func (la Launch) calcRangeDiff(t1, t2 time.Time, preRange, postRange time.Duration) (diff time.Duration, b bool) {
	diff = t2.Sub(t1)
	diffSec := diff.Seconds()
	preSec := -(preRange.Seconds())
	postSec := postRange.Seconds()
	b = preSec <= diffSec && diffSec <= postSec
	la.Lg.Debug(
		"calc range diff",
		logseal.F{"t1": t1, "t2": t2, "rpre": preSec, "rpost": postSec, "diff": diffSec},
	)
	return
}

func (la Launch) str2dur(s string, verbose bool) (dur time.Duration, err error) {
	dur, err = str2duration.ParseDuration(s)
	if verbose {
		if err != nil {
			la.Lg.Warn(err)
		}
	}
	return
}

func (la Launch) printTime(t time.Time) (s string) {
	return t.Format("2006-01-02 15:04:05 MST")
}
