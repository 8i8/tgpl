package xkcd

import (
	"fmt"
	"sort"

	"tgpl/ex_04.12-xkcd/ds"
)

// scanComicMap runs extract words on every text field in a Comic struct.
func scanComicMapList(m ds.MList, c Comic) ds.MList {

	m = ds.ExtractAndMap(m, c.Link, c.Number)
	m = ds.ExtractAndMap(m, c.News, c.Number)
	m = ds.ExtractAndMap(m, c.SafeTitle, c.Number)
	m = ds.ExtractAndMap(m, c.Transcript, c.Number)
	m = ds.ExtractAndMap(m, c.Alt, c.Number)
	m = ds.ExtractAndMap(m, c.Title, c.Number)
	m = ds.IdToMap(m, c.Number)

	return m
}

// buildSearchMap scans the comic database and creates a map of all words
// found, linking them to the comics that they are from.
func buildSearchMapList(comics *DataBase) ds.MList {

	// Scan and map comics.
	m := make(ds.MList)

	for _, comic := range comics.Edition {
		scanComicMapList(m, comic)
	}

	return m
}

// buildSearchTrie constructs a search trie from a search map.
func buildSearchTrieList(m ds.MList) *ds.Trie {

	t := new(ds.Trie)
	for word, indices := range m {
		t.Add(word, indices)
	}
	return t
}

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  Search
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// searchList prepares a list of comics that contain the given search terms.
func searchList(t *ds.Trie, comics *DataBase, args []string) []uint {

	if VERBOSE {
		fmt.Printf("xkcd: starting search list\n")
	}

	// Count occurrence of each search word over all comics, used to filter
	// out comics that do not contain all of the required search words.
	datalist := t.SearchWords(args)

	// Add all indicies to a map and count occurance.
	m := make(map[uint]int)
	for _, data := range datalist {
		temp := make(map[uint]int)
		// Generate a map from all the linked btrees such that only one
		// instance of every comic index can exist per word searched.
		for _, id := range data.Links {
			temp[id] = 0
		}
		// Count instances of each index so as to isolate only the
		// comics that contain all words sought.
		for id := range temp {
			m[id]++
		}
	}

	// Extract from map only those indicies that contain every search term,
	// check the number of occurance of the id againt the number of search words.
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

	m := buildSearchMapList(d)
	t := buildSearchTrieList(m)
	results := searchList(t, d, cleanArgs(args))
	if TITLE {
		d.printTitleList(results)
	} else {
		d.printList(results)
	}

	if VERBOSE {
		fmt.Printf("\nxkcd: ~~~ output end\n")
	}
}
