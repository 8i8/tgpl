// Exercise 05.17: Write a variadic function ElementsByTagName that,
// given given an HTML node tree and zero or more names, returns all the
// elements that match one of thoes names.
package main

import (
	"golang.org/x/net/html"
)

// ElementsByTagName given a html node tree, returns all of the nodes
// that match one of the given names.
func ElementsByTagName(doc *html.Node, names ...string) (nodes []*html.Node) {
	if len(names) == 0 || doc == nil {
		return nil
	}
	out := make([]*html.Node, 0)
	out = forEachNode(doc, out, seek, nil, names...)
	return out
}

// forEachNode iterates over an HTML document tree.
func forEachNode(n *html.Node, out []*html.Node,
	pre, post func(*html.Node, []*html.Node, ...string) []*html.Node,
	names ...string) []*html.Node {

	if pre != nil {
		out = pre(n, out, names...)
	}

	// Recursion if required.
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		out = forEachNode(c, out, pre, post, names...)
	}

	if post != nil {
		post(n, out, names...)
	}
	return out
}

// seek is called on arrival at a node to search its attributes.
func seek(n *html.Node, out []*html.Node, names ...string) []*html.Node {
	// Set state according to type for operations.
	for _, name := range names {
		if n.Data == name {
			out = append(out, n)
		}
	}
	return out
}

func main() {
}
