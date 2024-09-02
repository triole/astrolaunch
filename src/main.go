package main

import (
	"astrolaunch/src/calc"
	"astrolaunch/src/conf"
	"astrolaunch/src/launch"
	"fmt"
	"time"

	"github.com/triole/logseal"
)

func main() {
	parseArgs()
	lg := logseal.Init(CLI.LogLevel, CLI.LogFile, CLI.LogNoColors, CLI.LogJSON)
	conf := conf.Init(CLI.Conf, lg)
	calc := calc.Init(
		time.Now(), conf.Content.Location.Lat, conf.Content.Location.Lon,
	)
	lg.Debug(
		"run "+appName, logseal.F{
			"config": CLI.Conf, "log_level": CLI.LogLevel,
		},
	)
	lg.Debug("full configuration layout", logseal.F{"config": fmt.Sprintf("%+v", conf)})

	la := launch.Init(conf, calc, lg)
	la.Run()
}
