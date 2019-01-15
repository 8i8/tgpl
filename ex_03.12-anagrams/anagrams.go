package main

import (
	"bytes"
	"fmt"
	"unicode"
	"unicode/utf8"
)

// Remove all but letters and put those into lower case.
func readyString(s string) string {
	var buf bytes.Buffer
	runes := []rune(s)

	for _, r := range runes {
		if unicode.IsLetter(r) {
			buf.WriteRune(unicode.ToLower(r))
		}
	}
	return buf.String()
}

// Check both s1 and s2 for character count to asses status as possible
// anagrams. Both must contain both the same quantity of characters as well as
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
	for key, _ := range m1 {
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
