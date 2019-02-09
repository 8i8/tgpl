package github

import (
	"fmt"
)

// EditIssue opens and edits and existing issue.
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
		return fmt.Errorf("issue %s not found", conf.Number)
	} else if l > 1 {
		return fmt.Errorf("multiple issues returned")
	}

	if issue[0].Locked {
		return fmt.Errorf("the issue is currently locked")
	}

	// Edit existing data.
	json, err := editIssueData(conf, *issue[0])
	if err != nil {
		return fmt.Errorf("editIssueData: %v", err)
	}

	err = editIssue(conf, json)
	if err != nil {
		return fmt.Errorf("editIssue: %v", err)
	}

	return nil
}
