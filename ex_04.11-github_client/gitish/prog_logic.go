/*
NAME
	gitish - Command line client for the github issue API.

SYNOPSIS
	gitish [ user | repo | number ][Oauth2][options]

DESCRIPTION
	gitish is a github client designed for raising and tracking and
	updating github issues on the github platform from the users command
	line by way of the github HTTP API. Giving the user access from the
	command line or their favorite editor application.

GITHUB API HTTP
	┌───────┬───────┬───────┬───────┬───────┬───────┐
	│       │ GET   │ POST  │ PATCH │ PUT   │DELETE │
	├───────┼───────┼───────┼───────┼───────┼───────┤
	│ list  │   1   │       │       │       │       │
	├───────┼───────┼───────┼───────┼───────┼───────┤
	│ read  │   1   │       │       │       │       │
	├───────┼───────┼───────┼───────┼───────┼───────┤
	│ raise │       │   1   │       │       │       │
	├───────┼───────┼───────┼───────┼───────┼───────┤
	│ edit  │       │       │   1   │       │       │
	├───────┼───────┼───────┼───────┼───────┼───────┤
	│ lock  │       │       │       │   1   │       │
	├───────┼───────┼───────┼───────┼───────┼───────┤
        │unlock │       │       │       │       │   1   │
	└───────┴───────┴───────┴───────┴───────┴───────┘

	GET    /search/issues?q= user:[user] | repo:[repo] | author:[author]
	GET    /issues
 	GET    /user/issues
 	GET    /orgs/:org/issues
 	GET    /repos/:owner/:repo/issues
 	GET    /repos/:owner/:repo/issues/:number
 	POST   /repos/:owner/:repo/issues
 	PATCH  /repos/:owner/:repo/issues/:number
 	PUT    /repos/:owner/:repo/issues/:number/lock?lock_reason=[reason]
	DELETE /repos/:owner/:repo/issues/:number/lock

	https://api.github.com/search/issues

PROGRAM STATE TABLE
	Table representation of program states, the program has essentially two
	different primary states, the first of which is prevalent in the main
	function, designating the programs initial running mode to establish the
	type of HTTP request to be made. The second defines the formation of
	the URL for the request.

	┌────────┬────────┬────────┬────────┬──────────┬────────┬────────────┐
	│        │        │        │        │-k ?      │        │            │
	│-o org  │        │        │        │-l lock[r]│        │            │
	│-a auth │        │        │        │-e edit   │        │   State    │
	│-u user │-r repo │-n numb │-t token│-x raise  │-d[exec]│            │
	├────────┼────────┼────────┼────────┼──────────┼────────┼────────────┤
	│ yes    │        │        │ N/A    │ N/A      │ all    │ list  sear │
	├────────┼────────┼────────┼────────┼──────────┼────────┼────────────┤
	│        │ yes    │        │ N/A    │ N/A      │ all    │ list  sear │
	├────────┼────────┼────────┼────────┼──────────┼────────┼────────────┤
	│ yes    │ yes    │        │ N/A    │ N/A      │ all    │ list  sear │
	├────────┼────────┼────────┼────────┼──────────┼────────┼────────────┤
	│ yes    │ no/fill│ yes    │ N/A    │ N/A      │ all    │ list  sear │
	├────────┼────────┼────────┼────────┼──────────┼────────┼────────────┤
	│ no/fill│ yes    │ yes    │ N/A    │ N/A      │ all    │ list  sear │
	├────────┼────────┼────────┼────────┼──────────┼────────┼────────────┤
	│ yes    │ yes    │ yes    │ N/A    │ N/A      │ all    │ read  addr │
	├────────┼────────┼────────┼────────┼──────────┼────────┼────────────┤
	│ yes    │ yes    │ yes    │ yes    │ -x       │ all    │ raise addr │
	├────────┼────────┼────────┼────────┼──────────┼────────┼────────────┤
	│ yes    │ yes    │ yes    │ yes    │ -e       │ all    │ edit  addr │
	├────────┼────────┼────────┼────────┼──────────┼────────┼────────────┤
	│ yes    │ yes    │ yes    │ yes    │ -l       │ all    │ edit  lock │
	├────────┼────────┼────────┼────────┼──────────┼────────┼────────────┤
	│ yes    │ yes    │ yes    │ yes    │ -k       │ all    │ edit  unlk │
	└────────┴────────┴────────┴────────┴──────────┴────────┴────────────┘
*/
package gitish

import (
	"fmt"
	"net/url"
	"strings"
)

const (
	urlNone = iota
	urlSear
	urlAddr
	urlLock
	urlEdit
	urlRead
)

var state = urlNone

// Check if the requirements have been met to enter urlAddr mode.
func checkAddress(c Config) bool {
	return (len(c.Author) > 0 || len(c.User) > 0 || len(c.Org) > 0) &&
		len(c.Repo) > 0
}

