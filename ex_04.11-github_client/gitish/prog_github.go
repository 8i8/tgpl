package gitish

import (
	"log"
	"os"
	"time"
)

var errlog = log.New(os.Stderr, "gitish: ", log.Lshortfile)

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
   Search request.
*  ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// Reply is an evelope that encapsulates the programs http responce, the struct
// contained in the Msg is defined by the type.
type Reply struct {
	Type Flags
	Msg  interface{}
}

// IssuesSearchResult is a json object that contains an array of gitish issues.
type IssuesSearchResult struct {
	TotalCount int      `json:"total_count"`
	Items      []*Issue // Github issues.
}

// Issue represents a json object which contains the data from a gitish
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

// User represents a json object which contains a gitish user details.
type User struct {
	Author  string
	Login   string
	HTMLURL string `json:"html_url"`
}
