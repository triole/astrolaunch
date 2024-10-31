# Astrolaunch ![build](https://github.com/triole/astrolaunch/actions/workflows/build.yaml/badge.svg) ![test](https://github.com/triole/astrolaunch/actions/workflows/test.yaml/badge.svg)

## Synopsis

Astrolaunch is able to run commands based on the position of the sun. It compares the time at which is is run with the events in the configured operations. If the run time equals the event considering the given range, the commands are run. It enables you to run stuff based on the sun's position so to say.

Astrolaunch may get other event data in the future. For data listed below are available.

```go mdox-exec="r -a"
{
  "sun": {
    "dawn": "2024-10-31T05:26:25.503906048Z",
    "dusk": "2024-10-31T16:16:00.892179712Z",
    "golden_hour": "2024-10-31T14:49:15.091085824Z",
    "golden_hour_end": "2024-10-31T06:53:11.304999424Z",
    "nadir": "2024-10-30T22:51:13.198042624Z",
    "nautical_dawn": "2024-10-31T04:46:07.122052608Z",
    "nautical_dusk": "2024-10-31T16:56:19.274032896Z",
    "night": "2024-10-31T17:35:53.138203904Z",
    "night_end": "2024-10-31T04:06:33.2578816Z",
    "solar_noon": "2024-10-31T10:51:13.198042624Z",
    "sunrise": "2024-10-31T06:02:23.352712704Z",
    "sunrise_end": "2024-10-31T06:06:11.714579968Z",
    "sunset": "2024-10-31T15:40:03.043373056Z",
    "sunset_start": "2024-10-31T15:36:14.681505536Z"
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
      --log-level="info"          log level
      --log-no-colors             disable output colours, print plain text
      --log-json                  enable json log, instead of text one
  -n, --dry-run                   dry run, just print operations that would run
  -V, --version-flag              display version
```
