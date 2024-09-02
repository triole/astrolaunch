package launch

import (
	"astrolaunch/src/calc"
	"astrolaunch/src/conf"

	"github.com/triole/logseal"
)

type Launch struct {
	Conf conf.Conf
	Calc calc.Calc
	Lg   logseal.Logseal
}

func Init(conf conf.Conf, calc calc.Calc, lg logseal.Logseal) (la Launch) {
	return Launch{
		Conf: conf,
		Calc: calc,
		Lg:   lg,
	}
}
