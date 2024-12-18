package main

import (
	"astrolaunch/src/calc"
	"astrolaunch/src/conf"
	"astrolaunch/src/launch"
	"encoding/json"
	"fmt"
	"os"

	"github.com/triole/logseal"
)

func main() {
	parseArgs()
	lg := logseal.Init(cli.LogLevel, cli.LogFile, cli.LogNoColors, cli.LogJSON)
	// if cli.Date != "" {
	// 	tim, err := time.Parse("20060102", cli.Date)
	// 	lg.IfErrFatal("can not parse date string, use format YYYYMMDD", logseal.F{"error": err, "string": cli.Date})
	// 	now = tim
	// }
	cnf := conf.Init(cli.Conf, cli.DryRun, lg)
	clc := calc.Init(
		cnf.Now.UTC, cnf.Content.Location.Lat, cnf.Content.Location.Lon,
	)

	switch cli.Action {
	case "calc":
		var add int
		var res []calc.Calc
		for i := 0; i <= cli.Calc.Range; i++ {
			add = 1
			if i == 0 {
				add = 0
			}
			cnf.SetNow(cnf.Now.Local.AddDate(0, 0, add))
			clc = calc.Init(
				cnf.Now.UTC, cnf.Content.Location.Lat, cnf.Content.Location.Lon,
			)
			res = append(res, clc)
		}
		if cli.Calc.Range == 0 {
			pprint(res[0])
		} else {
			pprint(res)
		}
	case "exec":
		fmt.Println("Tomorrow.")
	case "ops":
		cnf.OpsFilter = cli.Ops.Filter
		cnf.ReadConf()
		cnf.ReadOps()
		lg.Info(
			"run "+appName, logseal.F{
				"config": cnf.FileName, "log_level": cli.LogLevel,
			},
		)
		lg.Debug("full config", logseal.F{"config": fmt.Sprintf("%+v", cnf)})
		lg.Debug("astro calculations", logseal.F{"calc": fmt.Sprintf("%+v", clc)})

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
