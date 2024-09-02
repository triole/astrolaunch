package conf

import (
	"time"

	"github.com/triole/logseal"
)

type Conf struct {
	Now      Now
	FileName string
	Content  ConfContent
	Lg       logseal.Logseal
	DryRun   bool
}

type Now struct {
	UTC   time.Time
	Local time.Time
}

type ConfContent struct {
	Location   Location    `yaml:"location"`
	Operations []Operation `yaml:"operations"`
}
type Operation struct {
	Name  string     `yaml:"name"`
	At    string     `yaml:"at"`
	Range string     `yaml:"range"`
	Exec  [][]string `yaml:"exec"`
}

type Location struct {
	Lat float64
	Lon float64
}
