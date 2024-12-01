package calc

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func (calc Calc) GetTime(ra string) (tim time.Time, err error) {
	arr := strings.Split(ra, ".")
	if len(arr) > 1 {
		cond := strings.ToLower(arr[0])
		switch cond {
		case "sun":
			tim, err = getVal(arr[1], calc.Sun)
		case "moon":
			tim, err = getVal(arr[1], calc.Moon)
		}
	}
	if tim.Unix() < 0 {
		err = errors.New("no astro event time for '" + arr[0] + "'")
	}
	tim = toLocalTime(tim)
	return
}

func toLocalTime(inp time.Time) time.Time {
	now := time.Now()
	location := now.Location()
	loc, err := time.LoadLocation(location.String())
	if err != nil {
		panic(err)
	}
	return inp.In(loc)
}

func getVal(s string, m map[string]time.Time) (val time.Time, err error) {
	var ok bool
	if val, ok = m[s]; ok {
		val = m[s]
	} else {
		err = errors.New(
			"can not fetch value from map: " + s + ", " +
				fmt.Sprintf("%+v", m),
		)
	}
	return
}
