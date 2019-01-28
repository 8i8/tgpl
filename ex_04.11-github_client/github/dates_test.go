package github

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
		for j := 0; j < 10; j++ {

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

/*
   1) The same calendar month and year.
   2) The same calendar year with one month difference; lesser day than the day
  	issued.
   3) One calendar year difference; January issue date December; lesser day
  	today than the day issued.
*/
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

/*
   1) The same calendar year; Greater month by more than one calendar month
  	than the date issued.
   2) The same calendar year; The month is one greater than the issued month;
  	equal or greater day today than the day issued.
   3) One calendar years difference; January the date of issue in December;
  	Equal or greater day today than the day issued.
   4) One calendar years difference; The current calendar month is lesser than
  	the month of the issue.
   5) One calendar years difference; The same calendar month; lesser day today
  	than the day issued.
*/
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

/*
   1) One calendar years difference; The same calendar month: Equal or greater
  	day than the day issued.
   2) Greater than one years difference.
*/
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
