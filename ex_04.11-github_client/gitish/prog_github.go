package gitish

import (
	"log"
	"os"
	"time"
)

var errlog = log.New(os.Stderr, "github: ", log.Lshortfile)

// The run state of the program, interpreted by commandline flags. This
// variable is set as in integer within the configuration sturct, by the
// function InitState(c Config) at program start.
const (
	MoNone = iota
	MoList
	MoRead
	MoEdit
	MoLock
	MoRaise
)

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
	Author   string
	HTMLURL string `json:"html_url"`
}

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
   Configuration and requests
*  ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// Config is a struct specific to the program that contains the principal
// program settings.
type Config struct {
	Author   string // Author user name.
	Token   string // Flag defined Oath token.
	Editor  string // Flag defined external editor.
	Edit    bool   // Signal request to edit an issue.
	Lock    bool   // Lock state.
	Reason  string // Reason for lock.
	Verbose bool   // Signals the program print out extra detail.
	Request        // Stores the users request data.
}

// Request is a struct containing the details of a particular request.
type Request struct {
	Mode    int      // Program running mode.
	User   string   // Repository owner,
	Org     string   // Organisation.
	Repo    string   // Repository name.
	Number  string   // Issue number.
	Queries []string // GET queries retrieved from the Args[] array.
}
