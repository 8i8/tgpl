package gitish

import (
	"fmt"
	"net/url"
	"strings"
)

// setURLSearchQuery structures and adds to the query array the required key
// values paries to instigate an API search query.
// https://api.gitish.com/search/issues
func setURLSearchQuery(c Config, e string) (Config, error) {

	var err error
	if f&cUSER > 0 && f&cREPO > 0 {
		c.Queries = append(
			c.Queries, "repo:"+c.User+"/"+c.Repo)
	} else if f&cUSER > 0 {
		c.Queries = append(
			c.Queries, "user:"+c.User)
	} else if f&cORG > 0 && f&cREPO > 0 {
		c.Queries = append(
			c.Queries, "org:"+c.Org+"/"+c.Repo)
	} else if f&cORG > 0 {
		c.Queries = append(c.Queries, "org:"+c.Org)
	} else if f&cAUTHOR > 0 && f&cREPO > 0 {
		c.Queries = append(
			c.Queries, "repo:"+c.Author+"/"+c.Repo)
	} else if f&cAUTHOR > 0 {
		c.Queries = append(
			c.Queries, "author:"+c.Author)
	} else if f&cREPO > 0 {
		c.Queries = append(
			c.Queries, "repo:"+c.Repo)
	} else {
		err = fmt.Errorf("setURLSearchQuery: %v", e)
	}
	return c, err
}

// setURLAddress is a helper funcion for setURL, defines the base address for
// the cREAD, cEDIT, cLOCK and cRAISE modes.
func setURLAddress(c Config, URL, e string) (Config, string, error) {

	var err error
	if len(c.User) > 0 && len(c.Repo) > 0 {
		URL = URL + "repos/" + c.User + "/" + c.Repo + "/issues"
	} else if len(c.Author) > 0 && len(c.Repo) > 0 {
		URL = URL + "repos/" + c.Author + "/" + c.Repo + "/issues"
	} else if len(c.Org) > 0 && len(c.Repo) > 0 {
		URL = URL + "orgs/" + c.Org + "/" + c.Repo + "/issues"
	} else {
		err = fmt.Errorf("setURLAddress: %v", e)
	}

	return c, URL, err
}

// setURL structures an http request from the given configuration.
func setURL(c Config) (Config, Address, error) {

	var addr Address
	var err error

	// Set the base address.
	addr.URL = "https://api.github.com/"

	// Dependant upon the program runnnig mode, generate the required URL
	// and or query set.
	switch {

	// Prepare URL for API search functionality
	case f&cLIST > 0:
		addr.HTTP = "GET"
		addr.URL = addr.URL + "search/issues"
		str := "url requirements were not met"
		c, err = setURLSearchQuery(c, str)
		if err != nil {
			return c, addr, fmt.Errorf("LIST: %v", err)
		}

	// Prepare URL for API reading repo issues directly by full address and
	// issue number.
	case f&cREAD > 0:
		addr.HTTP = "GET"
		str := "please specify owner, repository and issue number"
		c, addr.URL, err = setURLAddress(c, addr.URL, str)
		if err != nil {
			return c, addr, fmt.Errorf("READ: %v", err)
		}
		addr.URL += "/" + c.Number

	// Prepare for editing a preexisting repo.
	case f&cEDIT > 0:
		addr.HTTP = "PATCH"
		str := "please specify owner, repository and issue number"
		c, addr.URL, err = setURLAddress(c, addr.URL, str)
		if err != nil {
			return c, addr, fmt.Errorf("EDIT: %v", err)
		}
		addr.URL += "/" + c.Number

	// Prepare a URL to set the current issue status to resolved, requires
	// login.
	case f&cLOCK > 0:
		addr.HTTP = "PUT"
		str := "please specify owner, repository and issue number"
		c, addr.URL, err = setURLAddress(c, addr.URL, str)
		if err != nil {
			return c, addr, fmt.Errorf("LOCK: %v", err)
		}
		addr.URL += "/" + c.Number + "/lock"

	// Prepare URL to set the given issue status to open, requires login.
	case f&cUNLOCK > 0:
		addr.HTTP = "DELETE"
		str := "please specify owner, repository and issue number"
		c, addr.URL, err = setURLAddress(c, addr.URL, str)
		if err != nil {
			return c, addr, fmt.Errorf("UNLOCK: %v", err)
		}
		addr.URL += "/" + c.Number + "/lock"

	// Prepare URL for issue creation by way of a complete issue address
	// and the use of the POST function, requires login authorisation.
	case f&cRAISE > 0:
		addr.HTTP = "POST"
		str := "please specify owner and repository details"
		c, addr.URL, err = setURLAddress(c, addr.URL, str)
		if err != nil {
			return c, addr, fmt.Errorf("RAISE: %v", err)
		}
	}

	// Add queries to url.
	if len(c.Queries) > 0 && f&cLOCK == 0 {
		q := url.QueryEscape(strings.Join(c.Queries, " "))
		addr.URL = addr.URL + "?q=" + q
	}

	// If lock required, add query.
	// if f&cREASON > 0 {
	// 	addr.URL = addr.URL + "?lock_reason=" + c.Reason
	// }

	// If verbose flag is set print the address used.
	if f&cVERBOSE > 0 {
		fmt.Printf("Setting URL: %v %v\n", addr.HTTP, addr.URL)
	}

	return c, addr, err
}
