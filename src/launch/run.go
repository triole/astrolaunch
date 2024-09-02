package launch

import (
	"time"
)

func (la Launch) Run() {
	for _, op := range la.Conf.Content.Operations {
		now := time.Now().UTC()
		fits := false
		at, err := la.Calc.GetTime(op.At)
		if err == nil {
			rangeDuration, err := str2dur(op.Range)
			fits = rangeFits(now, at, rangeDuration)
			if err == nil {
			} else {
				la.Lg.Warn(err)
			}
		}
		if fits {
			la.execute(op.Exec)
		}
	}
}
