package github

import (
	"fmt"
)

// Open and edit and existing issue.
func EditIssue(conf Config) error {

	// Serch for requested issue.
	var issue []*Issue
	issue, err := searchIssues(conf)
	if err != nil {
		return fmt.Errorf("searchIssues: %v", err)
	}

	// If issue not found or to many.
	l := len(issue)
	if l == 0 {
		return fmt.Errorf("Issue %s not found.", conf.Number)
	} else if l > 1 {
		return fmt.Errorf("Multiple issues returned.")
	}

	// Edit existing data.
	json, err := editIssueData(conf, *issue[0])
	if err != nil {
		return fmt.Errorf("editIssueData: %v", err)
	}

	editIssue(conf, json)

	return nil
}
