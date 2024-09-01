package conf

import (
	"time"

	"astrolaunch/src/astro"

	"github.com/triole/logseal"
)

type Conf struct {
	FileName string
	Content  ConfContent
	Now      time.Time
	Astro    astro.Astro
	Lg       logseal.Logseal
}

type ConfContent struct {
	Location   Location    `yaml:"location"`
	Operations []Operation `yaml:"operations"`
}
type Operation struct {
	Name      string     `yaml:"name"`
	RunAt     string     `yaml:"run_at"`
	Tolerance string     `yaml:"tolerance"`
	Commands  [][]string `yaml:"commands"`
}

type Location struct {
	Lat float64
	Lon float64
}
