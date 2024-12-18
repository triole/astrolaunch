# Astrolaunch ![build](https://github.com/triole/astrolaunch/actions/workflows/build.yaml/badge.svg) ![test](https://github.com/triole/astrolaunch/actions/workflows/test.yaml/badge.svg)

## Synopsis

Astrolaunch is able to run commands based on the position of the sun. It compares the time at which is is run with the events in the configured operations. If the run time equals the event considering the given range, the commands are run. It enables you to run stuff based on the sun's position so to say.

Astrolaunch may get other event data in the future. For data listed below are available.

```go mdox-exec="r calc -d 20240601"
{
  "time": {
    "time": "2024-12-18T07:45:12.422695577Z"
  },
  "location": {
    "lat": 0,
    "lon": 0
  },
  "sun": {
    "dawn": "2024-12-18T06:31:54.012846336+01:00",
    "dusk": "2024-12-18T19:24:13.114789888+01:00",
    "golden_hour": "2024-12-18T18:31:54.012886528+01:00",
    "golden_hour_end": "2024-12-18T07:24:13.11474944+01:00",
    "nadir": "2024-12-18T00:58:03.56381824+01:00",
    "nautical_dawn": "2024-12-18T06:05:41.15314176+01:00",
    "nautical_dusk": "2024-12-18T19:50:25.974494208+01:00",
    "night": "2024-12-18T20:16:45.822935808+01:00",
    "night_end": "2024-12-18T05:39:21.304700416+01:00",
    "solar_noon": "2024-12-18T12:58:03.56381824+01:00",
    "sunrise": "2024-12-18T06:54:25.73132416+01:00",
    "sunrise_end": "2024-12-18T06:56:45.113187584+01:00",
    "sunset": "2024-12-18T19:01:41.396311808+01:00",
    "sunset_start": "2024-12-18T18:59:22.01444864+01:00"
  },
  "moon": {
    "rise": "2024-12-18T21:00:00Z",
    "set": "2024-12-18T08:00:00Z"
  }
}
```

## How to

Besides from your `location` the `operations` that you would like to run are configured in a configuration file. Every `operation` has four attributes.

| entry | explanation                                                |
|-------|------------------------------------------------------------|
| name  | operation name that will occur the logs                    |
| at    | event at which the operation should be run                 |
| range | time span in which the event is considered to be happening |
| exec  | list of commands to execute                                |

### At

Basically works like a JSON selector referring to the astro data available. See above or run `astrolaunch -a` to see this set of data. Examples: sun.dawn, sun.nadir etc.

### Range

Let's say astrolaunch is run two minutes before `sun.dusk`. Only if `range` is at least these two minutes, operations that should run at `sun.dusk` will be executed. Range only accounts to the time before. If an event is over, nothing will be run. Use a string to define a time range. Examples: 1h, 30m, 120s etc.

## Configuration

```go mdox-exec="tail -n+2 example/conf.yaml"
operations_dir: "{{.SELFDIR}}/operations"

location:
  lat: 51.50808063275697
  lon: -0.12806528535354245

default_range:
  pre: 2m
  post: 10m
```

## Help

```go mdox-exec="r -h"
Usage: astrolaunch <command> [flags]

launch commands at sun rise, sun dawn or other astro related times

Flags:
  -h, --help                      Show context-sensitive help.
  -c, --conf="/home/ole/.conf/astrolaunch/conf.yaml"
                                  path to config file
      --log-file="/dev/stdout"    log file
      --log-level="info"          log level
      --log-no-colors             disable output colours, print plain text
      --log-json                  enable json log, instead of text one
  -n, --dry-run                   dry run, just print operations that would run
  -V, --version-flag              display version

Commands:
  calc    list files matching the criteria
  exec    execute command, if event trigger matches
  ops     list files matching the criteria

Run "astrolaunch <command> --help" for more information on a command.
```
