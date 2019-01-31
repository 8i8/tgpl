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
	"os"
	"strings"
)

const (
	NONE = iota
	ORG
	USER
	LOGIN
)

var URL string
var state int = NONE

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

func setUrl(conf Config) string {

	str := "https://api.github.com/"

	switch conf.Mode {
	case "list":
		if conf.Name != "" && conf.Repo == "" {
			str = str + conf.Name + "/issues"
		} else if conf.Name != "" && conf.Repo != "" {
			str = str + "repos/" + conf.Name + "/" + conf.Repo + "/issues"
		} else if conf.Name != "" {
			str = str + "repos/" + conf.Name + "/issues"
		} else if conf.Org != "" && conf.Repo != "" {
			str = str + "orgs/" + conf.Org + "/" + conf.Repo + "/issues"
		} else if conf.Org != "" {
			str = str + "orgs/" + conf.Org + "/issues"
		} else if conf.Login != "" && conf.Repo != "" {
			str = str + "repos/" + conf.Login + "/" + conf.Repo + "/issues"
		} else if conf.Login != "" {
			str = str + "repos/" + conf.Login + "/issues"
		} else {
			str = str + "issues"
		}
	}

	return str
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(conf Config) (*IssuesSearchResult, error) {

	state = setState(conf)
	URL = setUrl(conf)

	q := url.QueryEscape(strings.Join(conf.Queries, " "))

	// Genereate request.
	// https://api.github.com/search/issues?q=repo%3Agolang%2Fgo+json+decoder
	// req, err := http.NewRequest("GET", "https://api.github.com/search/issues?q=repo%3Agolang%2Fgo+json+decoder", nil)
	req, err := http.NewRequest("GET", URL+"?q="+q, nil)
	fmt.Println("GET", URL+"?q="+q)
	if err != nil {
		msg := "New Request failed."
		fmt.Fprintf(os.Stderr, "error: %s: %s\n", msg, err.Error())
		return nil, err
	}

	fmt.Println(req)

	// Add header to request.
	req.Header.Set(
		"Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		msg := "Header failed."
		fmt.Fprintf(os.Stderr, "error: %s: %s\n", msg, err.Error())
		fmt.Println(resp.StatusCode)
		return nil, err
	}

	// Close resp.Body.
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	fmt.Println(resp)

	// Decode reply.
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		msg := "json decoder failed"
		fmt.Fprintf(os.Stderr, "error: %s: %s\n", msg, err.Error())
		return &result, err
	}
	resp.Body.Close()
	return &result, nil
}

// func GetIssue(num int) Issue {
// }

func RaiseIssue(data []Issue, conf Config) {

	URL = setUrl(conf)

	//str := `{"title": "The Go Issue","body": "This is the thing you see, it came from a Go application."}`
	//json := bytes.NewBufferString(str)
	str, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Json Marshal failed : %s\n", err)
	}
	json := bytes.NewBuffer(str)

	// Formulate post request
	req, err := http.NewRequest("POST", URL, json)
	if err != nil {
		fmt.Fprintf(os.Stderr, "POST failed : %s\n", err)
	}

	// Set header.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "token "+conf.Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "HEAD failed : %s\n", resp.Status)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "Issue creation failed: %s\n", resp.Status)
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
