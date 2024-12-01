package calc

import (
	"strings"
	"time"

	"github.com/sixdouglas/suncalc"
)

type Calc struct {
	Sun map[string]time.Time `json:"sun"`
}

func Init(now time.Time, lat, lon float64) (calc Calc) {
	calc.Sun = make(map[string]time.Time)
	arr := suncalc.GetTimes(now, lat, lon)
	for key, val := range arr {
		calc.Sun[toSnakeCase(string(key))] = toLocalTime(val.Value)
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
