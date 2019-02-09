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

// IssueMap is a struct for maintaining an issue map with its index.
type IssueMap struct {
	M *map[int]Issue
	I *[]date
}

// Len outputs the length of an issue map.
func (v IssueMap) Len() int {
	return len(*v.M)
}

// ListIssues retrieves a list of issues from the given repo that meet the;
// search criteria.
func ListIssues(conf Config) error {

	var err error
	// Retrieve data
	result, err := searchIssues(conf)
	if err != nil {
		return fmt.Errorf("no issues returned")
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
		return fmt.Errorf("empty result string")
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

	return err
}
