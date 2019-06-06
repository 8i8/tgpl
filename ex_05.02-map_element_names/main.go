package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprint(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
	// outline(nil, doc)
	m := make(map[string]int)
	mapOutline(m, doc)

	for key, value := range m {
		fmt.Printf("There are %d %v.\n", value, key)
	}
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func mapOutline(m map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		//m[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		mapOutline(m, n)
	}
}
