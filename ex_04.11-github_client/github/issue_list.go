package github

import (
	"fmt"
	"sort"
	"time"
)

type date struct {
	r int
	y int
	m time.Month
	d int
	p bool
}

// Struct for sorting issue map with an index.
type IssueMap struct {
	M *map[int]Issue
	I *[]date
}

// Length of issue map.
func (v IssueMap) Len() int {
	return len(*v.M)
}

// listIssues Retrieves a list of issues from the given repo that meet the
// search criteria.
func ListIssues(conf Config) {

	// Retrieve data
	result, err := searchIssues(conf)
	if err != nil {
		fmt.Printf("no issues returned.")
		return
	}

	// Prepare data structure.
	issue := make(map[int]Issue)
	index := make([]date, 0, len(result))
	resp := IssueMap{&issue, &index}

	// Fill map.
	for _, item := range result {
		issue[item.Number] = *item
	}

	// Check that there is a responst to print.
	if resp.Len() == 0 {
		fmt.Println("Empty result string.")
		return
	}

	// Construct index.
	for r, item := range issue {
		y, m, d := item.CreatedAt.Date()
		p := false
		issue := date{r, y, m, d, p}
		index = append(index, issue)
	}

	// Sort descending, newest first.
	sort.Slice(index, func(i, j int) bool {
		return index[i].r > index[j].r
	})

	// If there is something to print, print it.
	printIssues(resp)

	return
}
