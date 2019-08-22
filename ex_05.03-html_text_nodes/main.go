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
	}
	textnodes(nil, doc)
}

func textnodes(stack [], n *html.Node) {
	depth := 0
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return z.Err()
		case html.TextToken:
			if depth > 0 {
				// emitBytes should copy the []byte it receives,
				// if it doesn't process it immediately.
				emitBytes(z.Text())
			}
		case html.StartTagToken, html.EndTagToken:
			tn, _ := z.TagName()
			if len(tn) == 1 && tn[0] == 'a' {
				if tt == html.StartTagToken {
					depth++
				} else {
					depth--
				}
			}
		}
	}
}
