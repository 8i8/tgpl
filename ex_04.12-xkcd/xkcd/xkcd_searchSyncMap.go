package xkcd

import (
	"8i8/ds"
	"fmt"
	"sync"
)

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *   btree sync map
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// searchBtree prepares a list of comics that contain the given search terms.
func searchBtreeSyncMap(t *ds.Trie, comics *DataBase, args []string) []uint {

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

// goScanComicSyncMap runs extract words on every text field in a Comic struct.
func goScanComicSyncMap(m *sync.Map, c Comic, ch chan<- *sync.Map) {

	ch1 := make(chan *sync.Map)
	go ds.GoExtractAndSyncMap(m, c.Link, c.Number, ch1)
	go ds.GoExtractAndSyncMap(m, c.News, c.Number, ch1)
	go ds.GoExtractAndSyncMap(m, c.SafeTitle, c.Number, ch1)
	go ds.GoExtractAndSyncMap(m, c.Transcript, c.Number, ch1)
	go ds.GoExtractAndSyncMap(m, c.Alt, c.Number, ch1)
	go ds.GoExtractAndSyncMap(m, c.Title, c.Number, ch1)

	for i := 0; i < 6; i++ {
		m = <-ch1
	}

	ch <- m
}

// scanComicMap runs extract words on every text field in a Comic struct.
func scanComicSyncMap(m *sync.Map, c Comic) *sync.Map {

	m = ds.ExtractAndSyncMap(m, c.Link, c.Number)
	m = ds.ExtractAndSyncMap(m, c.News, c.Number)
	m = ds.ExtractAndSyncMap(m, c.SafeTitle, c.Number)
	m = ds.ExtractAndSyncMap(m, c.Transcript, c.Number)
	m = ds.ExtractAndSyncMap(m, c.Alt, c.Number)
	m = ds.ExtractAndSyncMap(m, c.Title, c.Number)

	return m
}

// buildSearchMap scans the comic database and creates a map of all words
// found, linking them to the comics that they are from.
func buildSearchSyncMap(comics *DataBase) *sync.Map {

	// Scan and map comics.
	m := new(sync.Map)

	for _, comic := range comics.Edition {
		m = scanComicSyncMap(m, comic)
	}

	return m
}

// buildSearchMap scans the comic database and creates a map of all words
// found, linking them to the comics that they are from.
func buildSearchComicSyncMap(comics *DataBase) *sync.Map {

	// Scan and map comics.
	m := new(sync.Map)
	ch := make(chan *sync.Map)

	for _, comic := range comics.Edition {
		go goScanComicSyncMap(m, comic, ch)
	}

	for range comics.Edition {
		m = <-ch
	}

	return m
}

// buildSearchTrieSyncMap constructs a search trie from a search map.
func buildSearchTrieSyncMap(m *sync.Map) *ds.Trie {

	t := new(ds.Trie)
	m.Range(func(word, btree interface{}) bool {
		t.AddBtree(word.(string), btree.(*ds.BtreeNode))
		return true
	})

	return t
}
