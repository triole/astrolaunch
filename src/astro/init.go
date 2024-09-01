package astro

import (
	"strings"
	"time"

	"github.com/sixdouglas/suncalc"
)

type Astro struct {
	Sun map[string]time.Time
}

func Init(now time.Time, lat, lon float64) (astro Astro) {
	astro.Sun = make(map[string]time.Time)
	arr := suncalc.GetTimes(now, lat, lon)
	for key, val := range arr {
		astro.Sun[toSnakeCase(string(key))] = val.Value
	}
	return
}

func toSnakeCase(s string) string {
	var result string
	for i, v := range s {
		if i > 0 && v >= 'A' && v <= 'Z' {
			result += "_"
		}
		result += string(v)
	}
	return strings.ToLower(result)
}
