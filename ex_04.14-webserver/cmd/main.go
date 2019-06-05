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
	// data, err = itr.UpdateCache(data)
	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// }

	data.GenerateIndex()
	data.SortByIdDes()

	//itr.HtmlReport(os.Stdout, itr.Data)

	http.HandleFunc("/", issueList)
	http.HandleFunc("/users", userList)
	http.HandleFunc("/milestones", milestoneList)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func issueList(w http.ResponseWriter, r *http.Request) {
	itr.HtmlIssueReport(w, itr.Data)
}

func userList(w http.ResponseWriter, r *http.Request) {
	itr.HtmlUserReport(w, itr.Data)
}

func milestoneList(w http.ResponseWriter, r *http.Request) {
	itr.HtmlMilestoneReport(w, itr.Data)
}
