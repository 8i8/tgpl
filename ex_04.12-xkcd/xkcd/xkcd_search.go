package xkcd

import (
	"8i8/ds"
	"fmt"
	"sort"
)

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  Build map
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// scanComicMap runs extract words on every text field in a Comic struct.
func scanComicMap(m ds.MData, c Comic) ds.MData {

	m = ds.ExtractStrings(m, c.Link, c.Number)
	m = ds.ExtractStrings(m, c.News, c.Number)
	m = ds.ExtractStrings(m, c.SafeTitle, c.Number)
	m = ds.ExtractStrings(m, c.Transcript, c.Number)
	m = ds.ExtractStrings(m, c.Alt, c.Number)
	m = ds.ExtractStrings(m, c.Title, c.Number)

	return m
}

// buildSearchMap scans the comic database and creates a map of all words
// found, linking them to the comics that they are from.
func buildSearchMap(comics *DataBase) ds.MData {

	// Scan and map comics.
	m := make(ds.MData)

	for _, comic := range comics.Edition {
		scanComicMap(m, comic)
	}

	return m
}

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  Build trie
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// buildSearchTrie constructs a search trie from a map of words.
func buildSearchTrie(m ds.MData) *ds.Trie {

	t := new(ds.Trie)
	for word, data := range m {
		t.Add(word, data)
	}

	return t
}

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  Search
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// search prepares a list of comics that contain the given search terms.
func search(t *ds.Trie, comics *DataBase, args []string) []uint {

	if VERBOSE {
		fmt.Printf("xkcd: starting search list\n")
	}

	// Count occurrence of each search word over all comics, used to filter
	// out comics that do not contain all of the required search words.
	datalist := t.SearchWords(args)

	// Add all indices to a map and count occurrence.
	m := make(ds.Count)
	for _, data := range datalist {
		temp := make(ds.Count)
		// Generate a map from all the linked data such that only one
		// instance of every comic index can exist per word searched.
		for _, data := range data.Datalist {
			for _, id := range data.List {
				temp[id] = 0
			}
		}
		// Count instances of each index so as to isolate only the
		// comics that contain all words sought.
		for id, _ := range temp {
			m[id]++
		}
	}

	// Extract from map only those indices that contain every search term,
	// check the number of occurrence of the id against the number of search words.
	var filter []uint
	l := len(args)
	for id, count := range m {
		if count == l {
			filter = append(filter, id)
		}
	}
	sort.Slice(filter, func(i, j int) bool {
		return filter[i] < filter[j]
	})

	if VERBOSE {
		fmt.Printf("xkcd: search list complete\n")
	}

	return filter
}

// Search searches the local database of comic descriptions for the given
// arguments.
func (d *DataBase) Search(args []string) {
	if VERBOSE {
		fmt.Printf("xkcd: output start ~~~\n\n")
	}

	m := buildSearchMap(d)
	t := buildSearchTrie(m)
	results := search(t, d, cleanArgs(args))
	printResults(d, results)

	if VERBOSE {
		fmt.Printf("\nxkcd: ~~~ output end\n")
	}
}
