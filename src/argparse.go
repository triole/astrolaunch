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

var CLI struct {
	Conf        string `help:"path to config file" short:"c" default:"${configFile}"`
	LogFile     string `help:"log file" default:"/dev/stdout"`
	Astro       bool   `help:"only print astro calculation results" short:"a"`
	LogLevel    string `help:"log level" default:"info" enum:"trace,debug,info,error"`
	LogNoColors bool   `help:"disable output colours, print plain text"`
	LogJSON     bool   `help:"enable json log, instead of text one"`
	DryRun      bool   `help:"dry run, just print operations that would run" short:"n"`
	VersionFlag bool   `help:"display version" short:"V"`
}

func parseArgs() {
	userdata := getUserdataMap()
	defaultConfigFolder := path.Join(userdata["home"], ".conf", appName)

	ctx := kong.Parse(&CLI,
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

	if CLI.VersionFlag {
		printBuildTags(BUILDTAGS)
		os.Exit(0)
	}
	// ctx.FatalIfErrorf(err)
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
