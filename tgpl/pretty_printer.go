package tgpl

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

const (
	strHttp  = "http://"
	strHttps = "https://"
)

func RemoveHttpPrefix(url string) string {
	switch {
	case strings.HasPrefix(strings.ToLower(url), strHttp) == true:
		url = url[len(strHttp):]
	case strings.HasPrefix(strings.ToLower(url), strHttps) == true:
		url = url[len(strHttps):]
	}
	return url
}

func CheckPrefix(url, prefix string) string {
	if strings.HasPrefix(strings.ToLower(url), strHttp) == false &&
		strings.HasPrefix(strings.ToLower(url), strHttps) == false {
		url = prefix + url
	}
	return url
}

func PrettyPrint(w io.Writer, url string) error {
	fname := "PrettyPrint"
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf(fname+"getting: %s\n", err)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf(fname+"parsing: %s\n", err)
	}
	forEachNodeWrite(w, doc, startElement, endElement)
	return nil
}

// forEachNode iterates over an HTML document tree.
func forEachNodeWrite(w io.Writer, n *html.Node, pre, post func(w io.Writer, n *html.Node)) {

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
		forEachNodeWrite(w, c, pre, post)
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

func attribute(a html.Attribute) html.Attribute {
	switch a.Key {
	case "href", "src", "data":
		if a.Val != "" || a.Val[0] == '/' {
			a.Val = "." + a.Val
		}
	}
	return a
}

func setPathBase(n *html.Node) {
	switch n.Data {
	case "a", "link", "img", "object":
		for i, a := range n.Attr {
			n.Attr[i] = attribute(a)
		}
	}
}

func startElementNode(w io.Writer, n *html.Node) {
	buf := bytes.Buffer{}
	if n.Data == "script" || n.Data == "style" {
		// Dissable short form mode.
		m = m &^ shortForm
		m = m | script
	}

	setPathBase(n)

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
