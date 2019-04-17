package xkcd

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"8i8/ds"
)

// lex extracts every word from a string.
func lex(trie *ds.Trie, s string, n uint) *ds.Trie {

	var word []rune
	var isInWord bool
	b := []byte(s)

	// Build word rune by rune, all lower case.
	for len(b) > 0 {
		// If rune is a letter add to byte slice and indicate the state
		// of bing in a word.
		r, size := utf8.DecodeRune(b)
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			word = append(word, unicode.ToLower(r))
			isInWord = true
		} else {
			// Break; If reading a word; Add both the word and it's
			// index of origin to the map. Skipping over any non
			// lexical characters.
			if isInWord {
				trie.Add(string(word), n)
				word = word[:0]
				isInWord = false
			}
		}
		b = b[size:]
	}
	return trie
}

// scanComic runs extract words on every text field in a Comic struct.
func scan(t *ds.Trie, c Comic) *ds.Trie {

	t = lex(t, c.Link, c.Number)
	t = lex(t, c.News, c.Number)
	t = lex(t, c.SafeTitle, c.Number)
	t = lex(t, c.Transcript, c.Number)
	t = lex(t, c.Alt, c.Number)
	t = lex(t, c.Title, c.Number)

	return t
}

// buildSearchGraph scans the comic database and creates a map of all words
// found, linking them to the comics that they are from.
func buildSearchGraph(comics *DataBase) *ds.Trie {

	if VERBOSE {
		fmt.Print("xkcd: building search data structure\n")
	}

	t := &ds.Trie{}
	for _, comic := range comics.Edition {
		scan(t, comic)
	}

	if VERBOSE {
		fmt.Print("xkcd: search data structure complete\n")
	}

	return t
}

// search prepares a list of comics that contain the given search terms.
func search(t *ds.Trie, comics *DataBase, args []string) []uint {

	if VERBOSE {
		fmt.Printf("xkcd: starting search list\n")
	}

	var results []uint

	// Count occurrence of each search word over all comics, used to filter
	// out comics that do not contain all of the required search words.
	btrees := t.SearchWords(args)

	// Add all indicies to a map and count occurance.
	m := make(ds.Count)
	for _, data := range btrees {
		for _, btree := range data.LinkedIds {
			// if VERBOSE {
			// 	fmt.Println("btree:")
			// 	ds.BtreePrint(btree)
			// }
			ds.BtreeCount(&m, btree)
		}
	}

	// Extract from map only those indicies that contain every search term,
	// check the number of occurance of the id againt the number of search words.
	var filter []uint
	l := len(args)
	for id, count := range m {

		// if VERBOSE {
		// 	fmt.Printf("output: id: %d count: %d\n", id, count)
		// }
		if count == l {
			filter = append(filter, id)
		}
	}
	if VERBOSE {
		fmt.Printf("filter: %v\n", filter)
	}

	// ~~~

	// If the comic contains the same number of found words as the length
	// of the list of search terms, add the comic to the results.
	// for num, i := range m {
	// 	if i == len(args) {
	// 		results = append(results, num-1)
	// 	}
	// }

	// Sort the results.
	// sort.Slice(results, func(i, j int) bool {
	// 	return results[j] > results[i]
	// })

	if VERBOSE {
		fmt.Printf("xkcd: search list complete\n")
	}

	return results
}

// Search searched the local database of comic descriptions for the given
// arguments.
func (d *DataBase) Search(args []string) {
	if VERBOSE {
		fmt.Printf("xkcd: output start ~~~\n\n")
	}
	t := buildSearchGraph(d)
	_ = search(t, d, args)
	// results := search(t, d, args)
	// printResults(d, results)
	if VERBOSE {
		fmt.Printf("\nxkcd: ~~~ output end\n")
	}
}
