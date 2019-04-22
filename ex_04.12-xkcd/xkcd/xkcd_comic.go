package xkcd

import (
	"fmt"
	"strconv"
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
	URL        string
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

// setURL auto generates the comics web url.
func setURL(i uint) string {
	return "https://xkcd.com/" + strconv.Itoa(int(i)) + "/"
}

// setURLOnAllcomics does just that, setting the web url of the comic on each
// database entry.
func setURLOnAll(comics *DataBase) error {

	for i, comic := range comics.Edition {
		comics.Edition[i].URL = setURL(comic.Number)
	}
	err := writeDatabase(comics)
	if err != nil {
		return fmt.Errorf("setURLOnAll: %v", err)
	}

	if VERBOSE {
		fmt.Printf("xkcd: url set on database.\n")
	}
	return nil
}
