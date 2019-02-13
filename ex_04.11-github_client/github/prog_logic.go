/*
Package github - Command line client for the github issue API.

SYNOPSIS
	github [user | repo | number][Oauth2][options]

DESCRIPTION
	github is a github client designed for raising and tracking and
	updating github issues on the github platform from the users command
	line by way of the github HTTP API. Giving the user access from the
	command line or their favorite editor application.

MAIN
	The github program has essentially five running modes, the mode is set
	from the main function according to the flags set state, defined in the
	SetState() function.

HTTP REQUESTS
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

	GET    /issues
 	GET    /user/issues
 	GET    /orgs/:org/issues
	GET    /search/issues?q= user:[user] | repo:[repo] | author:[author]
 	GET    /repos/:owner/:repo/issues
 	GET    /repos/:owner/:repo/issues/:number
 	POST   /repos/:owner/:repo/issues
 	PATCH  /repos/:owner/:repo/issues/:number
 	PUT    /repos/:owner/:repo/issues/:number/lock?lock_reason=[reason]
	DELETE /repos/:owner/:repo/issues/:number/lock

	https://api.github.com/search/issues

PROGRAM STATES
	Table representation of program states, the program has essentially two
	different primary states, the first of which is prevalent in the main
	function, designating the programs initial running mode to establish the
	type of HTTP request to be made. The second defines the formation of
	the URL for the request.

	┌────────┬────────┬────────┬────────┬────────┬────────┬──────────────┐
	│-o org  │        │        │        │-l lock │        │              │
	│-a auth │        │        │        │-e edit │        │    State     │
	│-u user │-r repo │-n numb │-t token│-x raise│-d[exec]│              │
	├────────┼────────┼────────┼────────┼────────┼────────┼──────────────┤
	│ yes    │        │        │ N/A    │ N/A    │ all    │ mList  rMany │
	├────────┼────────┼────────┼────────┼────────┼────────┼──────────────┤
	│        │ yes    │        │ N/A    │ N/A    │ all    │ mList  rMany │
	├────────┼────────┼────────┼────────┼────────┼────────┼──────────────┤
	│ yes    │ yes    │        │ N/A    │ N/A    │ all    │ mList  rMany │
	├────────┼────────┼────────┼────────┼────────┼────────┼──────────────┤
	│ yes    │ no/fill│ yes    │ N/A    │ N/A    │ all    │ mList  rMany │
	├────────┼────────┼────────┼────────┼────────┼────────┼──────────────┤
	│ no/fill│ yes    │ yes    │ N/A    │ N/A    │ all    │ mList  rMany │
	├────────┼────────┼────────┼────────┼────────┼────────┼──────────────┤
	│ yes    │ yes    │ yes    │ N/A    │ N/A    │ all    │ mRead  rLone │
	├────────┼────────┼────────┼────────┼────────┼────────┼──────────────┤
	│ yes    │ yes    │ yes    │ yes    │ -e     │ all    │ mEdit  rLone │
	├────────┼────────┼────────┼────────┼────────┼────────┼──────────────┤
	│ yes    │ yes    │ yes    │ yes    │ -x     │ all    │ mRais  rNone │
	├────────┼────────┼────────┼────────┼────────┼────────┼──────────────┤
	│ yes    │ yes    │ yes    │ yes    │ -l     │ all    │ mLock  rNone │
	└────────┴────────┴────────┴────────┴────────┴────────┴──────────────┘

	-v displays verbose report of the programs actions.
	-m defines the external editor to be used in editing.

*/
package github

import (
	"fmt"
	"net/url"
	"strings"
)

// The run state of the program, interpreted by commandline flags. This
// variable is set as in integer within the configuration sturct, by the
// function SetState(c Config) at program start.
const (
	mList = iota
	mRead
	mEdit
	mLock
	mRais
)

// Mode of the expected http response type.
const (
	rNone = iota
	rLone
	rMany
)

var state = rNone

