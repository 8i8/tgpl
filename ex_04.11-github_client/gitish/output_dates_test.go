package gitish

import (
	"testing"
	"time"
)

var cal []date
var now date

func TestDates(t *testing.T) {

	// Test all months.
	for i := 0; i < 12; i++ {
		// Test for many different start days.
		for j := 0; j < 1000; j++ {

			// Set the start date.
			d := time.Date(2011, 1, 10, 1, 0, 0, 0, time.UTC)
			// Augment date to cover a time period.
			d = d.AddDate(0, i, j)
			now = date{0, d.Year(), d.Month(), d.Day(), false}

			// Set an array of time periods of greater than two
			// years.
			for k := 730; k > 0; k-- {
				day := date{0, d.Year(), d.Month(), d.Day(), false}
				cal = append(cal, day)
				d = d.AddDate(0, 0, -1)
			}

			// Tests
			testLessThanMonth(t, i, j)
			testMonthToAYear(t, i, j)
			testYearOnward(t, i, j)
			testNotPrinted(t, i, j)
			cal = cal[:0]
		}
	}
}

func testLessThanMonth(t *testing.T, i, j int) {
	for n, d := range cal {
		if lessThanMonth(now, d) {
			if d.p == true {
				t.Error("error: LessThanMonth already printed",
					d.d, d.m, d.y, i, j)
			}
			cal[n].p = true
		}
	}
}

func testMonthToAYear(t *testing.T, i, j int) {
	for n, d := range cal {
		if monthToAYear(now, d) {
			if d.p == true {
				t.Error("error: MonthToAYear already printed",
					d.d, d.m, d.y, i, j)
			}
			cal[n].p = true
		}
	}
}

func testYearOnward(t *testing.T, i, j int) {
	for n, d := range cal {
		if yearOnward(now, d) {
			if d.p == true {
				t.Error("error: YearOnward already printed",
					d.d, d.m, d.y, i, j)
			}
			cal[n].p = true
		}
	}
}

/*
   Print an error for any dates remaining that have not yet been printed.
*/
func testNotPrinted(t *testing.T, i, j int) {
	for _, d := range cal {
		if d.p == false {
			t.Error("error: date not printed",
				d.d, d.m, d.y, i, j)
		}
	}
}
