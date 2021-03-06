package github

type cal uint

const (
	dMONTH cal = 1 << iota
	dYEAR
	dMORE
	dERROR
)

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

// dateSort returns a bitfield with the flag set.
func dateSort(now, rec date) cal {

	switch {
	case yearOnward(now, rec):
		return dMORE
	case monthToAYear(now, rec):
		return dYEAR
	case lessThanMonth(now, rec):
		return dMONTH
	default:
		return dERROR
	}
}
