package launch

import (
	"astrolaunch/src/conf"
	"time"

	"github.com/triole/logseal"
)

func (la Launch) WaitAndRun() (programExitCode int) {
	la.runTicker(la.Conf.Content.Operations[0])
	la.Run()
	return
}

func (la Launch) Run() (programExitCode int) {
	for _, op := range la.Conf.Content.Operations {
		var err error
		var diff time.Duration
		fits := false

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
			diff, fits = la.calcRangeDiff(
				op.AtTime, la.Conf.Now.UTC, preRange, postRange,
			)
			if fits {
				la.Lg.Info("exec operation", la.printOpStatus(op, diff, fits))
				_, exitcode, _ := la.execute(op.Exec)
				programExitCode += exitcode
			} else {
				la.Lg.Info("skip operation", la.printOpStatus(op, diff, fits))
			}
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
