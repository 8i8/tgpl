package github

import (
	"fmt"
	"strings"
	"time"
)

// Print to terminal order by date, separating newer than one month and within
// a year.
func PrintIssues(results IssueMap) {

	// Check that records exist to print.
	if results.Len() == 0 {
		fmt.Println("Empty result string.")
		return
	}

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
			printLine(item)
			index[n].p = true
		}
	}

	// Print issues that are less than a year old.
	fmt.Println("\nLess than a year")
	for n, d := range index {
		item := issue[d.r]
		if monthToAYear(now, d) {
			printLine(item)
			index[n].p = true
		}
	}

	// Print issues that are older than a year.
	fmt.Println("\nOlder than a year")
	for n, d := range index {
		item := issue[d.r]
		if yearOnward(now, d) {
			printLine(item)
			index[n].p = true
		}
	}
}

func printIssue(item Issue) {
	fmt.Printf("#%-6d %6.6s %6.6s %40.40s %10.10s\n",
		item.Number, item.User.Login,
		item.Repo[strings.LastIndex(item.Repo, "/")+1:],
		item.Title, item.CreatedAt.String())
	fmt.Printf("%v\n", item.Body)
}

func printLine(item Issue) {
	fmt.Printf("#%-6d %6.6s %6.6s %40.40s %10.10s\n",
		item.Number, item.User.Login,
		item.Repo[strings.LastIndex(item.Repo, "/")+1:],
		item.Title, item.CreatedAt.String())
}
