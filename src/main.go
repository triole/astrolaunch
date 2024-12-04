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
	lg := logseal.Init(CLI.LogLevel, CLI.LogFile, CLI.LogNoColors, CLI.LogJSON)
	now := time.Now()
	if CLI.Date != "" {
		tim, err := time.Parse("20060102", CLI.Date)
		lg.IfErrFatal("can not parse date string, use format YYYYMMDD", logseal.F{"error": err, "string": CLI.Date})
		now = tim
	}
	cnf := conf.Init(now, CLI.Conf, lg)
	cnf.DryRun = CLI.DryRun
	clc := calc.Init(
		cnf.Now.UTC, cnf.Content.Location.Lat, cnf.Content.Location.Lon,
	)

	if CLI.Astro {
		var add int
		for i := 0; i <= CLI.Range; i++ {
			add = 1
			if i == 0 {
				add = 0
			}
			now = now.AddDate(0, 0, add)
			cnf := conf.Init(now, CLI.Conf, lg)
			clc := calc.Init(
				cnf.Now.UTC, cnf.Content.Location.Lat, cnf.Content.Location.Lon,
			)
			pprint(clc)
		}
	} else {
		lg.Info(
			"run "+appName, logseal.F{
				"config": CLI.Conf, "log_level": CLI.LogLevel,
			},
		)
		lg.Debug("full config", logseal.F{"config": fmt.Sprintf("%+v", cnf)})
		lg.Debug("astro calculations", logseal.F{"config": fmt.Sprintf("%+v", clc)})

		la := launch.Init(cnf, clc, lg)
		programExitCode := la.Run()
		lg.Info("done", logseal.F{"occured_errors": programExitCode})
		os.Exit(programExitCode)
	}
}

func pprint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "  ")
	fmt.Println(string(s))
}
