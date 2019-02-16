package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// respDecode decodes an http responce dependant upon the expected responce
// state, into a single issue or an array of issues as required.
func respDecode(c Config, resp *http.Response) (interface{}, error) {

	var err error

	if c.Verbose {
		fmt.Println("respDecode: attempting decode")
	}

	// Decode into either an issue struct or an array of issue structs.
	if rState == rLone {

		// Single issue.
		var issue Issue
		if err = json.NewDecoder(resp.Body).Decode(&issue); err != nil {
			return nil, fmt.Errorf("json decoder failed: %v", err)
		}
		if c.Verbose {
			fmt.Println("respDecode: single issue decode")
		}
		return issue, err

	} else if rState == rMany {

		// Array of issues.
		var issue IssuesSearchResult
		if err = json.NewDecoder(resp.Body).Decode(&issue); err != nil {
			return nil, fmt.Errorf("json decoder failed: %v", err)
		}
		if c.Verbose {
			fmt.Println("respDecode: multiple issue decode")
		}
		result := issue.Items
		return result, err
	}

	return nil, nil
}

// issueToJSON marshals data into json format and returns it in a bytes buffer.
func issueToJSON(title, body string) (*bytes.Buffer, error) {

	// Write data into a struct.
	var issue Issue
	issue.Title = title
	issue.Body = body

	// Marshal the struct
	json, err := json.Marshal(issue)
	if err != nil {
		return nil, fmt.Errorf("Marshal: %v", err)
	}

	// Write into a byte buffer.
	var b bytes.Buffer
	b.Write(json)

	return &b, err
}
