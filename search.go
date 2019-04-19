package xkcd

import (
	"fmt"
	"sort"
	"unicode"
	"unicode/utf8"

	"tgpl/ex_04.12-xkcd/trie"
)

// nSearch is a map containing arrays of comic index numbers.
type mSearch map[string][]uint

// addToMap checks a mSearch map for a prexisting entry and returns the map with
// the number add if it was not already present.
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

// extractWords adds every word in a string to a mSearch map.
func extractWords(m mSearch, s string, n uint) mSearch {

	var word []rune
	var isInWord bool
	b := []byte(s)

	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			word = append(word, unicode.ToLower(r))
			isInWord = true
		} else {
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

// scanComic runs extract words on every text field in a Comic struct.
func scanComic(m mSearch, c Comic) mSearch {

	m = extractWords(m, c.Link, c.Number)
	m = extractWords(m, c.News, c.Number)
	m = extractWords(m, c.SafeTitle, c.Number)
	m = extractWords(m, c.Transcript, c.Number)
	m = extractWords(m, c.Alt, c.Number)
	m = extractWords(m, c.Title, c.Number)

	return m
}

// buildSearchMap scans the comic database and creates a map of all words
// found, linking them to the comics that they are from.
func buildSearchMap(comics Comics) mSearch {

	// Scan and map comics.
	m := make(mSearch)

	for _, comic := range comics.Edition {
		scanComic(m, comic)
	}

	return m
}

func buildSearchGraph(m mSearch) {
	t := trie.Trie{}

	for word, _ := range m {
		t.Add(word)
	}
	fmt.Print("trie made\n")
}

// search prepares a list of all comics that contain all of the given search
// terms.
func search(m mSearch, comics Comics, args []string) []uint {

	var results []uint
	count := make(map[uint]int)

	// Count occurrence of each search word over all comics, used to filter
	// out comics that do not contain all of the required search words.
	for _, arg := range args {
		indices := m[arg]
		for _, num := range indices {
			count[num]++
		}
	}

	// If the comic contains the same number of found words as the length
	// of the list of search terms, add the comic to the results.
	for num, i := range count {
		if i == len(args) {
			results = append(results, num-1)
		}
	}

	// Sort the results.
	sort.Slice(results, func(i, j int) bool {
		return results[j] > results[i]
	})

	return results
}
