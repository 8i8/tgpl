package main

import (
	"strings"
)

// Nonempty is an example of an in-place slice algorithm.

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// The nonempty function can also be written using append
func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

// an in-place function to eliminate adjacent duplicates in a []string
// slice.
func adjDuplicates(str []string) []string {
	var s1 string
	i := 0
	for _, s2 := range str {
		if strings.Compare(s2, s1) != 0 {
			str[i] = s2
			i++
		}
		s1 = s2
	}
	return str[:i]
}
