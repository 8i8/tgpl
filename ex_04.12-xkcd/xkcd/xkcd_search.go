package xkcd

import (
	"fmt"

	"8i8/ds"
)

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *   list
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

func searchList(t *ds.Trie, comics *DataBase, args []string) []uint {

	if VERBOSE {
		fmt.Printf("xkcd: starting search list\n")
	}

	// Count occurrence of each search word over all comics, used to filter
	// out comics that do not contain all of the required search words.
	results := t.SearchWordsList(args)

	// Add all indicies to a map and count occurance.
	m := make(ds.Count)
	for _, data := range results {
		for _, l := range data.Links {
			m[l]++
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
func scanComicMapList(m ds.MList, c Comic) ds.MList {

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
		t.AddList(word, indices)
	}
	return t
}

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
func scanComicMap(m ds.MBtree, c Comic) ds.MBtree {

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
func buildSearchMap(comics *DataBase) ds.MBtree {

	// Scan and map comics.
	m := make(ds.MBtree)

	for _, comic := range comics.Edition {
		scanComicMap(m, comic)
	}

	return m
}

// buildSearchTrie constructs a search trie from a search map.
func buildSearchTrie(m ds.MBtree) *ds.Trie {

	t := new(ds.Trie)
	for word, btree := range m {
		t.AddBtree(word, btree)
	}

	return t
}

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  Search
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// Search searches the local database of comic descriptions for the given
// arguments.
func (d *DataBase) Search(args []string) {
	if VERBOSE {
		fmt.Printf("xkcd: output start ~~~\n\n")
	}
	if LIST {
		m := buildSearchMapList(d)
		t := buildSearchTrieList(m)
		results := searchList(t, d, args)
		printResults(d, results)
	} else {
		m := buildSearchMap(d)
		t := buildSearchTrie(m)
		results := searchBtree(t, d, args)
		printResults(d, results)
	}
	if VERBOSE {
		fmt.Printf("\nxkcd: ~~~ output end\n")
	}
}
