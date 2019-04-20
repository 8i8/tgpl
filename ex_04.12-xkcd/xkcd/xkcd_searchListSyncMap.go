package xkcd

import (
	"8i8/ds"
	"fmt"
	"sync"
)

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *   list
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// searchList prepares a list of comics that contain the given search terms.
func searchListSyncMap(t *ds.Trie, comics *DataBase, args []string) []uint {

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
func scanComicSyncMapList(m *sync.Map, c Comic) *sync.Map {

	m = ds.ExtractAndSyncMapList(m, c.Link, c.Number)
	m = ds.ExtractAndSyncMapList(m, c.News, c.Number)
	m = ds.ExtractAndSyncMapList(m, c.SafeTitle, c.Number)
	m = ds.ExtractAndSyncMapList(m, c.Transcript, c.Number)
	m = ds.ExtractAndSyncMapList(m, c.Alt, c.Number)
	m = ds.ExtractAndSyncMapList(m, c.Title, c.Number)

	return m
}

// buildSearchSyncMapList scans the comic database and creates a map of all words
// found, linking them to the comics that they are from.
func buildSearchSyncMapList(comics *DataBase) *sync.Map {

	// Scan and map comics.
	m := new(sync.Map)

	for _, comic := range comics.Edition {
		scanComicSyncMapList(m, comic)
	}

	return m
}

// buildSearchTrie constructs a search trie from a search map.
func buildSearchTrieSyncMapList(m *sync.Map) *ds.Trie {

	t := new(ds.Trie)
	m.Range(func(word, list interface{}) bool {
		t.AddList(word.(string), list.([]uint))
		return true
	})

	return t
}
