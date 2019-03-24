package gitish

import (
	"fmt"
	"sort"
	"time"
)

// date is a struct used to translate a time.Time object into a sortable date
// format.
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
func listIssues(conf Config, reply Reply) error {

	//fmt.Println(reply)
	r, ok := reply.Msg.(IssuesSearchResult)
	if !ok {
		return fmt.Errorf("listIssues: type assertion failed")
	}
	retIssues := r.Items

	// Prepare data structure.
	issue := make(map[int]Issue)
	index := make([]date, 0, len(retIssues))
	resp := IssueMap{&issue, &index}

	// Fill map.
	for _, item := range retIssues {
		issue[item.Number] = *item
	}

	// Check that there is a response to print.
	if resp.Len() == 0 {
		return fmt.Errorf("listIssues: empty result string")
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

	return nil
}
