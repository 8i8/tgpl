// GET /issues                                  // all issues for current user
// GET /user/issues                             // all issues for user
// GET /orgs/:org/issues                        // all issues for organisation
// GET /repos/:owner/:repo/issues               // all issues for reps
// GET /repos/:owner/:repo/issues/:number       // single issue
// POST /repos/:owner/:repo/issues              // Create issue
// PATCH /repos/:owner/:repo/issues/:number     // edit issue
// PUT /repos/:owner/:repo/issues/:number/lock  // Lock or delete an issue
package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(URL string, terms []string) (*IssuesSearchResult, error) {

	q := url.QueryEscape(strings.Join(terms, " "))

	// Genereate request.
	req, err := http.NewRequest("GET", URL+"?q="+q, nil)
	if err != nil {
		return nil, err
	}

	// Add header to request.
	req.Header.Set(
		"Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Close resp.Body.
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	// Decode reply.
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

// func GetIssue(num int) Issue {
// }

func RaiseIssue(data []Issue, token string) {

	//str := `{"title": "The Go Issue","body": "This is the thing you see, it came from a Go application."}`
	//json := bytes.NewBufferString(str)
	str, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Json Marshal failed : %s\n", err)
	}
	json := bytes.NewBuffer(str)

	// Formulate post request
	req, err := http.NewRequest("POST", IssuesPostURL, json)
	if err != nil {
		fmt.Fprintf(os.Stderr, "POST failed : %s\n", err)
	}

	// Set header.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "token "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "HEAD failed : %s\n", resp.Status)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "Issue creation failed: %s\n", resp.Status)
	}

	resp.Body.Close()
}
