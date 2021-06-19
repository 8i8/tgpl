package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const prefix = "http://"
const prefixs = "https://"

func main() {
	for _, u := range os.Args[1:] {
		if strings.HasPrefix(strings.ToLower(u), prefix) == false &&
			strings.HasPrefix(strings.ToLower(u), prefixs) == false {
			u = prefix + u
		}

		resp, err := http.Get(u)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", u, err)
			os.Exit(1)
		}
	}
}
