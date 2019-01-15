package main

const LEN = 13

// reverse reverses a slice of int's in place.
func reverseSlice(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// Reverse for an array pointer requires that the arrays length be defined in
// along with the type.
func reverseArray(s *[LEN]int) *[LEN]int {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
	return s
}

// When the type is defined with its length any loops in the function may also
// be unwound.
func reverseArrayUnwound(s *[13]int) *[13]int {

	(*s)[0], (*s)[12] = (*s)[12], (*s)[0]
	(*s)[1], (*s)[11] = (*s)[11], (*s)[1]
	(*s)[2], (*s)[10] = (*s)[10], (*s)[2]
	(*s)[3], (*s)[9] = (*s)[9], (*s)[3]
	(*s)[4], (*s)[8] = (*s)[8], (*s)[4]
	(*s)[5], (*s)[7] = (*s)[7], (*s)[5]

	return s
}

func reverseArrayCopyUnwound(s [13]int) [13]int {

	s[0], s[12] = s[12], s[0]
	s[1], s[11] = s[11], s[1]
	s[2], s[10] = s[10], s[2]
	s[3], s[9] = s[9], s[3]
	s[4], s[8] = s[8], s[4]
	s[5], s[7] = s[7], s[5]

	return s
}
