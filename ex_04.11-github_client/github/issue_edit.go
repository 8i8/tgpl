package github

import (
	"bytes"
	"fmt"
)

// EditIssue opens and edits and existing issue.
// TODO NOW is it possible to have only one result possibele and negate the need
// for a test for 1?
func EditIssue(conf Config, issue Issue) (*bytes.Buffer, error) {

	// Edit existing data.
	json, err := editIssue(conf, issue)
	if err != nil {
		return nil, fmt.Errorf("editIssue: %v", err)
	}

	return json, nil
}
