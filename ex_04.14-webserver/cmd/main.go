package main

import (
	"fmt"
	"log"
	"net/http"

	"tgpl/ex_04.14-webserver/itr"
)

func main() {
	// Retreive data from either the system cache or from github.
	data, err := itr.LoadCache()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	// Update cache data from github.
	data, err = itr.UpdateCache(data)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	http.HandleFunc("/", issueTracker)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func issueTracker(w http.ResponseWriter, r *http.Request) {
	itr.HtmlReport(itr.Data.Issues)
}
