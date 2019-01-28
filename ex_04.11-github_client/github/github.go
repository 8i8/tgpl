package github

import (
	"time"
)

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

type Config struct {
	User
	Token  string
	Editor string
	Request
}

type Request struct {
	Name    User
	Repo    string
	Queries string
}

func (c Config) Strings() []string {
	var s []string
	s = append(s, "repo:"+c.Repo)
	s = append(s, c.Queries)
	return s
}
