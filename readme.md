# Astrolaunch ![build](https://github.com/triole/astrolaunch/actions/workflows/build.yaml/badge.svg) ![test](https://github.com/triole/astrolaunch/actions/workflows/test.yaml/badge.svg)

## Synopsis

Astrolaunch is able to run commands based on the position of the sun. It compares the time at which is is run with the events in the configured operations. If the run time equals the event considering the given range, the commands are run. It enables you to run stuff based on the sun's position so to say.

Astrolaunch may get other event data in the future. For data listed below are available.

```go mdox-exec="r -a -d 20240601"
{
  "sun": {
    "dawn": "2024-06-01T02:03:10.163053568Z",
    "dusk": "2024-06-01T20:08:15.888291328Z",
    "golden_hour": "2024-06-01T18:26:13.681104128Z",
    "golden_hour_end": "2024-06-01T03:45:12.370240768Z",
    "nadir": "2024-05-31T23:05:43.025672448Z",
    "nautical_dawn": "2024-06-01T00:49:58.605823488Z",
    "nautical_dusk": "2024-06-01T21:21:27.445521152Z",
    "night": "0001-01-01T00:00:00Z",
    "night_end": "0001-01-01T00:00:00Z",
    "solar_noon": "2024-06-01T11:05:43.025672448Z",
    "sunrise": "2024-06-01T02:50:47.40565248Z",
    "sunrise_end": "2024-06-01T02:55:18.451255552Z",
    "sunset": "2024-06-01T19:20:38.64569216Z",
    "sunset_start": "2024-06-01T19:16:07.600089344Z"
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
location:
  lat: 51.50808063275697
  lon: -0.12806528535354245

operations:
  - name: test
    at: sun.dawn
    range:
      pre: 2m
      post: 4h
    exec:
      - ["date"]
      - ["echo", "this is dawn"]
  - name: test
    at: sun.dusk
    range:
      pre: 2m
      post: 4h
    exec:
      - ["date"]
      - ["echo", "this is dusk"]
```

## Help

```go mdox-exec="r -h"
Usage: astrolaunch [flags]

launch commands at sun rise, sun dawn or other astro related times

Flags:
  -h, --help                      Show context-sensitive help.
  -c, --conf="/home/ole/.conf/astrolaunch/conf.yaml"
                                  path to config file
      --log-file="/dev/stdout"    log file
  -a, --astro                     only print astro calculation results
  -d, --date=STRING               print astro calculation for a certain date,
                                  format: YYYYMMDD
      --log-level="info"          log level
      --log-no-colors             disable output colours, print plain text
      --log-json                  enable json log, instead of text one
  -n, --dry-run                   dry run, just print operations that would run
  -V, --version-flag              display version
```
