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
		ElementByID(doc, "findMe")
		goto exit
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
		ElementByID(doc, "findMe")
		goto exit
	}
	if len(os.Args) == 1 {
		doc, err := html.Parse(bytes.NewReader([]byte(testPage)))
		if err != nil {
			fmt.Fprintf(os.Stdout, "testPage: input: html.Parse: %s", err)
			os.Exit(1)
		}
		// An eliment with the id findMe has been created in testPage.
		id := "findMe"
		ElementByID(doc, id)
	}
exit:
}

var found bool

// ElementByID returns the first node that it reaches with the given id.
func ElementByID(doc *html.Node, id string) *html.Node {
	n := forEachNode(doc, seek, nil, id)
	switch {
	case n != nil:
		fmt.Printf("An eliment with the id %q was found\n", id)
	case n == nil:
		fmt.Printf("No eliment with the id %q could be found\n", id)
	}
	return n
}

// forEachNode iterates over an HTML document tree.
func forEachNode(n *html.Node,
	pre, post func(n *html.Node, id string) bool, id string) *html.Node {

	if pre != nil {
		if !found {
			found = pre(n, id)
		}
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

var testPage = `<!DOCTYPE html>
<html lang="en-gb">
  <head>
    <meta charset="utf-8" />
    <meta name="robots" content="noindex, nofollow" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>
      home
    </title>
    <link rel="preload" href="/assets/javascript/scripts.js" />
    <link rel="preload" href="/assets/css/main.css" />
    <link rel="stylesheet" href="/assets/css/main.css" />
    <script src="/assets/javascript/scripts.js">
    </script>
  </head>
  <body>
    <!-- This is a comment to test that comments are also resotred. -->
    <script>
      beforeLoad();
    </script>
    <header>
      <a href="/fractal">
        <object type="image/svg+xml" data="/assets/images/Om.svg" width="80" height="80">
          <img src="/assets/images/Om.svg" alt="AUM" width="80" height="80" />
        </object>
      </a>
      <h1>
        Jyoḍi
      </h1>
    </header>
    <div class="all-site-wrapper">
      <nav class="black-grad">
        <ul>
          <li>
            <a href="/nav/jyoti.html">
              <object id="findMe" type="image/svg+xml" data="/assets/images/jyoti.bw.svg" width="80" height="80">
                <img src="/assets/images/jyoti.bw.svg.png" alt="What on earth" />
              </object>
            </a>
          </li>
          <li>
            <a href="/nav/nakshatra.html">
              <object type="image/svg+xml" data="/assets/images/nakshatra.bw.svg" width="80" height="80">
                <img src="/assets/images/nakshatra.bw.svg.png" alt="nakshatra" />
              </object>
            </a>
          </li>
          <li>
            <a href="/nav/graha.html">
              <object type="image/svg+xml" data="/assets/images/graha.bw.svg" width="80" height="80">
                <img src="/assets/images/graha.bw.svg.png" alt="graha" />
              </object>
            </a>
          </li>
          <li>
            <a href="/nav/rasi.html">
              <object type="image/svg+xml" data="/assets/images/rasi.bw.svg" width="80" height="80">
                <img src="/assets/images/rasi.bw.svg.png" alt="rasi" />
              </object>
            </a>
          </li>
          <li>
            <a href="/nav/bhava.html">
              <object type="image/svg+xml" data="/assets/images/bhava.bw.svg" width="80" height="80">
                <img src="/assets/images/bhava.bw.svg.png" alt="bhava" />
              </object>
            </a>
          </li>
          <li>
            <a href="/nav/books.html">
              <object type="image/svg+xml" data="/assets/images/books.bw.svg" width="80" height="80">
                <img src="/assets/images/books.bw.svg.png" alt="info" />
              </object>
            </a>
          </li>
        </ul>
      </nav>
      <div class="index">
      </div>
      <article>
      </article>
      <aside>
      </aside>
    </div>
    <footer>
      Page footer
    </footer>
  </body>
</html>
`
