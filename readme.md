# Astrolaunch ![build](https://github.com/triole/astrolaunch/actions/workflows/build.yaml/badge.svg) ![test](https://github.com/triole/astrolaunch/actions/workflows/test.yaml/badge.svg)

## Synopsis

Astrolaunch is able to run commands based on the position of the sun. It compares the time at which is is run with the events in the configured operations. If the run time equals the event considering the given range, the commands are run. It enables you to run stuff based on the sun's position so to say.

Astrolaunch may get other event data in the future. For data listed below are available.

```go mdox-exec="r calc -d 20240601"
{
  "time": {
    "time": "2024-06-01T00:00:00Z"
  },
  "location": {
    "lat": 52.5200066,
    "lon": 13.404954
  },
  "sun": {
    "dawn": "2024-06-01T04:03:10.163053568+02:00",
    "dusk": "2024-06-01T22:08:15.888291328+02:00",
    "golden_hour": "2024-06-01T20:26:13.681104128+02:00",
    "golden_hour_end": "2024-06-01T05:45:12.370240768+02:00",
    "nadir": "2024-06-01T01:05:43.025672448+02:00",
    "nautical_dawn": "2024-06-01T02:49:58.605823488+02:00",
    "nautical_dusk": "2024-06-01T23:21:27.445521152+02:00",
    "night": "0001-01-01T00:53:28+00:53",
    "night_end": "0001-01-01T00:53:28+00:53",
    "solar_noon": "2024-06-01T13:05:43.025672448+02:00",
    "sunrise": "2024-06-01T04:50:47.40565248+02:00",
    "sunrise_end": "2024-06-01T04:55:18.451255552+02:00",
    "sunset": "2024-06-01T21:20:38.64569216+02:00",
    "sunset_start": "2024-06-01T21:16:07.600089344+02:00"
  },
  "moon": {
    "rise": "2024-06-01T00:00:00Z",
    "set": "2024-06-01T13:00:00Z"
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

## Configurations

## Main config

```go mdox-exec="tail -n+2 example/conf.yaml"
operations_dir: "{{.SELFDIR}}/operations"

location:
  lat: 51.50808063275697
  lon: -0.12806528535354245

default_range:
  pre: 2m
  post: 10m
```

## Operations config

```go mdox-exec="tail -n+2 example/operations/test1.yaml"
name: test
at: sun.dawn
range:
  pre: 2m
  post: 4h
exec:
  - ["date"]
  - ["echo", "this is dawn"]
```

## Usage examples

```go mdox-exec="tail -n+3 example/usage.sh"
astrolaunch calc -d 20241112
astrolaunch calc -d 20241112 -r 3

astrolaunch exec -a sun.dawn -p 1m -q 2h ls -la /tmp

astrolaunch ops -f test
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
  calc       list files matching the criteria
  exec       execute command, if event trigger matches
  ops        list files matching the criteria
  version    display version

Run "astrolaunch <command> --help" for more information on a command.
```
