package github

import (
	"fmt"
	"strings"
	"time"
)

// Print a list of single line issues, ordered and grouped by date.
func printIssues(results IssueMap) {

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

// Print out a single issue including the message body.
func printIssue(item Issue) {
	fmt.Printf("number: #%-6d\nuser: %v\nrepo: %v\ntitle: %v\ncreated: %10.10v\n",
		item.Number, item.User.Author,
		item.Repo[strings.LastIndex(item.Repo, "/")+1:],
		item.Title, item.CreatedAt.String())
	if item.Locked {
		fmt.Printf("status: closed %v\n", item.Reason)
	} else {
		fmt.Println("status: open")
	}
	fmt.Printf("message: %v\n", item.Body)
}

// Print out a single line issue with only the title and details.
func printLine(item Issue) {
	fmt.Printf("#%-6d %6.6s %6.6s %40.40s %10.10s\n",
		item.Number, item.User.Author,
		item.Repo[strings.LastIndex(item.Repo, "/")+1:],
		item.Title, item.CreatedAt.String())
}

// OutputResponce prints the resultset to the terminal.
func OutputResponce(c Config, I interface{}) error {

	if c.Verbose {
		fmt.Println("Program output start\n~~~\n")
	}

	// Print out either a date ordered list of many issues else a detailed
	// print of one issue.
	var err error
	if I != nil {
		switch v := I.(type) {
		case []*Issue:
			err = listIssues(c, v)
		case Issue:
			printIssue(v)
		default:
			err = fmt.Errorf("unknown type")
		}
	}

	if c.Verbose {
		fmt.Println("\n~~~\nProgram output end\n")
	}
	return err
}