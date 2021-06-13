// Exercise 05.18: Without changing its behaviour, rewrite the fetch
// function to use defer to close the writable file.
package main

import (
	"io"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func(error) error {
		// Close file, but prefer error from Copy, if any.
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
		return err
	}(err)

	n, err = io.Copy(f, resp.Body)

	return local, n, err
}

func main() {
}
