package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// mode is used to pass the required state and action between recursive
// function calls.
type mode int

const (
	norm mode = iota
	script
	style
)

type links struct {
	m      mode
	link   []string
	text   []string
	img    []string
	script []string
	style  []string
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %s\n", err)
		os.Exit(1)
	}
	l := links{}
	l = visit(l, doc)
	printAll(l)
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
				l.text = append(l.text, str)
			case script:
				l.script = append(l.script, str)
			case style:
				l.style = append(l.style, str)
			}
		}
	}
	l.m = norm
	return l
}

func visitElementNode(l links, n *html.Node) links {

	l.m = norm
	switch n.Data {
	case "a":
		for _, a := range n.Attr {
			if a.Key == "href" {
				l.link = append(l.link, a.Val)
			}
		}
	case "img":
		for _, a := range n.Attr {
			if a.Key == "src" {
				l.img = append(l.img, a.Val)
			}
		}
	case "script":
		for _, a := range n.Attr {
			if a.Key == "src" {
				l.script = append(l.script, a.Val)
			}
		}
		l.m = script
	case "link":
		var ok bool
		for _, a := range n.Attr {
			if a.Key == "rel" && a.Val == "stylesheet" {
				ok = true
			}
		}
		if ok {
			for _, a := range n.Attr {
				if a.Key == "href" {
					l.style = append(l.style, a.Val)
				}
			}
			ok = false
		}
	case "style":
		l.m = style
	}
	return l
}

func printAll(l links) {

	buf := bytes.Buffer{}
	buf.WriteString("Text\n----\n")
	for _, o := range l.text {
		buf.WriteString(o)
		buf.WriteString("\n")
	}

	buf.WriteString("\nLink\n----\n")
	for _, o := range l.link {
		buf.WriteString(o)
		buf.WriteByte('\n')
	}

	buf.WriteString("\nImg\n---\n")
	for _, o := range l.img {
		buf.WriteString(o)
		buf.WriteByte('\n')
	}

	buf.WriteString("\nScript\n------\n")
	for _, o := range l.script {
		buf.WriteString(o)
		buf.WriteByte('\n')
	}

	buf.WriteString("\nStyle\n-----\n")
	for _, o := range l.style {
		buf.WriteString(o)
		buf.WriteByte('\n')
	}

	fmt.Println(buf.String())
}
