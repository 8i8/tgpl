package itr

import (
	"time"
)

// Issue a struct that stores a github API json object, containing the data
// from a github repository issue.
type Issue struct {
	Id            uint64
	NodeId        string `json:"node_id"`
	URL           string `json:"url"`
	RepositoryURL string `json:"repository_url"`
	LablesURL     string `json:"labels_url"`
	CommentsURL   string `json:"comments_url"`
	EventsURL     string `json:"events_url"`
	HtmlURL       string `json:"html_url"`
	Number        uint64
	State         string
	Title         string
	Body          string   // Markdown format
	User          uint64   // User
	Labels        []uint64 // Label
	Assignee      uint64   // User
	Assignees     []uint64 // User
	Milestone     uint64
	Locked        bool
	Reason        string `json:"active_lock_reason"`
	Comments      int
	PullRequest   PullRequest
	ClosedAt      time.Time `json:"closed_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	ClosedBy      uint64    `json:"closed_by"`
}

// User is a struch that stores a github API json object, containing a github
// user's details.
type User struct {
	Login            string
	Id               uint64
	NodeId           string `json:"node_id"`
	AvatarURL        string `json:"avatar_url"`
	GravitarId       string `json:"gravitar_id"`
	URL              string `json:"url"`
	HtmlURL          string `json:"html_url"`
	FollowersURL     string `json:"followers_url"`
	FollowingURL     string `json:"following_url"`
	GistsURL         string `json:"gists_url"`
	StarredURL       string `json:"starred_url"`
	SubscriptionsURL string `json:"subscriptions_url"`
	OrganizationsURl string `json:"organizations_url"`
	ReposURL         string `json:"repos_url"`
	EventsURL        string `json:"events_url"`
	Type             string
	SiteAdmin        bool `json:"site_admin"`
}

// Labels are a struct that stores imformation about a github issues; Labels on
// GitHub help you organize and prioritize your work.
type Label struct {
	Id          uint64
	NodeId      string `json:"node_id"`
	URL         string `json:"url"`
	Name        string
	Description string
	Color       string
	Default     bool
}

// Milestone is a struct that stores milestone data; Milestones are used to
// track progress on groups of issues or pull requests in a repository.
type Milestone struct {
	URL          string `json:"url"`
	HtmlURL      string `json:"html_url"`
	LabelsURL    string `json:"labels_url"`
	Id           uint64
	NodeId       string `json:"node_id"`
	Number       uint64
	State        string
	Title        string
	Description  string
	Creator      uint64    // User
	OpenIssues   int       `json:"open_issues"`
	ClosedIssues int       `json:"closed_issues"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ClosedAt     time.Time `json:"closed_at"`
	DueOn        time.Time `json:"due_on"`
}

// PullRequest is a struct that contains a github pull request, an issue is
// created when a pull request is made, as such they may or may not conjoin an
// issue; A pull request is always atached to and issue but an issue may exist
// without a pull request.
type PullRequest struct {
	URL      string `json:"url"`
	HtmlURL  string `json:"html_url"`
	DiffURL  string `json:"diff_url"`
	PatchURL string `json:"patch_url"`
}
