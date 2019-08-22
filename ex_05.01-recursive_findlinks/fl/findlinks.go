package fl

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  Recursive version.
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// FindlinksRec prints the links in an HTML document read from standard input.
func FindlinksRec(in io.Reader) {
	doc, err := html.Parse(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visitRec(nil, doc) {
		fmt.Println(link)
	}
}

// visitRec appends to links each link found in n and returns the result.
func visitRec(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if n.FirstChild != nil {
		links = visitRec(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visitRec(links, n.NextSibling)
	}

	return links
}
