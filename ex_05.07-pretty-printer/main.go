package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// main is designed such that the program can except either input via an
// UNIX type operating system pipe or as addresses given as arguments after the
// program is called.
func main() {
	out := os.Stdout
	// Stream info.
	stream := os.Stdin
	info, err := stream.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "io stream: %s", err)
		os.Exit(1)
	}
	// Check for data in the pipe.
	if info.Mode()&os.ModeNamedPipe != 0 {
		doc, err := html.Parse(stream)
		if err != nil {
			fmt.Fprintf(os.Stderr, "io stream parse: %s\n", err)
			os.Exit(1)
		}
		forEachNode(out, doc, startElement, endElement)
	}
	// If there are args, assume that they are urls.
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cmd url: %s\n", err)
			os.Exit(1)
		}
		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "cmd parse: %s\n", err)
			os.Exit(1)
		}
		forEachNode(out, doc, startElement, endElement)
	}
}

// forEachNode iterates over an HTML document tree.
func forEachNode(w io.Writer, n *html.Node, pre, post func(w io.Writer, n *html.Node)) {

	// If there are no children set short form for the elements style.
	if n.FirstChild == nil {
		m = m | shortForm
	}

	// Perform pre function.
	if pre != nil {
		pre(w, n)
	}

	// Recursion if required.
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(w, c, pre, post)
	}

	// Perform post function, unless short form mode is active.
	if post != nil && m&shortForm == 0 {
		post(w, n)
	}

	// Reset all modes to 0.
	m = norm
}

// mode keeps the state of the pretty printer.
type mode int

const (
	norm mode = 1 << iota
	script
	shortForm
)

var m mode
var depth int

// startElement is called on arrival at a node.
func startElement(w io.Writer, n *html.Node) {

	// Set state according to type for operations.
	switch n.Type {
	case html.DocumentNode:
		return
	case html.DoctypeNode:
		fmt.Fprintf(w, "<!DOCTYPE %s>\n", n.Data)
	case html.ElementNode:
		startElementNode(w, n)
	case html.TextNode:
		startTextNode(w, n)
	case html.CommentNode:
		startCommentNode(w, n)
	case html.ErrorNode:
		log.Fatalf("startElement: ErrorNode: %s", n.Data)
	}
}

// endElement is called just before leaving a node.
func endElement(w io.Writer, n *html.Node) {

	// Set state according to type for operations.
	switch n.Type {
	case html.DocumentNode:
		return
	case html.DoctypeNode:
		return
	case html.ElementNode:
		endElementNode(w, n)
	case html.TextNode:
		endTextNode(w, n)
	case html.CommentNode:
		return
	case html.ErrorNode:
		log.Fatalf("endElement: ErrorNode: %s", n.Data)
	}
}

func startElementNode(w io.Writer, n *html.Node) {
	buf := bytes.Buffer{}
	if n.Data == "script" || n.Data == "style" {
		// Dissable short form mode.
		m = m &^ shortForm
		m = m | script
	}
	for _, a := range n.Attr {
		buf.WriteString(fmt.Sprintf(" %s=\"%s\"", a.Key, a.Val))
	}
	if m&shortForm > 0 {
		fmt.Fprintf(w, "%*s<%s%s />\n", depth*2, "", n.Data, buf.String())
		return
	}
	fmt.Fprintf(w, "%*s<%s%s>\n", depth*2, "", n.Data, buf.String())
	depth++
}

func endElementNode(w io.Writer, n *html.Node) {
	if m&shortForm > 0 {
		return
	}
	depth--
	fmt.Fprintf(w, "%*s</%s>\n", depth*2, "", n.Data)
}

func startTextNode(w io.Writer, n *html.Node) {
	str := strings.TrimSpace(n.Data)
	if len(str) > 0 && str != "\n" {
		if m&script > 0 {
			fmt.Fprintf(w, "%s\n", str)
			return
		}
		fmt.Fprintf(w, "%*s%s\n", depth*2, "", str)
	}
}

func endTextNode(w io.Writer, n *html.Node) {
	if m&shortForm > 0 {
		return
	}
	depth--
}

func startCommentNode(w io.Writer, n *html.Node) {
	fmt.Fprintf(w, "%*s<!--%s-->\n", depth*2, "", n.Data)
}
