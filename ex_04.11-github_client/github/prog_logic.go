/*
	       |GET  |POST |PATCH|PUT
	-------|-----|-----|-----|-----
	list   | 1   |     |     |
	-------|-----|-----|-----|-----
	read   | 1   |     |     |
	-------|-----|-----|-----|-----
	raise  |     | 1   |     |
	-------|-----|-----|-----|-----
	edit   |     |     | 1   |
	-------|-----|-----|-----|-----
	resolve|     |     |     | 1

	GET /issues                                  // all issues, make explicit to avout waste
 	GET /user/issues                             // all issues for user
 	GET /orgs/:org/issues                        // all issues for organisation
 	GET /repos/:owner/:repo/issues               // all issues for reps
 	GET /repos/:owner/:repo/issues/:number       // single issue
 	POST /repos/:owner/:repo/issues              // Create issue
 	PATCH /repos/:owner/:repo/issues/:number     // edit issue
 	PUT /repos/:owner/:repo/issues/:number/lock  // Lock or delete an issue

	https://api.github.com/search/issues

*/
package github

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

const (
	NONE = iota
	SEARCH
	ADDRESS
)

var state int = NONE

// Check if requirements have been met to enter ADDRESS mode.
func checkAddress(c Config) bool {
	return (len(c.Login) > 0 || len(c.Author) > 0 ||
		len(c.Owner) > 0 || len(c.Org) > 0) &&
		len(c.Repo) > 0
}

// Define a state by which to run the program, from the selection of possible
// uses inferred from the given flags.
func SetState(c *Config) int {

	// If an issue number has been given and all parameters exist for a
	// direct HTTP access then do so, else add the number to the query
	// listing as a search paramiter.
	if len(c.Number) > 0 {
		if checkAddress(*c) {
			c.Mode = "read"
		} else {
			c.Queries = append(c.Queries, c.Number)
		}
	}
	// Set to the default mode if no other mode has been set.
	if c.Mode == "def" {
		c.Mode = "list"
	}

	// Set the run state.
	if c.Mode == "list" {
		state = SEARCH
	} else if c.Mode == "read" {
		state = ADDRESS
	} else if c.Mode == "raise" {
		state = ADDRESS
	}
	return state
}

// Structure an http request from available data.
func setUrl(conf Config) (string, string, error) {

	var HTTP string
	var err error
	URL := "https://api.github.com/"

	switch conf.Mode {

	// Prepare URL for API search functionaltiy, add falg designated
	// information to the query list.
	case "list":
		HTTP = "GET"
		URL = URL + "search/issues"
		if len(conf.Owner) > 0 && len(conf.Repo) > 0 {
			conf.Queries = append(conf.Queries, "repo:"+conf.Owner+"/"+conf.Repo)
		} else if len(conf.Owner) > 0 {
			conf.Queries = append(conf.Queries, "user:"+conf.Owner)
		} else if len(conf.Org) > 0 && len(conf.Repo) > 0 {
			conf.Queries = append(conf.Queries, "org:"+conf.Org+"/"+conf.Repo)
		} else if len(conf.Org) > 0 {
			conf.Queries = append(conf.Queries, "org:"+conf.Org)
		} else if len(conf.Author) > 0 && len(conf.Repo) > 0 {
			conf.Queries = append(conf.Queries, "repo:"+conf.Author+"/"+conf.Repo)
		} else if len(conf.Author) > 0 {
			conf.Queries = append(conf.Queries, "author:"+conf.Author)
		} else if len(conf.Login) > 0 && len(conf.Repo) > 0 {
			conf.Queries = append(conf.Queries, "repo:"+conf.Login+"/"+conf.Repo)
		} else if len(conf.Login) > 0 {
			conf.Queries = append(conf.Queries, "author:"+conf.Login)
		} else {
			err = errors.New("state list; definition requirments were not completed.")
		}

	// Prepare URL for API readin repo issues directly by full address and
	// issue number.
	case "read":
		HTTP = "GET"
		if len(conf.Owner) > 0 && len(conf.Repo) > 0 && len(conf.Number) > 0 {
			URL = URL + "repos/" + conf.Owner + "/" + conf.Repo + "/issues/" + conf.Number
		} else if len(conf.Login) > 0 && len(conf.Repo) > 0 && len(conf.Number) > 0 {
			URL = URL + "repos/" + conf.Login + "/" + conf.Repo + "/issues/" + conf.Number
		} else if len(conf.Author) > 0 && len(conf.Repo) > 0 && len(conf.Number) > 0 {
			URL = URL + "repos/" + conf.Author + "/" + conf.Repo + "/issues/" + conf.Number
		} else if len(conf.Org) > 0 && len(conf.Repo) > 0 && len(conf.Number) > 0 {
			URL = URL + "orgs/" + conf.Org + "/" + conf.Repo + "/issues/" + conf.Number
		} else {
			err = errors.New("state read; Please provide owner, repository and issue number.")
		}

	// Prepare URL for issue creation by way of a compleet issue address
	// and the use of the POST function.
	case "raise":
		HTTP = "POST"
		if len(conf.Owner) > 0 && len(conf.Repo) > 0 {
			URL = URL + "repos/" + conf.Owner + "/" + conf.Repo + "/issues"
		} else if len(conf.Login) > 0 && len(conf.Repo) > 0 {
			URL = URL + "repos/" + conf.Login + "/" + conf.Repo + "/issues"
		} else if len(conf.Author) > 0 && len(conf.Repo) > 0 {
			URL = URL + "repos/" + conf.Author + "/" + conf.Repo + "/issues"
		} else if len(conf.Org) > 0 && len(conf.Repo) > 0 {
			URL = URL + "orgs/" + conf.Org + "/" + conf.Repo + "/issues"
		} else {
			err = errors.New("state raise; Please provide owner and repository details.")
		}
	}

	// Add queries to url.
	if len(conf.Queries) > 0 {
		q := url.QueryEscape(strings.Join(conf.Queries, " "))
		URL = URL + "?q=" + q
	}

	// If verbose flag is set print the address used.
	if conf.Verbose {
		fmt.Println(HTTP, URL)
	}

	return HTTP, URL, err
}
