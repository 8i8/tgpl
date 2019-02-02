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

// Struct to link issues to there index.
type IssueMap struct {
	M *map[int]Issue
	I *[]date
}

// listIssues Retrieves a list of issues from the given repo that meet the
// search criteria.
func ListIssues(conf Config) (IssueMap, error) {

	issue := make(map[int]Issue)
	index := make([]date, 0, len(issue))
	resp := IssueMap{&issue, &index}

	// Retrieve data
	result, err := SearchIssues(conf)
	if err != nil {
		Log.Printf("error: %v", err.Error())
		return resp, err
	}

	// Make and fill a map from the result array.
	for _, item := range result.Items {
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

	return resp, err
}

// Print to terminal order by date, separating newer than one month and within
// a year.
func PrintIssues(results IssueMap) {

	issue := *results.M
	index := *results.I
	now := date{}
	now.y, now.m, now.d = time.Now().Date()
	fmt.Printf("%d issues:\n", len(index))

	// Print issues that are less than a month old.
	fmt.Println("\nLess than a month")
	for n, d := range index {
		item := issue[d.r]
		if lessThanMonth(now, d) {
			printLine(item, d)
			index[n].p = true
		}
	}

	// Print issues that are less than a year old.
	fmt.Println("\nLess than a year")
	for n, d := range index {
		item := issue[d.r]
		if monthToAYear(now, d) {
			printLine(item, d)
			index[n].p = true
		}
	}

	// Print issues that are older than a year.
	fmt.Println("\nOlder than a year")
	for n, d := range index {
		item := issue[d.r]
		if yearOnward(now, d) {
			printLine(item, d)
			index[n].p = true
		}
	}
}

func printLine(item Issue, d date) {
	fmt.Printf("#%-5d %9.9s %55.55s  %.2d %s %d\n",
		item.Number, item.User.Login, item.Title, d.d, d.m, d.y)
}
