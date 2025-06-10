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
	lg := logseal.Init(cli.LogLevel, cli.LogFile, cli.LogNoColors, cli.LogJSON)
	cnf := conf.Init(cli.Conf, cli.DryRun, lg)
	cnf.ReadConf()
	clc := calc.Init(
		cnf.Now.UTC, cnf.Content.Location.Lat, cnf.Content.Location.Lon,
	)

	switch cli.Action {
	case "calc":
		var add int
		var res []calc.Calc
		if cli.Calc.Date != "" {
			tim, err := time.Parse("20060102", cli.Calc.Date)
			lg.IfErrFatal(
				"can not parse date string, use format YYYYMMDD",
				logseal.F{"error": err, "string": cli.Calc.Date},
			)
			cnf.SetNow(tim)
		}
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
		os.Exit(0)
	case "exec":
		var op conf.Operation
		op.Exec = [][]string{cli.Exec.Cmd}
		op.At = cli.Exec.At
		op.Range.Pre = cnf.Content.DefaultRange.Pre
		op.Range.Post = cnf.Content.DefaultRange.Post
		if cli.Exec.Pre != "" {
			op.Range.Pre = cli.Exec.Pre
		}
		if cli.Exec.Post != "" {
			op.Range.Post = cli.Exec.Post
		}
		cnf.Content.Operations = append(cnf.Content.Operations, op)
	case "list":
		for k := range clc.Sun {
			fmt.Printf("sun.%s\n", k)
		}
		for k := range clc.Moon {
			fmt.Printf("moon.%s\n", k)
		}
	case "ops":
		cnf.OpsFilter = cli.Ops.Filter
		cnf.ReadOps()
		lg.Info(
			"run "+appName, logseal.F{
				"config": cnf.FileName, "log_level": cli.LogLevel,
			},
		)
		lg.Debug("full config", logseal.F{"config": fmt.Sprintf("%+v", cnf)})
		lg.Debug("astro calculations", logseal.F{"calc": fmt.Sprintf("%+v", clc)})
	}

	var programExitCode int
	la := launch.Init(cnf, clc, lg)
	if cli.Exec.Wait {
		programExitCode = la.WaitAndRun()
	} else {
		programExitCode = la.Run()
	}
	lg.Info("done", logseal.F{"occured_errors": programExitCode})
	os.Exit(programExitCode)
}

func pprint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "  ")
	fmt.Println(string(s))
}
