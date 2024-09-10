package main

import (
	"astrolaunch/src/calc"
	"astrolaunch/src/conf"
	"astrolaunch/src/launch"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/triole/logseal"
)

func main() {
	parseArgs()
	now := time.Now()
	lg := logseal.Init(CLI.LogLevel, CLI.LogFile, CLI.LogNoColors, CLI.LogJSON)
	conf := conf.Init(now, CLI.Conf, lg)
	conf.DryRun = CLI.DryRun
	calc := calc.Init(
		conf.Now.UTC, conf.Content.Location.Lat, conf.Content.Location.Lon,
	)

	if CLI.Astro {
		pprint(calc)
	} else {
		lg.Info(
			"run "+appName, logseal.F{
				"config": CLI.Conf, "log_level": CLI.LogLevel,
			},
		)
		fmt.Printf("%+v\n", conf)
		lg.Debug("full config", logseal.F{"config": fmt.Sprintf("%+v", conf)})
		lg.Debug("astro calculations", logseal.F{"config": fmt.Sprintf("%+v", calc)})

		la := launch.Init(conf, calc, lg)
		programExitCode := la.Run()
		lg.Info("done", logseal.F{"occured_errors": programExitCode})
		os.Exit(programExitCode)
	}
}

func pprint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "  ")
	fmt.Println(string(s))
}
