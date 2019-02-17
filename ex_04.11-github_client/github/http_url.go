package github

import (
	"fmt"
	"net/url"
	"strings"
)

// setURLSearchQuery structures and adds to the query array the required key
// values paries to instigate an API search query.
// https://api.github.com/search/issues
func setURLSearchQuery(c *Config, e string) error {

	var err error
	if len(c.User) > 0 && len(c.Repo) > 0 {
		c.Queries = append(
			c.Queries, "repo:"+c.User+"/"+c.Repo)
	} else if len(c.User) > 0 {
		c.Queries = append(
			c.Queries, "user:"+c.User)
	} else if len(c.Org) > 0 && len(c.Repo) > 0 {
		c.Queries = append(
			c.Queries, "org:"+c.Org+"/"+c.Repo)
	} else if len(c.Org) > 0 {
		c.Queries = append(c.Queries, "org:"+c.Org)
	} else if len(c.Author) > 0 && len(c.Repo) > 0 {
		c.Queries = append(
			c.Queries, "repo:"+c.Author+"/"+c.Repo)
	} else if len(c.Author) > 0 {
		c.Queries = append(
			c.Queries, "author:"+c.Author)
	} else {
		err = fmt.Errorf("%v: setURLSearchQuery: %v", mState[c.Mode], e)
	}
	return err
}

// setURLAddress is a helper funcion for setURL, defines the base address for
// the mRead, mEdit, mLock and mRais modes.
func setURLAddress(c Config, URL, e string) (string, error) {

	var err error
	if len(c.User) > 0 && len(c.Repo) > 0 {
		URL = URL + "repos/" + c.User + "/" + c.Repo + "/issues"
	} else if len(c.Author) > 0 && len(c.Repo) > 0 {
		URL = URL + "repos/" + c.Author + "/" + c.Repo + "/issues"
	} else if len(c.Org) > 0 && len(c.Repo) > 0 {
		URL = URL + "orgs/" + c.Org + "/" + c.Repo + "/issues"
	} else {
		err = fmt.Errorf("%v: setURLAddress: %v", mState[c.Mode], e)
	}

	return URL, err
}

// setURL structures an http request from the given configuration.
func setURL(c Config) (Address, error) {

	var addr Address
	var err error

	// Set the base address.
	addr.URL = "https://api.github.com/"

	// Dependant upon the program runnnig mode, generate the required URL
	// and or query set.
	switch c.Mode {

	// Prepare URL for API search functionality
	case mList:
		addr.HTTP = "GET"
		addr.URL = addr.URL + "search/issues"
		str := "url requirements were not met"
		err := setURLSearchQuery(&c, str)
		if err != nil {
			return addr, err
		}

	// Prepare URL for API reading repo issues directly by full address and
	// issue number.
	case mRead:
		addr.HTTP = "GET"
		str := "please specify owner, repository and issue number"
		addr.URL, err = setURLAddress(c, addr.URL, str)
		if err != nil {
			return addr, err
		}
		addr.URL += "/" + c.Number

	case mEdit:
		// Prepare for editing a preexisting repo.
		addr.HTTP = "PATCH"
		str := "please specify owner, repository and issue number"
		addr.URL, err = setURLAddress(c, addr.URL, str)
		if err != nil {
			return addr, err
		}
		addr.URL += "/" + c.Number

	case mLock:
		// Prepare a URL to set the current issue status to resolved,
		// requires login.
		addr.HTTP = "PUT"
		str := "please specify owner, repository and issue number"
		addr.URL, err = setURLAddress(c, addr.URL, str)
		if err != nil {
			return addr, err
		}
		addr.URL += "/" + c.Number + "/lock"

	// Prepare URL for issue creation by way of a complete issue address
	// and the use of the POST function, requires login authorisation.
	case mRais:
		addr.HTTP = "POST"
		str := "please specify owner and repository details"
		addr.URL, err = setURLAddress(c, addr.URL, str)
		if err != nil {
			return addr, err
		}
	case mRaw:
		if len(c.Queries) < 2 {
			str := "please provide http request type and address"
			return addr, fmt.Errorf(str)
		}
		// Fill the address fields from the command line query
		// arguments.
		addr.HTTP = c.Queries[0]
		addr.URL = c.Queries[1]
		// Remove the first two queries.
		c.Queries = c.Queries[2:]
	}

	// Add queries to url.
	if len(c.Queries) > 0 && c.Mode != mLock {
		q := url.QueryEscape(strings.Join(c.Queries, " "))
		addr.URL = addr.URL + "?q=" + q
	}

	// If lock required, add query.
	if c.Mode == mLock {
		addr.URL = addr.URL + "?lock_reason=" + c.Reason
	}

	// If verbose flag is set print the address used.
	if c.Verbose {
		fmt.Printf("Setting URL: %v %v\n", addr.HTTP, addr.URL)
	}

	return addr, err
}
