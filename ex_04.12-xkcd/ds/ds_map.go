package ds

import (
	"unicode"
	"unicode/utf8"
)

// nSearch is a map containing arrays of comic index numbers.
type mSearch map[string][]uint

// addToMap checks a mSearch map for a prexisting entry and returns the map with
// the index and word added, if not already present.
func addToMap(m mSearch, s string, n uint) mSearch {

	list := m[s]

	// If empty.
	if len(list) == 0 {
		list = append(list, n)
		m[s] = list
		return m
	}

	// If already there.
	for _, num := range list {
		if num == n {
			return m
		}
	}
	// Add the number.
	list = append(list, n)
	m[s] = list
	return m
}

// extractMap adds every word in a string to a mSearch map.
func extractMap(m mSearch, s string, n uint) mSearch {

	var word []rune
	var isInWord bool
	b := []byte(s)

	// Build word rune by rune, all lower case.
	for len(b) > 0 {
		// If rune is a letter add to byte slice and signify the state
		// of processing of a word.
		r, size := utf8.DecodeRune(b)
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			word = append(word, unicode.ToLower(r))
			isInWord = true
		} else {
			// Break, If reading a word, add both it and its index
			// of origin to the map. Else, skip over any non
			// lexical characters.
			if isInWord {
				m = addToMap(m, string(word), n)
				word = word[:0]
				isInWord = false
			}
		}
		b = b[size:]
	}
	return m
}
