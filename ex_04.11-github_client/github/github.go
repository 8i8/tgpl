// GET /issues                               // all issues for current user
// GET /user/issues                          // all issues for user
// GET /orgs/:org/issues                     // all issues for organisation
// GET /repos/:owner/:repo/issues            // all issues for reps
// GET /repos/:owner/:repo/issues/:number    // single issue
// POST /repos/:owner/:repo/issues           // Create issue
// PATCH /repos/:owner/:repo/issues/:number  // edit issue

package github

import "time"

// go run issues.go repo:golang/go is:open json decoder
const IssuesURL = "https://api.github.com/search/issues"

//const IssuesURL = "https://api.github.com/repo/8i8/search/issues"
const IssuesPostURL = "https://api.github.com/repos/8i8/test/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

//!-
