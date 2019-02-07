package github

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
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

	// Get user input.
	json, err := editIssueData(conf, *issue[0])
	if err != nil {
		return fmt.Errorf("editIssueData: %v", err)
	}

	editIssue(conf, json)
	return nil
}

// Edit an existing issue.
func editIssueData(conf Config, issue Issue) (*bytes.Buffer, error) {

	var err error
	// Retreive data from issue.
	currentTitle, currentBody := issue.Title, issue.Body

	// Display the current title and offer the occasion to change it.
	sc := bufio.NewScanner(os.Stdin)
	var title string
	for title == "" {
		fmt.Print("leave empty and press enter to keep the current title.")
		fmt.Print(currentTitle, " : ")
		title = sc.Text()
		title = strings.TrimSpace(title)
		if title == "" {
			title = currentTitle
		}
	}

	// Require a message body be entered, open an editor if desired.
	var body string
	body, err = openEditorBody(conf, currentBody)
	if err != nil {
		return nil, fmt.Errorf("openEditorBody: %+v", err)
	}

	// If either field are empty, refuse to generate.
	if len(title) == 0 || len(body) == 0 {
		return nil, errors.New("Issue fields are empty, refusing to send request.")
	}

	str := `{"title":"` + title + `","body":"` + body + `"}`
	json := bytes.NewBufferString(str)

	return json, nil
}
