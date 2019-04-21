package xkcd

import (
	"8i8/ds"
	"fmt"
)

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *   btree
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// searchBtree prepares a list of comics that contain the given search terms.
func searchBtree(t *ds.Trie, comics *DataBase, args []string) []uint {

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
		// Generate a map from all the linked btrees such that only one
		// instance of every comic index can exist per word searched.
		for _, btree := range data.LinkedIds {
			ds.BtreeToMap(&temp, btree)
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
	if VERBOSE {
		fmt.Printf("xkcd: search list complete\n")
	}

	return filter
}

// scanComicMap runs extract words on every text field in a Comic struct.
func scanComicMap(m ds.MData, c Comic) ds.MData {

	m = ds.ExtractAndMap(m, c.Link, c.Number)
	m = ds.ExtractAndMap(m, c.News, c.Number)
	m = ds.ExtractAndMap(m, c.SafeTitle, c.Number)
	m = ds.ExtractAndMap(m, c.Transcript, c.Number)
	m = ds.ExtractAndMap(m, c.Alt, c.Number)
	m = ds.ExtractAndMap(m, c.Title, c.Number)

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

// buildSearchTrie constructs a search trie from a search map.
func buildSearchTrie(m ds.MData) *ds.Trie {

	t := new(ds.Trie)
	for word, data := range m {
		t.AddBtree(word, data.Btree)
	}

	return t
}
