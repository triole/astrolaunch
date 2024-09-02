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
		if strings.EqualFold(arr[0], "sun") {
			tim, err = getVal(arr[1], calc.Sun)
		}
	}
	return
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