// InitState defines the state in which to run the program, set by the
// configuration of user input flags and data.
func InitState(c *Config) error {

	var err error
	if c.Edit {
		c.Mode = MoEdit
		state = urlEdit
	}
	// If a lock type has been set, force lock mode.
	if c.Lock {
		c.Mode = MoEdit
		state = urlLock
	}
	// If an issue number has been given and all parameters exist for
	// a direct HTTP access then do so, else add the number to the
	// query listing as a search parameter.
	if len(c.Number) > 0 && state != urlEdit && state != urlLock {
		if checkAddress(*c) {
			c.Mode = MoRead
			state = urlRead
		} else {
			c.Queries = append(c.Queries, c.Number)
		}
	}
	// Set to the default mode if none designated.
	if c.Mode == MoNone {
		c.Mode = MoList
	}
	// Set the run state.
	if c.Mode == MoList {
		state = urlSear
	} else if c.Mode == MoRead {
		state = urlAddr
	} else if c.Mode == MoRaise {
		state = urlAddr
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

	// Prepare URL for API search functionality, add flag designated
	// information to the query list.
	case MoList:
		HTTP = "GET"
		URL = URL + "search/issues"
		if len(conf.User) > 0 && len(conf.Repo) > 0 {
			conf.Queries = append(
				conf.Queries, "repo:"+conf.User+"/"+conf.Repo)
		} else if len(conf.User) > 0 {
			conf.Queries = append(
				conf.Queries, "user:"+conf.User)
		} else if len(conf.Org) > 0 && len(conf.Repo) > 0 {
			conf.Queries = append(
				conf.Queries, "org:"+conf.Org+"/"+conf.Repo)
		} else if len(conf.Org) > 0 {
			conf.Queries = append(conf.Queries, "org:"+conf.Org)
		} else if len(conf.Author) > 0 && len(conf.Repo) > 0 {
			conf.Queries = append(
				conf.Queries, "repo:"+conf.Author+"/"+conf.Repo)
		} else if len(conf.Author) > 0 {
			conf.Queries = append(
				conf.Queries, "author:"+conf.Author)
		} else {
			err = fmt.Errorf("%v: url definition requirements "+
				"were not met", conf.Mode)
		}

	// Prepare URL for API reading repo issues directly by full address and
	// issue number.
	case MoRead:
		HTTP = "GET"
		str := "Please specify owner, repository and issue number."
		URL, err = urlAddrIssues(conf, URL, "read", str)
		if err != nil {
			return HTTP, URL, err
		}
		URL += conf.Number

	case MoEdit:
		switch state {
		// Prepare for editing a preexisting repo.
		case urlEdit:
			HTTP = "PATCH"
			str := "Please specify owner, repository and issue number."
			URL, err = urlAddrIssues(conf, URL, "edit", str)
			if err != nil {
				return HTTP, URL, err
			}
			URL += conf.Number

		// Prepare a URL to set the current issue status to resolved,
		// requires login.
		case urlLock:
			HTTP = "PUT"
			str := "Please specify owner, repository and issue number."
			URL, err = urlAddrIssues(conf, URL, "lock", str)
			if err != nil {
				return HTTP, URL, err
			}
			URL += conf.Number + "/lock"
		}

	// Prepare URL for issue creation by way of a complete issue address
	// and the use of the POST function, requires login authorisation.
	case MoRaise:
		HTTP = "POST"
		str := "Please specify owner and repository details"
		URL, err = urlAddrIssues(conf, URL, "raise", str)
		if err != nil {
			return HTTP, URL, err
		}
	}

	// Add queries to url.
	if len(conf.Queries) > 0 && state != urlLock {
		q := url.QueryEscape(strings.Join(conf.Queries, " "))
		URL = URL + "?q=" + q
	}

	// If lock required, add query.
	if state == urlLock {
		URL = URL + "?lock_reason=" + conf.Reason
	}

	// If verbose flag is set print the address used.
	if conf.Verbose {
		fmt.Printf("Setting URL: %v: %v\n", HTTP, URL)
	}

	return HTTP, URL, err
}

// urlAddrIssues sets the url.
func urlAddrIssues(conf Config, URL, mode, e string) (string, error) {

	var err error
	if len(conf.User) > 0 && len(conf.Repo) > 0 && len(conf.Number) > 0 {

		URL = URL + "repos/" + conf.User + "/" + conf.Repo + "/issues/"

	} else if len(conf.Author) > 0 && len(conf.Repo) > 0 && len(conf.Number) > 0 {

		URL = URL + "repos/" + conf.Author + "/" + conf.Repo + "/issues/"

	} else if len(conf.Org) > 0 && len(conf.Repo) > 0 && len(conf.Number) > 0 {

		URL = URL + "orgs/" + conf.Org + "/" + conf.Repo + "/issues/"

	} else {
		err = fmt.Errorf("%v: %v", mode, e)
	}

	return URL, err
}
