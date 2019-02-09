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
	cNone = iota
	cSearch
	cAddress
	cLock
)

var state = cNone

// Check if requirements have been met to enter cAddress mode.
func checkAddress(c Config) bool {
	return (len(c.Login) > 0 || len(c.Author) > 0 ||
		len(c.Owner) > 0 || len(c.Org) > 0) &&
		len(c.Repo) > 0
}

// SetState defines a state by which to run the program, from the selection of
// possible uses inferred from the given flags.
func SetState(c *Config) error {

	var err error
	// If a lock type has been set, force lock mode.
	if len(c.Lock) > 0 {
		c.Mode = "lock"
	}
	// If an issue number has been given and all parameters exist for a
	// direct HTTP access then do so, else add the number to the query
	// listing as a search paramiter.
	if len(c.Number) > 0 && c.Mode != "edit" && c.Mode != "lock" {
		if checkAddress(*c) {
			c.Mode = "read"
		} else {
			c.Queries = append(c.Queries, c.Number)
		}
	}
	// Set to the default mode if none designated.
	if c.Mode == "def" {
		c.Mode = "list"
	}
	// Set the run state.
	if c.Mode == "list" {
		state = cSearch
	} else if c.Mode == "read" {
		state = cAddress
	} else if c.Mode == "raise" {
		state = cAddress
	} else if c.Mode == "lock" {
		if len(c.Lock) == 0 {
			err = errors.New("please designate a reason for locking the thread using the -k flag")
			c.Mode = "error"
		}
		state = cLock
	} else if c.Mode == "edit" {
		if len(c.Editor) == 0 {
			err = errors.New("please designate an external editor `-e <command>` and try again")
			c.Mode = "error"
		}
		state = cAddress
	}

	if c.Verbose {
		fmt.Printf("Setting mode: %v\n", c.Mode)
	}
	return err
}

// Structure an http request from available data.
func setURL(conf Config) (string, string, error) {

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
			err = errors.New("state list; definition requirments were not completed")
		}

	// Prepare URL for API reading repo issues directly by full address and
	// issue number.
	case "read":
		HTTP = "GET"
		URL, err = urlAddressNumber(conf, URL)

	// Prepare for editing a preexisting repo.
	case "edit":
		HTTP = "PATCH"
		URL, err = urlAddressNumber(conf, URL)

	// Prepare URL for issue creation by way of a compleet issue address
	// and the use of the POST function, requires login aurorisation.
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
			err = errors.New("state raise; Please provide owner and repository details")
		}

	// Prepare a URL to set the current issue status to resolved, requires
	// login.
	case "lock":
		// PUT /repos/:owner/:repo/issues/:number/lock
		HTTP = "PUT"
		URL, err = urlAddressNumber(conf, URL)
		URL += "/lock"
	}

	// Add queries to url.
	if len(conf.Queries) > 0 && state != cLock {
		q := url.QueryEscape(strings.Join(conf.Queries, " "))
		URL = URL + "?q=" + q
	}

	if state == cLock {
		URL = URL + "?lock_reason=" + conf.Lock
	}

	// If verbose flag is set print the address used.
	if conf.Verbose {
		fmt.Printf("Setting URL: %v: %v\n", HTTP, URL)
	}

	return HTTP, URL, err
}

// urlAddressNumber sets the url.
func urlAddressNumber(conf Config, URL string) (string, error) {

	var err error
	if len(conf.Owner) > 0 && len(conf.Repo) > 0 && len(conf.Number) > 0 {
		URL = URL + "repos/" + conf.Owner + "/" + conf.Repo + "/issues/" + conf.Number
	} else if len(conf.Login) > 0 && len(conf.Repo) > 0 && len(conf.Number) > 0 {
		URL = URL + "repos/" + conf.Login + "/" + conf.Repo + "/issues/" + conf.Number
	} else if len(conf.Author) > 0 && len(conf.Repo) > 0 && len(conf.Number) > 0 {
		URL = URL + "repos/" + conf.Author + "/" + conf.Repo + "/issues/" + conf.Number
	} else if len(conf.Org) > 0 && len(conf.Repo) > 0 && len(conf.Number) > 0 {
		URL = URL + "orgs/" + conf.Org + "/" + conf.Repo + "/issues/" + conf.Number
	} else {
		err = errors.New("state read; Please provide owner, repository and issue number")
	}

	return URL, err
}
