package xkcd

import (
	"fmt"
)

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  Search
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// Search searches the local database of comic descriptions for the given
// arguments.
func (d *DataBase) Search(args []string) {
	if VERBOSE {
		fmt.Printf("xkcd: output start ~~~\n\n")
	}
	if BTREE {
		m := buildSearchMap(d)
		t := buildSearchTrie(m)
		results := searchBtree(t, d, args)
		printResults(d, results)
	} else if BTREESYNC {
		m := buildSearchSyncMap(d)
		t := buildSearchTrieSyncMap(m)
		results := searchBtree(t, d, args)
		printResults(d, results)
	} else if LISTSYNC {
		m := buildSearchSyncMapList(d)
		t := buildSearchTrieSyncMapList(m)
		results := searchList(t, d, args)
		printResults(d, results)
	} else {
		m := buildSearchMapList(d)
		t := buildSearchTrieList(m)
		results := searchList(t, d, args)
		printResults(d, results)
	}
	if VERBOSE {
		fmt.Printf("\nxkcd: ~~~ output end\n")
	}
}
