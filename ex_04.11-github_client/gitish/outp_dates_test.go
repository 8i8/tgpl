package gitish

import (
	"testing"
	"time"
)

var calendar []date
var now date

func TestDates(t *testing.T) {

	// Months.
	for i := 0; i < 12; i++ {
		// Days, test for many different start days, 400 is just over a year
		// and a month.
		for j := 0; j < 40; j++ {

			for k := 0; k < 50; k++ {
				// Arbritrary starting point.
				d := time.Date(1980+k, 6, 10, 1, 0, 0, 0, time.UTC)

				// Augment date to cover the given time period.
				d = d.AddDate(0, i, j)
				now = date{0, d.Year(), d.Month(), d.Day(), false}

				// Create calendar that spans two years.
				for l := 730; l > 0; l-- {
					day := date{0, d.Year(), d.Month(), d.Day(), false}
					calendar = append(calendar, day)
					d = d.AddDate(0, 0, -1)
				}

				// Test the entire range of dates.
				for n, d := range calendar {
					c := testDateSort(now, d)
					if c&dERROR > 0 {
						t.Error("error: ", d.d, d.m, d.y, i, j)
						return
					}
					if c != dMORE && c != dYEAR && c != dMONTH {
						t.Error("doubled: ", d.d, d.m, d.y, i, j)
						return
					}
					calendar[n].p = true
				}

				// Check for dates that have not been printed.
				testNotPrinted(t, i, j, k)

				// reset the calendar
				calendar = calendar[:0]
			}
		}
	}
}

// testDateSort returns a bitfield with the flag set.
func testDateSort(now, rec date) cal {

	var c cal
	if yearOnward(now, rec) {
		c |= dMORE
	}
	if monthToAYear(now, rec) {
		c |= dYEAR
	}
	if lessThanMonth(now, rec) {
		c |= dMONTH
	}
	if c == 0 {
		c |= dERROR
	}
	return c
}

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  Benchmarks
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

func BenchmarkDate(b *testing.B) {

	for k := 0; k < b.N; k++ {
		// Months.
		for i := 0; i < 12; i++ {
			// Days, test for many different start days, 400 is just over a year
			// and a month.
			for j := 0; j < 400; j++ {

				// Arbritrary starting point.
				d := time.Date(2011, 6, 10, 1, 0, 0, 0, time.UTC)

				// Augment date to cover the given time period.
				d = d.AddDate(0, i, j)
				now = date{0, d.Year(), d.Month(), d.Day(), false}

				// Create calendar that spans two years.
				for k := 730; k > 0; k-- {
					day := date{0, d.Year(), d.Month(), d.Day(), false}
					calendar = append(calendar, day)
					d = d.AddDate(0, 0, -1)
				}

				for _, d := range calendar {
					dateSort(now, d)
				}

				// reset the calendar
				calendar = calendar[:0]
			}
		}
	}
}

func BenchmarkDateOld(b *testing.B) {

	for k := 0; k < b.N; k++ {
		// Months.
		for i := 0; i < 12; i++ {
			// Days, test for many different start days, 400 is just over a year
			// and a month.
			for j := 0; j < 400; j++ {

				// Arbritrary starting point.
				d := time.Date(2011, 6, 10, 1, 0, 0, 0, time.UTC)

				// Augment date to cover the given time period.
				d = d.AddDate(0, i, j)
				now = date{0, d.Year(), d.Month(), d.Day(), false}

				// Create calendar that spans two years.
				for k := 730; k > 0; k-- {
					day := date{0, d.Year(), d.Month(), d.Day(), false}
					calendar = append(calendar, day)
					d = d.AddDate(0, 0, -1)
				}

				for _, d := range calendar {
					//dateSortOld(now, d)
					testDateSort(now, d)
				}

				// reset the calendar
				calendar = calendar[:0]
			}
		}
	}
}

/*
   Print an error for any dates remaining that have not yet been printed.
*/
func testNotPrinted(t *testing.T, i, j, k int) {
	for _, d := range calendar {
		if d.p == false {
			t.Error("error: date not printed",
				d.d, d.m, d.y, i, j, k)
		}
	}
}

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  Old method left in code for benchmarking.
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// dateSort returns a bitfield with the flag set.
func dateSortOld(now, rec date) cal {

	switch {
	case yearOnwardOld(now, rec):
		return dMORE
	case monthToAYearOld(now, rec):
		return dYEAR
	case lessThanMonthOld(now, rec):
		return dMONTH
	default:
		return dERROR
	}
}

func lessThanMonthOld(now, rec date) bool {
	return now.y == rec.y && now.m == rec.m ||
		now.y == rec.y && now.m-1 == rec.m && now.d < rec.d ||
		now.y-1 == rec.y && now.m == 1 && rec.m == 12 && now.d < rec.d
}

func monthToAYearOld(now, rec date) bool {
	return now.y == rec.y && now.m-1 == rec.m && now.d >= rec.d ||
		now.y == rec.y && now.m-1 > rec.m ||
		now.y-1 == rec.y && now.m < rec.m && rec.m-now.m != 11 ||
		now.y-1 == rec.y && now.m == 1 && rec.m == 12 && now.d >= rec.d ||
		now.y-1 == rec.y && now.m == rec.m && now.d < rec.d
}

func yearOnwardOld(now, rec date) bool {
	return now.y-1 > rec.y ||
		now.y-1 == rec.y && now.m > rec.m ||
		now.y-1 == rec.y && now.m == rec.m && now.d >= rec.d
}
