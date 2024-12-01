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
		calc.Sun.Light[toSnakeCase(string(key))] = toLocalTime(val.Value)
	}

	res1 := suncalc.GetPosition(calc.Time.Time, calc.Location.Lat, calc.Location.Lon)
	calc.Sun.Position["altitude"] = res1.Altitude
	calc.Sun.Position["azimuth"] = res1.Azimuth

	res2 := suncalc.GetMoonTimes(now, calc.Location.Lat, calc.Location.Lon, false)
	calc.Moon.Light["rise"] = res2.Rise
	calc.Moon.Light["set"] = res2.Set
	calc.Moon.Light["always_up"] = res2.AlwaysUp
	calc.Moon.Light["always_down"] = res2.AlwaysDown

	res3 := suncalc.GetMoonPosition(now, calc.Location.Lat, calc.Location.Lon)
	calc.Moon.Position["altitude"] = res3.Altitude
	calc.Moon.Position["azimuth"] = res3.Azimuth
	calc.Moon.Position["distance"] = res3.Distance
	calc.Moon.Position["parallactic_angle"] = res3.ParallacticAngle

	res4 := suncalc.GetMoonIllumination(now)
	calc.Moon.Illumination["fraction"] = res4.Fraction
	calc.Moon.Illumination["phase"] = res4.Phase
	calc.Moon.Illumination["angle"] = res4.Angle
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
