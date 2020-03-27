package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "CountWordsAndImages failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("URL %s returned %d words and %d images.\n", url, words, images)
	}
}

// CountWordsAndImages counts the words and images found at the given url.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	l := links{}
	l = visit(l, n)
	words = l.wCount
	images = l.imgCount
	return
}

// mode is used to pass the required state and action between recursive
// function calls.
type mode int

const (
	norm mode = iota
	script
	style
)

type links struct {
	m        mode
	imgCount int
	wCount   int
	link     []string
	text     []string
	img      []string
	script   []string
	style    []string
}

func visit(l links, n *html.Node) links {

	switch n.Type {
	case html.TextNode:
		l = visitTextNode(l, n)
	case html.ElementNode:
		l = visitElementNode(l, n)
	}

	// If we have just visited an element node and have changed mode, the
	// consequent mode set and then passed into the recursive call so that the node
	// the node content can be dealt with accordingly.
	if n.FirstChild != nil {
		l = visit(l, n.FirstChild)
	}
	if n.NextSibling != nil {
		l = visit(l, n.NextSibling)
	}
	return l
}

func visitTextNode(l links, n *html.Node) links {
	if len(n.Data) > 0 {
		str := strings.TrimSpace(n.Data)
		if len(str) > 1 {
			switch l.m {
			case norm:
				l.wCount += len(strings.Split(str, " "))
			}
		}
	}
	l.m = norm
	return l
}

func visitElementNode(l links, n *html.Node) links {

	l.m = norm
	switch n.Data {
	case "img":
		l.imgCount++
	case "script":
		l.m = script
	case "style":
		l.m = style
	}
	return l
}
