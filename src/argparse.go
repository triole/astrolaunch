package main

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"regexp"
	"strings"

	"github.com/alecthomas/kong"
)

var (
	BUILDTAGS      string
	appName        = "astrolaunch"
	appDescription = "launch commands at sun rise, sun dawn or other astro related times"
	appMainversion = "0.1"
)

var cli struct {
	Action      string `kong:"-" enum:"conf,exec,calc" default:"conf"`
	Conf        string `help:"path to config file" short:"c" default:"${configFile}"`
	LogFile     string `help:"log file" default:"/dev/stdout"`
	LogLevel    string `help:"log level" default:"info" enum:"trace,debug,info,error"`
	LogNoColors bool   `help:"disable output colours, print plain text"`
	LogJSON     bool   `help:"enable json log, instead of text one"`
	DryRun      bool   `help:"dry run, just print operations that would run" short:"n"`
	VersionFlag bool   `help:"display version" short:"V"`

	Calc struct {
		Date  string `help:"print astro calculation for a certain date, format: YYYYMMDD, [use with -a]" short:"d"`
		Range int    `help:"range of days, astro calculation for a multiple days, [use with -a]" short:"r"`
	} `cmd:"" help:"list files matching the criteria"`

	Exec struct {
		Cmd  []string `help:"command to run, flags always have to be in front" arg:"" optional:"" passthrough:""`
		At   string   `help:"event at which exec should trigger" short:"a"`
		Pre  string   `help:"pre range" short:"p"`
		Post string   `help:"post range" short:"q"`
	} `cmd:"" help:"execute command, if event trigger matches"`

	Ops struct {
		Filter string `help:"only execute operations whose conf files match the regex filter" short:"f" default:".*"`
	} `cmd:"" help:"list files matching the criteria"`
}

func parseArgs() {
	userdata := getUserdataMap()
	defaultConfigFolder := path.Join(userdata["home"], ".conf", appName)

	ctx := kong.Parse(&cli,
		kong.Name(appName),
		kong.Description(appDescription),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}),
		kong.Vars{
			"configFile": defaultConfigFolder + "/conf.yaml",
		},
	)
	_ = ctx.Run()
	// ctx.FatalIfErrorf(err)
	cli.Action = strings.Split(ctx.Command(), " ")[0]
	if cli.Action == "version" {
		printBuildTags(BUILDTAGS)
		os.Exit(0)
	}
}

func printBuildTags(buildtags string) {
	regexp, _ := regexp.Compile(`({|}|,)`)
	s := regexp.ReplaceAllString(buildtags, "\n")
	s = strings.Replace(s, "_subversion: ", "Version: "+appMainversion+".", -1)
	fmt.Printf("%s\n", s)
}

func getUserdataMap() map[string]string {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	m := make(map[string]string)
	m["user_id"] = user.Uid
	m["group_id"] = user.Gid
	m["username"] = user.Username
	m["name"] = user.Name
	m["home"] = user.HomeDir + "/"
	return m
}
