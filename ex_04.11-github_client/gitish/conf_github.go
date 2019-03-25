package gitish

import (
	"time"
)

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
   Search request.
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// Reply is an evelope to encapsulate http responses.
type Reply struct {
	Type flags
	Msg  interface{}
}

// IssuesSearchResult is github API json object, a wrapper for an array of
// github issues.
type IssuesSearchResult struct {
	TotalCount int      `json:"total_count"`
	Items      []*Issue // Github issues.
}

// Issue represents a github API json object, containing the data from a github
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

// User represents a github API json object, containing a github user's details.
type User struct {
	Author  string
	Login   string
	HTMLURL string `json:"html_url"`
}
