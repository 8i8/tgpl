package github

import (
	"log"
	"os"
	"time"
)

var errlog = log.New(os.Stderr, "github: ", log.Lshortfile)

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
   Search request.
*  ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// IssuesSearchResult is a json object that contains an array of github issues.
type IssuesSearchResult struct {
	TotalCount int      `json:"total_count"`
	Items      []*Issue // Github issues.
}

// Issue represents a json object which contains the data from a github
// repository issue.
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string `json:"title"`
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body"` // in Markdown format
	Repo      string    `json:"repository_url"`
	Locked    bool      `json:"locked"`
	Reason    string    `json:"active_lock_reason"`
}

// User represents a json object which contains a github user details.
type User struct {
	Author  string
	Login   string
	HTMLURL string `json:"html_url"`
}
