package main

import (
	"../ex_03.04-svg_server/orchid"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		orchid.Orchid(w)
	})
	http.HandleFunc("/hello", homepage)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "Hello World")
}
