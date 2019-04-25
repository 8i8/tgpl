package ds

import (
	"strconv"
	"unicode"
	"unicode/utf8"
)

// MList is a map containing arrays of indices.
type MList map[string][]int

// addToMap checks a MList map for a preexisting entry and returns the
// map with the index and word added, if not already present.
func addToMap(m MList, s string, index int) MList {

	list := m[s]

	// If empty.
	if len(list) == 0 {
		list = append(list, index)
		m[s] = list
		return m
	}

	// If already there.
	for _, i := range list {
		if i == index {
			return m
		}
	}

	// Add the number.
	list = append(list, index)
	m[s] = list
	return m
}

// ExtractAndMap adds every word in a string to a MList map.
func ExtractAndMap(m MList, s string, index int) MList {

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
			// lexicographical characters.
			if isInWord {
				m = addToMap(m, string(word), index)
				word = word[:0]
				isInWord = false
			}
		}
		b = b[size:]
	}
	return m
}

func IdToMap(m MList, index int) MList {

	return addToMap(m, strconv.Itoa(int(index)), index)
}
