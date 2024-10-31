package launch

import (
	"fmt"
	"testing"
	"time"
)

func TestCalcRangeDiff(t *testing.T) {
	assertCalcRangeDiff(
		toTime("2024-01-01T10:00:00+01:00"),
		toTime("2024-01-01T10:00:00+01:00"),
		toDur("2m"), toDur("2h"), true, t,
	)
	assertCalcRangeDiff(
		toTime("2024-01-02T08:00:00+01:00"),
		toTime("2024-01-02T10:00:00+01:00"),
		toDur("2m"), toDur("2h"), true, t,
	)
	assertCalcRangeDiff(
		toTime("2024-01-03T10:00:00+01:00"),
		toTime("2024-01-03T08:00:00+01:00"),
		toDur("2h"), toDur("2m"), true, t,
	)
	assertCalcRangeDiff(
		toTime("2024-01-02T08:00:00+01:00"),
		toTime("2024-01-02T10:00:00+01:00"),
		toDur("2m"), toDur("1h"), false, t,
	)
	assertCalcRangeDiff(
		toTime("2024-01-03T10:00:00+01:00"),
		toTime("2024-01-03T08:00:00+01:00"),
		toDur("1h"), toDur("2m"), false, t,
	)
}

func assertCalcRangeDiff(t1, t2 time.Time, r1, r2 time.Duration, exp bool, t *testing.T) {
	la := new()
	diff, match := la.calcRangeDiff(t1, t2, r1, r2)
	if match != exp {
		t.Errorf(
			"calcRangeDiff fail: %v, %v, %v, %v, diff: %v",
			la.printTime(t1), la.printTime(t2), r1, r2, diff,
		)
	}
}

func toDur(str string) (dur time.Duration) {
	la := new()
	dur, _ = la.str2dur(str)
	return
}

func toTime(str string) (tim time.Time) {
	loc, _ := time.LoadLocation("Europe/Berlin")
	tim, err := time.ParseInLocation(time.RFC3339, str, loc)
	if err != nil {
		fmt.Println("time parse error: " + err.Error())
	}
	return
}
