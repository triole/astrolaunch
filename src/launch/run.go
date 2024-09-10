package launch

import (
	"astrolaunch/src/conf"
	"time"

	"github.com/triole/logseal"
)

func (la Launch) Run() (programExitCode int) {
	for _, op := range la.Conf.Content.Operations {
		now := time.Now().UTC()
		var diff time.Duration
		fits := false
		at, err := la.Calc.GetTime(op.At)
		if err == nil {
			rangeDuration, err := str2dur(op.Range)
			diff, fits = calcRangeDiff(now, at, rangeDuration)
			if err == nil {
			} else {
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
		"name": op.Name, "at": op.At, "exec": op.Exec, "range": op.Range,
		"diff": diff, "fits": fits,
	}
	return
}
