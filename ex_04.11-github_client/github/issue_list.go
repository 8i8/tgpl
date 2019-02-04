package github

import (
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

// Struct to link issues to there index.
type IssueMap struct {
	M *map[int]Issue
	I *[]date
}

func (v IssueMap) Len() int {
	return len(*v.M)
}

// listIssues Retrieves a list of issues from the given repo that meet the
// search criteria.
func ListIssues(conf Config) error {

	// Retrieve data
	result, err := SearchIssues(conf)
	if err != nil {
		Log.Printf("no issue returned.")
		return err
	}

	issue := make(map[int]Issue)
	index := make([]date, 0, len(result))
	resp := IssueMap{&issue, &index}

	// Make and fill a map from the result array.
	for _, item := range result {
		issue[item.Number] = *item
	}

	// Make and fill an array to index the map.
	for r, item := range issue {
		y, m, d := item.CreatedAt.Date()
		p := false
		issue := date{r, y, m, d, p}
		index = append(index, issue)
	}

	// Sort the results, order by reference numbers.
	sort.Slice(index, func(i, j int) bool {
		return index[i].r > index[j].r
	})

	// If there is something to print, print it.
	PrintIssues(resp)

	return err
}
