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
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	NONE = iota
	ORG
	USER
	LOGIN
	SEARCH
	ADDRESS
)

var state int = NONE

// Check whether all requirements are met to enter ADDRESS mode.
func checkAddress(c Config) bool {
	return (len(c.Login) > 0 || len(c.Author) > 0 ||
		len(c.Owner) > 0 || len(c.Org) > 0) &&
		len(c.Repo) > 0
}

// Define a state by which to run the program, from the selection of possible
// uses inferred from available information.
func SetState(c *Config) int {

	// If there is a number given and all paramiters exist for a direct
	// HTTP access then do so, else add the number to the query listing and
	// search.
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
	}
	return state
}

// Structure an https request from the available data given.
func setUrl(conf Config) (string, string) {

	var HTTP string
	URL := "https://api.github.com/"

	switch conf.Mode {

	case "list":
		HTTP = "GET"
		URL = URL + "search/issues"
		if conf.Owner != "" && conf.Repo != "" {
			conf.Queries = append(conf.Queries, "repo:"+conf.Owner+"/"+conf.Repo)
		} else if conf.Owner != "" {
			conf.Queries = append(conf.Queries, "user:"+conf.Owner)
		} else if conf.Org != "" && conf.Repo != "" {
			conf.Queries = append(conf.Queries, "org:"+conf.Repo+"/"+conf.Repo)
		} else if conf.Org != "" {
			conf.Queries = append(conf.Queries, "org:"+conf.Org)
		} else if conf.Author != "" && conf.Repo != "" {
			conf.Queries = append(conf.Queries, "repo:"+conf.Author+"/"+conf.Repo)
		} else if conf.Author != "" {
			conf.Queries = append(conf.Queries, "author:"+conf.Author)
		} else if conf.Login != "" && conf.Repo != "" {
			conf.Queries = append(conf.Queries, "repo:"+conf.Login+"/"+conf.Repo)
		} else if conf.Login != "" {
			conf.Queries = append(conf.Queries, "author:"+conf.Login)
		}

	case "read":
		HTTP = "GET"
		if len(conf.Owner) > 0 && len(conf.Repo) > 0 && len(conf.Number) > 0 {
			URL = URL + "repos/" + conf.Owner + "/" + conf.Repo + "/issues/" + conf.Number
		} else if len(conf.Login) > 0 && len(conf.Repo) > 0 && len(conf.Number) > 0 {
			URL = URL + "repos/" + conf.Login + "/" + conf.Repo + "/issues"
		} else if len(conf.Author) > 0 && len(conf.Repo) > 0 && len(conf.Number) > 0 {
			URL = URL + "repos/" + conf.Author + "/" + conf.Repo + "/issues"
		} else if len(conf.Org) > 0 && len(conf.Repo) > 0 && len(conf.Number) > 0 {
			URL = URL + "orgs/" + conf.Org + "/" + conf.Repo + "/issues"
		} else {
			fmt.Println("Please provide owner, repo and number information.")
		}
	}

	// Add queries to url.
	if len(conf.Queries) > 0 {
		q := url.QueryEscape(strings.Join(conf.Queries, " "))
		URL = URL + "?q=" + q
	}

	return HTTP, URL
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(conf Config) ([]*Issue, error) {

	HTTP, URL := setUrl(conf)
	fmt.Println(HTTP, URL)

	// Generate request.
	req, err := http.NewRequest(HTTP, URL, nil)
	if err != nil {
		Log.Printf("error: %v", err.Error())
		return nil, err
	}

	// Add header to request.
	req.Header.Set(
		"Accept", "application/vnd.github.v3.text-match+json")
	//"Accept", "application/vnd.github.machine-man-preview")
	if conf.Token != "" {
		req.Header.Set("Authorization", "token "+conf.Token)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		Log.Printf("error: %v: http response: %v %v", err.Error(),
			resp.StatusCode, http.StatusText(resp.StatusCode))
		return nil, err
	}

	var result []*Issue

	// Close without decoding.
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("http response: %v %v\n", resp.StatusCode,
			http.StatusText(resp.StatusCode))
		resp.Body.Close()
		return result, err
	}

	// Decode reply ADDRESS for a direct http request and SEARCH using the
	// API search function.
	if state == ADDRESS {
		var issue Issue
		if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
			Log.Printf("error: %v", err.Error())
		}
		result = append(result, &issue)
	} else if state == SEARCH {
		var resultStruct IssuesSearchResult
		if err := json.NewDecoder(resp.Body).Decode(&resultStruct); err != nil {
			Log.Printf("error: %v", err.Error())
		}
		result = resultStruct.Items
	}

	resp.Body.Close()
	return result, err
}

// Generate a new issue.
func RaiseIssue(conf Config) {

	HTTP, URL := setUrl(conf)
	fmt.Println(HTTP, URL)

	data := writeIssue()
	str, err := json.Marshal(data)
	if err != nil {
		Log.Printf("error: %v", err.Error())
		return
	}
	json := bytes.NewBuffer(str)

	// Formulate post request
	req, err := http.NewRequest(HTTP, URL, json)
	if err != nil {
		Log.Printf("error: %v", err.Error())
		return
	}

	// Set header.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "token "+conf.Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		Log.Printf("error: %v", err.Error())
		return
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		Log.Printf("error: %v", err.Error())
		return
	}

	resp.Body.Close()
}

func writeIssue() Issue {
}

func EditIssue() {
	//cmd := exec.Command(c.Editor, "temp.md")
	// if err := cmd.Run(); err != nil {
	// 	_, file, line, _ := runtime.Caller(0)
	// 	msg := "Failed to open tempory file"
	// 	fmt.Fprintf(os.Stderr, "error: %s: %v %v: %s\n", msg, file, line, err.Error())
	// 	fmt.Println(c.Editor)
	// }
}
