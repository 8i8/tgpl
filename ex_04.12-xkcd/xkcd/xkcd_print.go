package xkcd

import (
	"fmt"
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

func printComic(i uint) {
	printSingle(comics.Edition[i])
}

func printSingle(c Comic) {
	fmt.Printf("Number:     %d\n", c.Number)
	fmt.Printf("Month:      %s\n", c.Month)
	fmt.Printf("Link:       %s\n", c.Link)
	fmt.Printf("News:	%s\n", c.News)
	fmt.Printf("SafeTitle:  %s\n", c.SafeTitle)
	fmt.Printf("Transcript: %s\n", c.Transcript)
	fmt.Printf("Alt:        %s\n", c.Alt)
	fmt.Printf("Img:        %s\n", c.Img)
	fmt.Printf("Title:      %s\n", c.Title)
	fmt.Printf("Day:        %s\n", c.Day)
}

func printTitle(c Comic) {
	fmt.Printf("xkcd: %d %v\n", c.Number, c.SafeTitle)
}

func printMap(m mSearch) {
	for word, _ := range m {
		fmt.Println(word)
	}
}

func printResults(c Comics, r []uint) {
	for _, i := range r {
		printTitle(c.Edition[i])
	}
}
