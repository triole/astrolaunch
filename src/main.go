package main

import (
	"astrolaunch/src/conf"
	"fmt"

	"github.com/triole/logseal"
)

func main() {
	parseArgs()
	lg := logseal.Init(CLI.LogLevel, CLI.LogFile, CLI.LogNoColors, CLI.LogJSON)
	conf := conf.Init(CLI.Conf, lg)
	lg.Info(
		"run "+appName, logseal.F{
			"config": CLI.Conf, "log_level": CLI.LogLevel,
		},
	)
	lg.Debug("full configuration layout", logseal.F{"config": fmt.Sprintf("%+v", conf)})
	// if CLI.ValidateConf {
	// 	pprint(conf)
	// 	os.Exit(0)
	// }
}
