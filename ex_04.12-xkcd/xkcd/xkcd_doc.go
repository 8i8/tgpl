/*
Package xkcd - Command line client and index for the xkcd comic website.

SYNOPSIS
	xkcd [flag][value][args]

DESCRIPTION
	xkcd is a command line client and search index for the online xkcd
	comic. Once a database of the repetoir has been made by scanning the
	site, comics can be located using the clients search function or
	browsed by number.

FLAGS
	xkcd -s hello world

	-s	Search for <args> amongst the comic descriptions in the local database.

	xkcd -n 571

	-n	Display the description of comix 'n' from the database index.
	-w	Download and display the description of comic 'n' from the web.

	xkcd -u -s hello world

	-u	Update first with the latest comics descriptions.
	-v	Verbose mode, for a detailed output of the programs actions.
	-help	Prints out the programs help file.
	-h

HTTP REQUESTS

	Information of the sites API
		https://xkcd.com/json.html

	http://xkcd.com/info.0.json (current comic)
	http://xkcd.com/614/info.0.json (comic #614)

*/
/*
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
*/

// Number:     12
// Month:      1
// Link:
// News:
// SafeTitle:  Poisson
// Transcript: [[A stick figure says to another black-hat-wearing figure.]]
// Man: I'm a poisson distribution!
// Man: Still a poisson distribution.
// Hat Guy: What the hell, man.  Why do you keep saying that?
// Man: Because I'm totally a poisson distribution.
// Hat Guy: I'm less than zero.
// [[Man is gone; Hat Guy is whistling.]]
// {{alt text: Poisson distributions have no value over negative numbers}}
// Alt:        Poisson distributions have no value over negative numbers
// Img:        https://imgs.xkcd.com/comics/poisson.jpg
// Title:      Poisson
// Day:        1

package xkcd
