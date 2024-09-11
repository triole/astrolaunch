# Astrolaunch

## Synopsis

Astrolaunch is able to run commands based on the position of the sun. It compares the time at which is is run with the events in the configured operations. If the run time equals the event considering the given range, the commands are run. It enables you to run stuff based on the sun's position so to say.

Astrolaunch may get other event data in the future. For data listed below are available.

```go mdox-exec="r -a"
{
  "sun": {
    "dawn": "2024-09-11T04:00:55.59912192Z",
    "dusk": "2024-09-11T18:07:53.54798464Z",
    "golden_hour": "2024-09-11T16:48:05.580553216Z",
    "golden_hour_end": "2024-09-11T05:20:43.566553344Z",
    "nadir": "2024-09-10T23:04:24.573553152Z",
    "nautical_dawn": "2024-09-11T03:19:00.147990016Z",
    "nautical_dusk": "2024-09-11T18:49:48.999116544Z",
    "night": "2024-09-11T19:34:55.41903744Z",
    "night_end": "2024-09-11T02:33:53.72806912Z",
    "solar_noon": "2024-09-11T11:04:24.573553152Z",
    "sunrise": "2024-09-11T04:35:39.378726656Z",
    "sunrise_end": "2024-09-11T04:39:11.69851648Z",
    "sunset": "2024-09-11T17:33:09.768379904Z",
    "sunset_start": "2024-09-11T17:29:37.44859008Z"
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
    range: 300s
    exec:
      - ["date"]
      - ["echo", "this is dawn"]
  - name: test
    at: sun.dusk
    range: 300s
    exec:
      - ["date"]
      - ["echo", "this is dusk"]
```
