package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Remove all but letters and put those into lower case.
func readyString(s string) string {
	var buf bytes.Buffer

	for _, r := range s {
		if unicode.IsLetter(r) {
			buf.WriteRune(unicode.ToLower(r))
		}
	}
	return buf.String()
}

// Here the code uses a map to illimiate the case of double checking multiple
// instances of letters, it turns out to be slower than not using the map in
// the short length test cases that are being used.
func anagramNewMap(s1, s2 string) bool {

	// Remove all non alphabet characters and whitespace.
	s1 = readyString(s1)
	s2 = readyString(s2)

	// Check for equal length.
	l1 := utf8.RuneCountInString(s1)
	l2 := utf8.RuneCountInString(s2)
	if l1 != l2 {
		return false
	}

	// Check each individual rune.
	m := make(map[rune]bool)
	for _, r := range s1 {
		if !m[r] && strings.Count(s1, string(r)) != strings.Count(s2, string(r)) {
			return false
		}
		// add the rune to the map to avoid checking the same rune
		// twice.
		m[r] = true
	}
	return true
}

// This version is by far the most efficient.
func anagramNew(s1, s2 string) bool {

	// Remove all non alphabet characters and whitespace.
	s1 = readyString(s1)
	s2 = readyString(s2)

	// Check for equal length.
	l1 := utf8.RuneCountInString(s1)
	l2 := utf8.RuneCountInString(s2)
	if l1 != l2 {
		return false
	}

	// Check each individual rune.
	for _, r := range s1 {
		if strings.Count(s1, string(r)) != strings.Count(s2, string(r)) {
			return false
		}
	}
	return true
}

// Check both s1 and s2 for character count to assess status as possible
// anagram. Both must contain both the same quantity of characters as well as
// the same quantity of each; The strings are evaluated for unicode code points.
func anagram(s1, s2 string) bool {

	// Remove white space and change relevant characters to lower case.
	s1 = readyString(s1)
	s2 = readyString(s2)

	// Check for equal length.
	l1 := utf8.RuneCountInString(s1)
	l2 := utf8.RuneCountInString(s2)
	if l1 != l2 {
		return false
	}

	// Morph into runes to make note of unichar code points.
	r1 := []rune(s1)
	r2 := []rune(s2)

	m1 := make(map[string]int)
	m2 := make(map[string]int)

	// Add each rune to a map increment the map to count occurrences.
	for _, v := range r1 {
		m1[string(v)]++
	}
	for _, v := range r2 {
		m2[string(v)]++
	}

	// Check one map against the other for discrepancies.
	for key := range m1 {
		if m1[key] != m2[key] {
			return false
		}
	}
	return true
}

func main() {
	var a, b string

	a = "This is a test"
	b = "This is a test!"
	if anagram(a, b) {
		fmt.Printf("`%v` is an anagram of `%v`.\n", a, b)
	} else {
		fmt.Printf("`%v` and `%v` are not anagrams.\n", a, b)
	}

	a = "This also is a test"
	b = "This is a test!"
	if anagram(a, b) {
		fmt.Printf("`%v` is an anagram of `%v`.\n", a, b)
	} else {
		fmt.Printf("`%v` and `%v` are not anagrams.\n", a, b)
	}
}
