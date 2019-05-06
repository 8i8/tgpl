package dates

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"tgpl/ex_04.14-github.text.template/github"
)

// IssueMap stores issues sorted by date.
type IssueMap struct {
	M         *map[int]github.Issue
	issuelist []date
	Imonth    []date
	Iyear     []date
	Imore     []date
}

// Len outputs the length of an issue map.
func (v IssueMap) Len() int {
	return len(*v.M)
}

// listIssues retrieves a list of issues from the given repo that meet the
// search criteria.
func listIssues(r *github.IssuesSearchResult) (IssueMap, error) {

	resp := IssueMap{}
	retIssues := r.Items

	// Prepare data structures.
	issues := make(map[int]github.Issue)
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
		c |= DateSort(now, da)
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

// ListIssues sorts by date and then displays a list of issues in the terminal.
func ListIssues(reply *github.IssuesSearchResult) error {
	resp, err := listIssues(reply)
	if err != nil {
		log.Fatal(err)
	}
	printIssues(resp)
	return nil
}

// Print a list of single line issues, grouped by date.
func printIssues(Issues IssueMap) {

	issue := *Issues.M
	imonth := Issues.Imonth
	iyear := Issues.Iyear
	imore := Issues.Imore
	now := date{}
	now.y, now.m, now.d = time.Now().Date()
	fmt.Printf("%d issues:\n", len(issue))

	// Print issues that are less than a month old.
	printList(imonth, issue, "Less than a month")

	// Print issues that are less than a year old.
	printList(iyear, issue, "Less than a year")

	// Print issues that are older than a year.
	printList(imore, issue, "Older than a year")
}

// printList prints out a list of issues under the given header.
func printList(index []date, issue map[int]github.Issue, header string) {
	if len(index) > 0 {
		fmt.Println(header)
		for n, d := range index {
			item := issue[d.i]
			printLine(item)
			index[n].p = true
		}
		println()
	}
}

// Print out a single line issue with only the title and details.
func printLine(item github.Issue) {
	fmt.Printf("#%-6d %6.6s %6.6s %40.40s %10.10s\n",
		item.Number, item.User.Login,
		item.Repo[strings.LastIndex(item.Repo, "/")+1:],
		item.Title, item.CreatedAt.String())
}
