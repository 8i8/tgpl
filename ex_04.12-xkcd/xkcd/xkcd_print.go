package xkcd

import (
	"fmt"
)

func (d *DataBase) printList(list []int) {
	for _, i := range list {
		if TITLE {
			printTitle(d.Edition[i])
		} else {
			printSingle(d.Edition[i])
		}
	}
}

func (d *DataBase) printComic(n int) {
	if TITLE {
		printTitle(d.Edition[n-1])
	} else {
		printSingle(d.Edition[n-1])
	}
}

func printSingle(c Comic) {
	fmt.Printf("~~~~~~\n")
	fmt.Printf("xkcd: %d: %v ~ %v\n", c.Number, c.SafeTitle, c.URL)
	fmt.Printf("~~~\n")
	fmt.Printf("Image: %s\n", c.Img)
	fmt.Printf("~~~\n")
	fmt.Printf("Alt text:\n%s\n", c.Alt)
	fmt.Printf("~~~\n")
	fmt.Printf("Transcript: %s\n", c.Transcript)
	fmt.Printf("~~~~~~\n")
}

func printTitle(c Comic) {
	fmt.Printf("xkcd: %4.d: %v ~ %v\n", c.Number, c.SafeTitle, c.URL)
}

func printEverything(c Comic) {
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
	fmt.Printf("URL:        %s\n", c.URL)
}
