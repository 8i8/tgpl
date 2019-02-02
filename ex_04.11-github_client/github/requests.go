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
	"net/http"
	"net/url"
	"strings"
)

const (
	NONE = iota
	ORG
	USER
	LOGIN
)

var state int = NONE

// Define a state by which to run the program, from the selection of possible
// uses infered frim available information.
func setState(conf Config) int {
	if conf.Name != "" {
		return USER
	} else if conf.Org != "" {
		return ORG
	} else if conf.Login != "" {
		return LOGIN
	}
	return NONE
}

// Structure an https request from the availabe data given.
func setUrl(conf Config) (string, string) {

	var HTTP string
	URL := "https://api.github.com/"

	switch conf.Mode {
	case "list":
		HTTP = "GET"
		if conf.Name != "" && conf.Repo == "" {
			URL = URL + conf.Name + "/issues"
		} else if conf.Name != "" && conf.Repo != "" {
			URL = URL + "repos/" + conf.Name + "/" + conf.Repo + "/issues"
		} else if conf.Name != "" {
			URL = URL + "repos/" + conf.Name + "/issues"
		} else if conf.Org != "" && conf.Repo != "" {
			URL = URL + "orgs/" + conf.Org + "/" + conf.Repo + "/issues"
		} else if conf.Org != "" {
			URL = URL + "orgs/" + conf.Org + "/issues"
		} else if conf.Login != "" && conf.Repo != "" {
			URL = URL + "repos/" + conf.Login + "/" + conf.Repo + "/issues"
		} else if conf.Login != "" {
			URL = URL + "repos/" + conf.Login + "/issues"
		} else {
			URL = URL + "search/issues"
		}
	case "read":
		HTTP = "GET"
		URL = URL + "search/issues"
	}

	// Add queries to url.
	if len(conf.Queries) > 0 {
		q := url.QueryEscape(strings.Join(conf.Queries, " "))
		URL = URL + "?q=" + q
	}

	return HTTP, URL
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(conf Config) (*IssuesSearchResult, error) {

	state = setState(conf)
	HTTP, URL := setUrl(conf)

	// Genereate request.
	req, err := http.NewRequest(HTTP, URL, nil)
	if err != nil {
		Log.Printf("error: %v", err.Error())
		return nil, err
	}

	// Add header to request.
	req.Header.Set(
		"Accept", "application/vnd.github.v3.text-match+json")
	if conf.Token != "" {
		req.Header.Set("Authorization", "token "+conf.Token)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		Log.Printf("error: %v: http response: %v", err.Error(), resp.StatusCode)
		return nil, err
	}

	// Close resp.Body.
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		Log.Printf("http response: %v", resp.StatusCode)
		return nil, err
	}

	// Decode reply.
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		Log.Printf("error: %v", err.Error())
		return &result, err
	}

	resp.Body.Close()
	return &result, nil
}

// Generate a new issue.
func RaiseIssue(data []Issue, conf Config) {

	HTTP, URL := setUrl(conf)

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

func EditIssue() {
	//cmd := exec.Command(c.Editor, "temp.md")
	// if err := cmd.Run(); err != nil {
	// 	_, file, line, _ := runtime.Caller(0)
	// 	msg := "Failed to open tempory file"
	// 	fmt.Fprintf(os.Stderr, "error: %s: %v %v: %s\n", msg, file, line, err.Error())
	// 	fmt.Println(c.Editor)
	// }
}
