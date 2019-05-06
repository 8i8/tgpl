package github

import (
	"fmt"
	"strings"
	"time"
)

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
func printList(index []date, issue map[int]Issue, header string) {
	if len(index) > 0 {
		fmt.Println(header)
		for n, d := range index {
			item := issue[d.i]
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

// outputResponse prints the result set to the terminal.
func outputResponse(c Config, reply Reply) error {

	if f&cVERBOSE > 0 {
		fmt.Printf("Program output start\n~~~\n\n")
	}

	// Print out either a date ordered list of many issues else a detailed
	// print of one issue.
	var err error
	switch {
	case f&cLIST > 0:
		resp := IssueMap{}
		resp, err = listIssues(reply)
		printIssues(resp)
	case f&cREAD > 0:
		if f&cEDITOR > 0 {
			editIssue(c, reply)
		} else {
			printIssue(reply)
		}
	default:
		err = fmt.Errorf("OutputResponse: end of switch stament")
	}

	if f&cVERBOSE > 0 {
		fmt.Printf("\n~~~\nProgram output end\n\n")
	}
	return err
}
