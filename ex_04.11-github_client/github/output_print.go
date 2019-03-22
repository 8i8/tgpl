package github

import (
	"fmt"
	"strings"
	"time"
)

// Print a list of single line issues, ordered and grouped by date.
func printIssues(retIssues IssueMap) {

	issue := *retIssues.M
	index := *retIssues.I
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
func printIssue(reply Reply) {
	item := reply.Msg.(Issue)
	fmt.Printf("number: #%-6d\nuser: %v\nrepo: %v\ntitle: %v\ncreated: %10.10v\n",
		item.Number, item.User.Login,
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
		item.Number, item.User.Login,
		item.Repo[strings.LastIndex(item.Repo, "/")+1:],
		item.Title, item.CreatedAt.String())
}

// OutputResponse prints the resultset to the terminal.
func OutputResponse(c Config, reply Reply) error {

	if f&cVERBOSE > 0 {
		fmt.Println("Program output start\n~~~\n")
	}

	// Print out either a date ordered list of many issues else a detailed
	// print of one issue.
	var err error
	switch {
	case f&cLIST > 0:
		err = listIssues(c, reply)
	case f&cLIST == 0:
		printIssue(reply)
	default:
		err = fmt.Errorf("OutputResponse: end of switch stament")
	}

	if f&cVERBOSE > 0 {
		fmt.Println("\n~~~\nProgram output end\n")
	}
	return err
}
