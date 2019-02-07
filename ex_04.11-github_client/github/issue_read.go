package github

import "fmt"

func ReadIssue(conf Config) error {

	// Serch for requested issue.
	issue, err := searchIssues(conf)
	if err != nil {
		return fmt.Errorf("searchIssue: %v", err)
	}

	// Display.
	printIssue(*issue[0])

	return err
}
