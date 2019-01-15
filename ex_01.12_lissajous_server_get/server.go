package main

import (
	"../ex_01.12_lissajous_server/lissa"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	lissa.SeedRand()
	http.HandleFunc("/", lissaHandler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/test", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
	mu.Lock()
	count++
	mu.Unlock()
}

func lissaHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	strMap := r.URL.Query()
	if len(strMap) > 0 {
		num, _ := strconv.Atoi(strMap["cycles"][0])
		lissa.Lissajous(w, num)
	} else {
		lissa.Lissajous(w, 0)
	}
}
