package launch

import (
	"astrolaunch/src/conf"
	"time"

	"github.com/triole/logseal"
)

func (la Launch) Run() (programExitCode int) {
	for _, op := range la.Conf.Content.Operations {
		var err error
		now := time.Now().UTC()
		var diff time.Duration
		fits := false
		op.AtTime, err = la.Calc.GetTime(op.At)
		if err == nil {
			preRange, err := la.str2dur(op.Range.Pre)
			if err != nil {
				la.Lg.Warn(err)
			}
			postRange, err := la.str2dur(op.Range.Post)
			if err != nil {
				la.Lg.Warn(err)
			}
			diff, fits = la.calcRangeDiff(op.AtTime, now, preRange, postRange)
			if err != nil {
				la.Lg.Warn(err)
			}
		}
		if fits {
			la.Lg.Info("exec operation", la.printOpStatus(op, diff, fits))
			_, exitcode, _ := la.execute(op.Exec)
			programExitCode += exitcode
		} else {
			la.Lg.Info("skip operation", la.printOpStatus(op, diff, fits))
		}
	}
	return
}

func (la Launch) printOpStatus(op conf.Operation, diff time.Duration, fits bool) (f logseal.F) {
	f = logseal.F{
		"name": op.Name, "at_str": op.At, "at_time": op.AtTime,
		"exec": op.Exec, "range": op.Range,
		"diff": diff, "fits": fits,
	}
	return
}
