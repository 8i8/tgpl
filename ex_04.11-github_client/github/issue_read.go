package github

import "fmt"

func ReadIssue(conf Config) {

	issue, err := SearchIssues(conf)
	if err != nil {
		Log.Printf("Issue %s not found.", conf.Number)
		return
	}

	if len(issue) == 0 {
		fmt.Println("No issue returned.")
		return
	}
	printIssue(*issue[0])
}
