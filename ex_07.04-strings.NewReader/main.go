// Exercise 07.04: The strings.NewReader function returns a value that
// satisfies the io.Reader interface (and others) by reading from its
// arguments, a string. Implement a simple version of NewReader
// yourself, and use it to make the HTML parser (5.2) take input from a
// string.
package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

// Reader stores a string that is to be read along with the index of the
// next bytes to be read.
type Reader struct {
	s string
	i int // index
}

// NewReader implements the Read interface upon a string.
func NewReader(s string) *Reader {
	return &Reader{s, 0}
}

// Read implements the Read method upon the Reader struct.
func (r *Reader) Read(b []byte) (n int, err error) {
	if r.i >= len(r.s) {
		return 0, io.EOF
	}
	n = copy(b, r.s[r.i:])
	r.i += n
	return
}

func main() {
	r := NewReader(testPage)
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	m := make(map[string]int)
	outline(m, doc)
	for key, value := range m {
		fmt.Println(key, value)
	}
}

func outline(m map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		m[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(m, c)
	}
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
