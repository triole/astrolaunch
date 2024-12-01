package calc

import "time"

type Calc struct {
	Time     tTime     `json:"time" toml:"time" yaml:"time"`
	Location tLocation `json:"location" toml:"location" yaml:"location"`
	Sun      tSun      `json:"sun" toml:"sun" yaml:"sun"`
	Moon     tMoon     `json:"moon" toml:"moon" yaml:"moon"`
}

type tTime struct {
	Time time.Time `json:"time" toml:"time" yaml:"time"`
	// TZ   string    `json:"tz" toml:"tz" yaml:"tz"`
}

type tLocation struct {
	Lat float64 `json:"lat" toml:"lat" yaml:"lat"`
	Lon float64 `json:"lon" toml:"lon" yaml:"lon"`
}

type tSun struct {
	Light    tSunLight `json:"light" toml:"light" yaml:"light"`
	Position tPosition `json:"position" toml:"position" yaml:"position"`
}

type tMoon struct {
	Light        tMoonLight        `json:"light" toml:"light" yaml:"light"`
	Position     tPosition         `json:"position" toml:"position" yaml:"position"`
	Illumination tMoonIllumination `json:"illumination" toml:"illumination" yaml:"illumination"`
}

type tSunLight map[string]time.Time
type tMoonLight map[string]interface{}
type tMoonIllumination map[string]interface{}
type tPosition map[string]float64

func newDataset() (ds Calc) {
	ds.Sun.Light = make(tSunLight)
	ds.Sun.Position = make(tPosition)
	ds.Moon.Light = make(tMoonLight)
	ds.Moon.Position = make(tPosition)
	ds.Moon.Illumination = make(tMoonIllumination)
	return
}
