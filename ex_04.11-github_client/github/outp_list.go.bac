package gitish

import (
	"fmt"
	"sort"
	"time"
)

// date is used to translate a time.Time object into a calculable format.
type date struct {
	i int
	y int
	m time.Month
	d int
	p bool
}

// IssueMap stores issues sorted by date.
type IssueMap struct {
	M         *map[int]Issue
	issuelist []date
	Imonth    []date
	Iyear     []date
	Imore     []date
}

// Len outputs the length of an issue map.
func (v IssueMap) Len() int {
	return len(*v.M)
}

// ListIssues retrieves a list of issues from the given repo that meet the
// search criteria.
func listIssues(reply Reply) (IssueMap, error) {

	resp := IssueMap{}
	// Approprate type for the reply.
	r, ok := reply.Msg.(IssuesSearchResult)
	if !ok {
		return resp, fmt.Errorf("listIssues: type assertion failed")
	}
	retIssues := r.Items

	// Prepare data structures.
	issues := make(map[int]Issue)
	resp.M = &issues
	resp.issuelist = make([]date, 0, len(retIssues))

	// Fill map with links to the corresponding issues.
	for _, issue := range retIssues {
		issues[issue.Number] = *issue
	}

	// Check that there is data to print.
	if resp.Len() == 0 {
		return resp, fmt.Errorf("listIssues: empty result string")
	}

	// Retrieve the current date.
	now := date{}
	now.y, now.m, now.d = time.Now().Date()

	// Var for counting how many or each time designation exist.
	a, b, c := 0, 0, 0
	// Construct index.
	for i, issue := range issues {

		// Flag for sorting date
		var c cal

		y, m, d := issue.CreatedAt.Date()
		da := date{i, y, m, d, false}

		// Add to issue list.
		resp.issuelist = append(resp.issuelist, da)

		// Count how many dates fall into each time bracket.
		c |= dateSort(now, da)
		switch {
		case c&dMONTH > 0:
			a++
		case c&dYEAR > 0:
			b++
		case c&dMORE > 0:
			c++
		}

	}

	// Sort descending most recent first.
	sort.Slice(resp.issuelist, func(i, j int) bool {
		return resp.issuelist[i].i > resp.issuelist[j].i
	})

	// Shunt each count, b and c, to the end of each consecutive segment
	// for ease of reading.
	b = b + a
	c = c + b

	// Group into slices: Under a month; Between month and a year; Over a year.
	resp.Imonth = resp.issuelist[:a]
	resp.Iyear = resp.issuelist[a:b]
	resp.Imore = resp.issuelist[c:]

	return resp, nil
}

// // ListIssues retrieves a list of issues from the given repo that meet the
// // search criteria.
// func listIssues(reply Reply) error {

// 	// Approprate type
// 	r, ok := reply.Msg.(IssuesSearchResult)
// 	if !ok {
// 		return fmt.Errorf("listIssues: type assertion failed")
// 	}
// 	retIssues := r.Items

// 	// Prepare data structures.
// 	issue := make(map[int]Issue)
// 	l := len(retIssues) / 3
// 	imonth := make([]date, 0, l)
// 	iyear := make([]date, 0, l)
// 	imore := make([]date, 0, l)
// 	resp := IssueMap{&issue, &imonth, &iyear, &imore}

// 	// Fill map.
// 	for _, item := range retIssues {
// 		issue[item.Number] = *item
// 	}

// 	// Check that there is data to print.
// 	if resp.Len() == 0 {
// 		return fmt.Errorf("listIssues: empty result string")
// 	}

// 	now := date{}
// 	now.y, now.m, now.d = time.Now().Date()

// 	// Construct index.
// 	for r, item := range issue {

// 		// Flag for sorting date
// 		var c cal

// 		y, m, d := item.CreatedAt.Date()
// 		issue := date{r, y, m, d, false}
// 		c |= dateSort(now, issue)
// 		switch {
// 		case c&dMORE > 0:
// 			imore = append(imore, issue)
// 		case c&dYEAR > 0:
// 			iyear = append(iyear, issue)
// 		case c&dMONTH > 0:
// 			imonth = append(imonth, issue)
// 		}
// 	}

// 	// Sort descending, newest first and then print.
// 	sort.Slice(imonth, func(i, j int) bool {
// 		return imonth[i].r > imonth[j].r
// 	})
// 	sort.Slice(iyear, func(i, j int) bool {
// 		return iyear[i].r > iyear[j].r
// 	})
// 	sort.Slice(imore, func(i, j int) bool {
// 		return imore[i].r > imore[j].r
// 	})

// 	printIssues(resp)

// 	return nil
// }
