package main

import (
	"fmt"
	"net/http"

	"tgpl/ex_04.14-webserver/itr"
)

func main() {
	err := itr.ServerStartup()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	//http.HandleFunc("/", issueTracker)
	//log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func issueTracker(w http.ResponseWriter, r *http.Request) {
}
