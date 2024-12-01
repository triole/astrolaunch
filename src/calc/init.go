package calc

import (
	"strings"
	"time"

	"github.com/sixdouglas/suncalc"
)

func Init(now time.Time, lat, lon float64) (calc Calc) {
	calc = newDataset()
	calc.Time.Time = now
	calc.Location.Lat = lat
	calc.Location.Lon = lon
	arr := suncalc.GetTimes(now, lat, lon)
	for key, val := range arr {
		calc.Sun[toSnakeCase(string(key))] = toLocalTime(val.Value)
	}
	res2 := suncalc.GetMoonTimes(now, calc.Location.Lat, calc.Location.Lon, false)
	calc.Moon["rise"] = res2.Rise
	calc.Moon["set"] = res2.Set
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
