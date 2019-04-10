package xkcd

import (
	"fmt"
)

// Comic contains an xkcd cartoon.
type Comic struct {
	Month      string
	Number     uint `json:"num"`
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

// Num returns the comic number.
func (c Comic) Num() uint {
	return c.Number
}

// WebGet outputs the given comic from the internet by making a http request.
func WebGet(n uint) {

	if VERBOSE {
		fmt.Printf("xkcd: make http request\n")
	}

	comic, code, err := newComicHTTP(n)
	if err != nil && code != 404 {
		fmt.Printf("error: newComicHTTP: %v\n", err)
	} else if code == 404 {
		fmt.Printf("xkcd: http: 404 comic not found\n")
	}

	if VERBOSE {
		fmt.Printf("xkcd: http request closed\n")
	}

	if err == nil {
		printSingle(comic)
	}
}

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  Print
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

func printComic(c *DataBase, i uint) {
	printSingle(c.Edition[i])
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
