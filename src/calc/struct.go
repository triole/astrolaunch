package calc

import "time"

type Calc struct {
	Time     tTime      `json:"time" toml:"time" yaml:"time"`
	Location tLocation  `json:"location" toml:"location" yaml:"location"`
	Sun      tSunLight  `json:"sun" toml:"sun" yaml:"sun"`
	Moon     tMoonLight `json:"moon" toml:"moon" yaml:"moon"`
}

type tTime struct {
	Time time.Time `json:"time" toml:"time" yaml:"time"`
	// TZ   string    `json:"tz" toml:"tz" yaml:"tz"`
}

type tLocation struct {
	Lat float64 `json:"lat" toml:"lat" yaml:"lat"`
	Lon float64 `json:"lon" toml:"lon" yaml:"lon"`
}

type tSunLight map[string]time.Time
type tMoonLight map[string]time.Time

func newDataset() (ds Calc) {
	ds.Sun = make(tSunLight)
	ds.Moon = make(tMoonLight)
	return
}
