package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// main is designed such that the program can except either input via an
// UNIX type operating system pipe or as addresses given as arguments after the
// program is called.
func main() {
	// Stream info.
	stream := os.Stdin
	info, err := stream.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "os.Stdin: in.Stat: %v", err)
		os.Exit(1)
	}
	// Check for data in the pipe.
	if info.Mode()&os.ModeNamedPipe != 0 {
		doc, err := html.Parse(stream)
		if err != nil {
			fmt.Fprintf(os.Stderr, "stream: html.Parse: %s\n", err)
			os.Exit(1)
		}
		n := ElementByID(doc, "findMe")
		fmt.Println(n)
	}
	// If there are args, assume that they are urls.
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			return
		}
		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			err = fmt.Errorf("args: html.Parse: %s", err)
			return
		}
		n := ElementByID(doc, "findMe")
		fmt.Println(n)
	}
	if len(os.Args) == 1 {
		doc, err := html.Parse(bytes.NewReader([]byte(testPage)))
		if err != nil {
			fmt.Fprintf(os.Stdout, "testPage: input: html.Parse: %s", err)
			os.Exit(1)
		}
		// An eliment with the id findMe has been created in testPage.
		n := ElementByID(doc, "findMe")
		fmt.Println(n)
	}
}

var found bool

// ElementByID returns the first node that it reaches with the given id.
func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, seek, nil, id)
}

// forEachNode iterates over an HTML document tree.
func forEachNode(n *html.Node,
	pre, post func(n *html.Node, id string) bool, id string) *html.Node {

	if pre != nil {
		found = pre(n, id)
		if found {
			return n
		}
	}

	// Recursion if required.
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		n = forEachNode(c, pre, post, id)
		if found {
			return n
		}
	}

	if post != nil {
		found = post(n, id)
		if found {
			return n
		}
	}
	return nil
}

// seek is called on arrival at a node to search its attributes.
func seek(n *html.Node, id string) bool {
	// Set state according to type for operations.
	switch n.Type {
	case html.DocumentNode:
	case html.DoctypeNode:
	case html.ElementNode:
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return true
			}
		}
	case html.TextNode:
	case html.CommentNode:
	case html.ErrorNode:
		log.Fatalf("startElement: ErrorNode: %s", n.Data)
	}
	return false
}
