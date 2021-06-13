package xkcd

import (
	"fmt"
	"sort"
	"xkcd/ds"
)

// scanComicMap runs extract words on every text field in a Comic struct.
func scanComicMapList(m ds.MList, c Comic) ds.MList {

	// Offset comic number for array index.
	index := c.Number - 1

	m = ds.ExtractAndMap(m, c.Link, index)
	m = ds.ExtractAndMap(m, c.News, index)
	m = ds.ExtractAndMap(m, c.SafeTitle, index)
	m = ds.ExtractAndMap(m, c.Transcript, index)
	m = ds.ExtractAndMap(m, c.Alt, index)
	m = ds.ExtractAndMap(m, c.Title, index)
	m = ds.IdToMap(m, index)

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

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  Search
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// searchXkcd prepares a list of comics that contain the given search terms.
func searchXkcd(t *ds.Trie, comics *DataBase, args []string) []int {

	if VERBOSE {
		fmt.Printf("xkcd: starting search list\n")
	}

	// Count occurrence of each search word over all comics, used to filter
	// out comics that do not contain all of the required search words.
	datalist := t.SearchTrie(args)

	// Add all indices to a map and count occurrence.
	m := make(map[int]int)
	for _, data := range datalist {
		temp := make(map[int]int)
		// Generate a map from all the linked btrees such that only one
		// instance of every comic index can exist per word searched.
		for _, id := range data.List {
			temp[id] = 0
		}
		// Count instances of each index so as to isolate only the
		// comics that contain all words sought.
		for id := range temp {
			m[id]++
		}
	}

	// Extract from map only those indices that contain every search term,
	// check the number of occurrence of the id against the number of search words.
	var filter []int
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

// SearchXkcd searches the local database of comic descriptions for the given
// arguments.
func (d *DataBase) SearchXkcd(args []string) {

	// Get verbal.
	if VERBOSE {
		fmt.Printf("xkcd: output start ~~~\n\n")
	}

	// Construct a search tree of all words in the database.
	t := &ds.Trie{}
	m := buildSearchMapList(d)
	t.BuildFromMap(m)

	// Convert all words in argument array to lowercase and remove spaces.
	args, ok := cleanArgs(args)
	if !ok {
		return
	}

	// Launch a search of the xkcd database.
	results := searchXkcd(t, d, args)
	d.printList(results)
	if VERBOSE {
		fmt.Printf("\nxkcd: ~~~ output end\n")
	}
}
