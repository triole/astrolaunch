package launch

import (
	"astrolaunch/src/conf"
	"time"

	"github.com/triole/logseal"
)

func (la Launch) runTicker(op conf.Operation) {
	var err error
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		la.Conf.SetNow(time.Now())
		op.AtTime, err = la.Calc.GetTime(op.At)
		if err != nil {
			la.Lg.Warn(
				"operation skipped to due value fetch error",
				logseal.F{"error": err, "operation_name": op.Name},
			)
		} else {
			preRange, _ := la.str2dur(op.Range.Pre, true)
			postRange, _ := la.str2dur(op.Range.Post, true)
			diff, fits := la.calcRangeDiff(
				op.AtTime, la.Conf.Now.UTC, preRange, postRange,
			)
			la.Lg.Debug(
				"diff too large, continue to tick",
				logseal.F{"time": la.Conf.Now.Local, "diff": diff},
			)
			if fits {
				ticker.Stop()
				break
			}
		}
	}
}
