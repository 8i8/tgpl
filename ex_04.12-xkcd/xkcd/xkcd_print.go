package xkcd

import "tgpl/ex_04.12-xkcd/ds"

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
	//fmt.Println(r)
	for _, i := range r {
		// -1 offset due to array of comics starting at 0.
		printTitle(c.Edition[i-1])
	}
}