// isFullAddress checks if the requirements have been met to enter rLone
// mode.
func isFullAddress(c Config) bool {
	return (len(c.Author) > 0 || len(c.User) > 0 || len(c.Org) > 0) &&
		len(c.Repo) > 0
}

// checkMode verifies that there are not two contradicting flags set.
func checkMode(c Config) bool {
	tally := 0
	if c.Edit {
		tally++
	}
	if c.Lock {
		tally++
	}
	if c.Raise {
		tally++
	}
	if tally > 1 {
		return true
	}
	return false
}

// SetState defines the state in which to run the program, set by the
// configuration the of users input data.
func SetState(c *Config) error {

	var err error
	c.Mode = mList

	if checkMode(*c) {
		str := "Please provide only one of the following flags x e or l"
		return fmt.Errorf(str)
	}

	if c.Edit {
		c.Mode = mEdit
		state = rLone
	} else if c.Lock {
		c.Mode = mEdit
		state = rLone
	} else if c.Raise {
		c.Mode = mRais
		state = rNone
	}
	// If an issue number has been given and all parameters exist for
	// a direct HTTP access then do so, else add the number to the
	// query listing as a search parameter, expect multiple results.
	if len(c.Number) > 0 && (!c.Edit || !c.Lock) {
		if isFullAddress(*c) {
			c.Mode = mRead
		} else {
			c.Queries = append(c.Queries, c.Number)
			c.Mode = mList
		}
	}
	// Set the run state.
	if c.Mode == mList {
		state = rMany
	} else if c.Mode == mRead {
		state = rLone
	} else if c.Mode == mRais {
		state = rNone
	}

	if c.Verbose {
		fmt.Printf("Setting mode: %v\n", c.Mode)
	}
	return err
}

// Structure an http request from available data.
func setURL(conf Config) (Address, error) {

	var addr Address
	var err error
	addr.Url = "https://api.github.com/"

	switch conf.Mode {

	// Prepare URL for API search functionality, add flag designated
	// information to the query list.
	case mList:
		addr.Http = "GET"
		addr.Url = addr.Url + "search/issues"
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
	case mRead:
		addr.Http = "GET"
		str := "Please specify owner, repository and issue number."
		addr.Url, err = urlAddrIssues(conf, addr.Url, "read", str)
		if err != nil {
			return addr, err
		}
		addr.Url += conf.Number

	case mEdit:
		// Prepare for editing a preexisting repo.
		addr.Http = "PATCH"
		str := "Please specify owner, repository and issue number."
		addr.Url, err = urlAddrIssues(conf, addr.Url, "edit", str)
		if err != nil {
			return addr, err
		}
		addr.Url += conf.Number

	case mLock:
		// Prepare a URL to set the current issue status to resolved,
		// requires login.
		addr.Http = "PUT"
		str := "Please specify owner, repository and issue number."
		addr.Url, err = urlAddrIssues(conf, addr.Url, "lock", str)
		if err != nil {
			return addr, err
		}
		addr.Url += conf.Number + "/lock"

	// Prepare URL for issue creation by way of a complete issue address
	// and the use of the POST function, requires login authorisation.
	case mRais:
		addr.Http = "POST"
		str := "Please specify owner and repository details"
		addr.Url, err = urlAddrIssues(conf, addr.Url, "raise", str)
		if err != nil {
			return addr, err
		}
	}

	// Add queries to url.
	if len(conf.Queries) > 0 && conf.Mode != mLock {
		q := url.QueryEscape(strings.Join(conf.Queries, " "))
		addr.Url = addr.Url + "?q=" + q
	}

	// If lock required, add query.
	if conf.Mode == mLock {
		addr.Url = addr.Url + "?lock_reason=" + conf.Reason
	}

	// If verbose flag is set print the address used.
	if conf.Verbose {
		fmt.Printf("Setting URL: %v %v\n", addr.Http, addr.Url)
	}

	return addr, err
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
