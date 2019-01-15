package main

const LEN = 6

// reverse reverses a slice of int's in place.
func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// Rotates the content of a slice by pairs.
func rotate(s []int) []int {
	// Rotate s left by two positions.
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	return s
}

// Rotate using arrays in a single pass.
func rotateSinglePass(s [LEN]int) [LEN]int {
	s[0], s[1], s[2], s[3], s[4], s[5] = s[2], s[3], s[4], s[5], s[0], s[1]
	return s
}
