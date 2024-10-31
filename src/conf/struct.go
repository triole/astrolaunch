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
	Range Range      `yaml:"range"`
	Exec  [][]string `yaml:"exec"`
}

type Range struct {
	Pre  string `yaml:"pre"`
	Post string `yaml:"post"`
}

type Location struct {
	Lat float64
	Lon float64
}
