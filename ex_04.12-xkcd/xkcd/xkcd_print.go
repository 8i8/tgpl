package xkcd

import (
	"8i8/ds"
)

// type Comic struct {
// 	Month      string
// 	Number     uint `json:"num"`
// 	Link       string
// 	Year       string
// 	News       string
// 	SafeTitle  string `json:"safe_title"`
// 	Transcript string
// 	Alt        string
// 	Img        string
// 	Title      string
// 	Day        string
// }

func printMap(m *ds.Trie) {
}

func printResults(c *DataBase, r []uint) {
	for _, i := range r {
		// -1 offset due to array of comics starting at 0.
		printTitle(c.Edition[i-1])
	}
}
