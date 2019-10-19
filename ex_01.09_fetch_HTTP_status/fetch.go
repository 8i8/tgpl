package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const prefix = "http://"

func main() {
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(strings.ToLower(url), prefix) == false {
			url = prefix + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			fmt.Fprintf(os.Stdout, "fetch: %v\n", resp.Status)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		fmt.Fprintf(os.Stdout, "fetch: %v\n", resp.Status)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			fmt.Fprintf(os.Stdout, "fetch: %v\n", resp.Status)
			os.Exit(1)
		}
	}
}
