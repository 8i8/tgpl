package github

/*
   1) The same calendar year and month.
   2) The same calendar year with one month difference; lesser day than the day
  	issued.
   3) One calendar year difference; January issue date December; lesser day
  	today than the day issued.
*/
func lessThanMonth(now, rec date) bool {
	return now.y == rec.y && now.m == rec.m ||
		now.y == rec.y && now.m-1 == rec.m && now.d < rec.d ||
		now.y-1 == rec.y && now.m == 1 && rec.m == 12 && now.d < rec.d
}

/*
   1) The same calendar year; Greater month by more than one calendar month
  	than the date issued.
   2) One calendar years difference; The current calendar month is lesser than
  	the month of the issue.
   3) The same calendar year; The month is one greater than the issued month;
  	equal or greater day today than the day issued.
   4) One calendar years difference; The same calendar month; lesser day today
  	than the day issued.
   5) One calendar years difference; January the date of issue in December;
  	Equal or greater day today than the day issued.
*/
func monthToAYear(now, rec date) bool {
	return now.y == rec.y && now.m-1 == rec.m && now.d >= rec.d ||
		now.y == rec.y && now.m-1 > rec.m ||
		now.y-1 == rec.y && now.m < rec.m && rec.m-now.m != 11 ||
		now.y-1 == rec.y && now.m == 1 && rec.m == 12 && now.d >= rec.d ||
		now.y-1 == rec.y && now.m == rec.m && now.d < rec.d
}

/*
   1) Greater than one year differance from the year issued.
   2) One calendar years differance; The month is greater than that of the month issued.
   3) One calendar years difference; The same calendar month; Equal or greater
  	day than the day issued.
*/
func yearOnward(now, rec date) bool {
	return now.y-1 > rec.y ||
		now.y-1 == rec.y && now.m > rec.m ||
		now.y-1 == rec.y && now.m == rec.m && now.d >= rec.d
}
