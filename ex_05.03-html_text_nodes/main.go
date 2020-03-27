package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
	}
	text := textNodes(nil, doc)
	for _, n := range text {
		fmt.Println(n)
	}
}

func textNodes(text []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		if len(n.Data) > 0 {
			str := strings.TrimSpace(n.Data)
			if len(str) > 1 {
				text = append(text, str)
			}
		}
	}

	if n.FirstChild != nil && n.Data != "script" && n.Data != "style" {
		text = textNodes(text, n.FirstChild)
	}
	if n.NextSibling != nil {
		text = textNodes(text, n.NextSibling)
	}
	return text
}
