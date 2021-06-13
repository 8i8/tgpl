package main

import (
	"bytes"
	"fmt"
	"testing"

	"golang.org/x/net/html"
)

func TestElementById(t *testing.T) {

	doc, err := html.Parse(bytes.NewReader([]byte(page)))
	if err != nil {
		t.Errorf("ElementById: input: html.Parse: %s", err)
	}

	n1 := ElementByID(doc, "findMe")
	n2 := `&{0xc000193730 0xc000193880 0xc000193960 0xc0001937a0 0xc0001939d0 3 object object  [{ id findMe} { type image/svg+xml} { data /assets/images/jyoti.bw.svg} { width 80} { height 80}]}`
	if fmt.Sprint(n1) == n2 {
		t.Errorf("ElementById: error: returned %v expected %v", n1, n2)
	}
}

var page = `<!DOCTYPE html>
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
