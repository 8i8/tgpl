package xkcd

import (
	"8i8/ds"
	"fmt"
)

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *   list
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// searchList prepares a list of comics that contain the given search terms.
func searchList(t *ds.Trie, comics *DataBase, args []string) []uint {

	if VERBOSE {
		fmt.Printf("xkcd: starting search list\n")
	}

	// Count occurrence of each search word over all comics, used to filter
	// out comics that do not contain all of the required search words.
	datalist := t.SearchWordsList(args)

	// Add all indicies to a map and count occurance.
	m := make(ds.Count)
	for _, data := range datalist {
		temp := make(ds.Count)
		// Generate a map from all the linked btrees such that only one
		// instance of every comic index can exist per word searched.
		for _, id := range data.Links {
			temp[id] = 0
		}
		// Count instances of each index so as to isolate only the
		// comics that contain all words sought.
		for id, _ := range temp {
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
	if VERBOSE {
		fmt.Printf("xkcd: search list complete\n")
	}

	return filter
}

// scanComicMap runs extract words on every text field in a Comic struct.
func scanComicMapList(m ds.MData, c Comic) ds.MData {

	m = ds.ExtractAndMapList(m, c.Link, c.Number)
	m = ds.ExtractAndMapList(m, c.News, c.Number)
	m = ds.ExtractAndMapList(m, c.SafeTitle, c.Number)
	m = ds.ExtractAndMapList(m, c.Transcript, c.Number)
	m = ds.ExtractAndMapList(m, c.Alt, c.Number)
	m = ds.ExtractAndMapList(m, c.Title, c.Number)

	return m
}

// buildSearchMap scans the comic database and creates a map of all words
// found, linking them to the comics that they are from.
func buildSearchMapList(comics *DataBase) ds.MData {

	// Scan and map comics.
	m := make(ds.MData)

	for _, comic := range comics.Edition {
		scanComicMapList(m, comic)
	}

	return m
}

// buildSearchTrie constructs a search trie from a search map.
func buildSearchTrieList(m ds.MData) *ds.Trie {

	t := new(ds.Trie)
	for word, indices := range m {
		t.AddList(word, indices)
	}
	return t
}
