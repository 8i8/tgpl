package gitish

// TODO these functions could be joined into one single function that returns a
// flag such that one single iteration over an array of dates be required.

// lessThanMonth returns true for dates that are less than a month apart.
func lessThanMonth(now, rec date) bool {
	diff := (now.y*12 + int(now.m)) - (rec.y*12 + int(rec.m))
	return diff == 0 || diff == 1 && now.d < rec.d
}

// monthToAYear returns true for dates that are greater than a month but less than
// a year apart.
func monthToAYear(now, rec date) bool {
	diff := (now.y*12 + int(now.m)) - (rec.y*12 + int(rec.m))
	return diff == 1 && now.d >= rec.d ||
		diff > 1 && diff < 12 ||
		diff == 12 && now.d < rec.d
}

// yearOnward returns true for dates that are greater than a year apart.
func yearOnward(now, rec date) bool {
	diff := (now.y*12 + int(now.m)) - (rec.y*12 + int(rec.m))
	return diff == 12 && now.d >= rec.d || diff > 12
}

func lessThanMonth2(now, rec date) bool {
	return now.y == rec.y && now.m == rec.m ||
		now.y == rec.y && now.m-1 == rec.m && now.d < rec.d ||
		now.y-1 == rec.y && now.m == 1 && rec.m == 12 && now.d < rec.d
}

func monthToAYear2(now, rec date) bool {
	return now.y == rec.y && now.m-1 == rec.m && now.d >= rec.d ||
		now.y == rec.y && now.m-1 > rec.m ||
		now.y-1 == rec.y && now.m < rec.m && rec.m-now.m != 11 ||
		now.y-1 == rec.y && now.m == 1 && rec.m == 12 && now.d >= rec.d ||
		now.y-1 == rec.y && now.m == rec.m && now.d < rec.d
}

func yearOnward2(now, rec date) bool {
	return now.y-1 > rec.y ||
		now.y-1 == rec.y && now.m > rec.m ||
		now.y-1 == rec.y && now.m == rec.m && now.d >= rec.d
}
