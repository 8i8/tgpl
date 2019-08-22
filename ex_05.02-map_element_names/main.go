package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprint(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
	// The initial code, uncomment to use.
	// outline(nil, doc)

	m := make(map[string]int)

	// First attempt at a solution, fails due to stack smashing.
	// mapOutline(m, doc)

	// Solution using Tokenizer rather than Parse.
	// mapOutlineTokens(m, os.Stdin)

	// Solution using a localy defined function recursivly.
	m := mapOutline2(os.Stdin)

	// Required for the map solutions.
	for key, value := range m {
		if len(key) != 0 {
			fmt.Printf("%10v%3d\n", key, value)
		}
	}
}

// outline uses the []string to print data and then after every recursive call
// the buffer returns to its previous state, thus there is no overflow of the
// stack.
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

// mapOutline stores every singe incidence of each tag into the map and its
// recursion causes a buffer overflow.
func mapOutline(m map[string]int, n *html.Node) {

	if n.Type == html.ElementNode {
		m[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		mapOutline(m, n)
	}
}

// mapOutline2 avoids overflow by using the recursion of a localy defined
// function.
func mapOutline2(r io.Reader) map[string]int {
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprint(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
	m := make(map[string]int)
	var f func(*html.Node, map[string]int)
	f = func(n *html.Node, m map[string]int) {
		if n.Type == html.ElementNode {
			m[n.Data]++
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c, m)
		}
	}
	f(doc, m)
	return m
}

func mapOutlineTokens(m map[string]int, n io.Reader) {
	z := html.NewTokenizer(n)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			return
		}
		str, _ := z.TagName()
		m[string(str)]++
	}
}
