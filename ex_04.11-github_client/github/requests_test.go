// GET https://api.github.com/repos/golang/go/issues?q=json+decoder

// var query1 []string
// var query2 []string

// func init() {
// 	query1 = append(query1, "is:open")
// 	query2 = append(query2, "repo:golang/go")
// 	query2 = append(query2, "is:open")
// }

//URL := "https://api.github.com/users/octocat/orgs"
//URL := "https://api.github.com/orgs/octokit/repos"
//URL := "https://api.github.com/search/issues?q=repo:8i8/test"
//URL := "https://api.github.com/repos/8i8/test/issues"
//URL := "https://api.github.com/repos/8i8/test/issues"

//URL := "https://api.github.com/users/octocat/orgs"
//URL := "https://api.github.com/orgs/octokit/repos"
//URL := "https://api.github.com/search/issues?q=repo:8i8/test"
//URL := "https://api.github.com/repos/8i8/test/issues"
//URL := "https://api.github.com/repos/8i8/test/issues"

// var query []string
// query = append(query, "repo:8i8/test")
// query = append(query, "is:open")
// q := url.QueryEscape(strings.Join(terms, " "))

package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"testing"
)

const t_IssuesAddrURL = "https://api.github.com/repos/8i8/test/issues"
const t_IssuesQueryURL = "https://api.github.com/search/issues?q=repo:8i8/test"

var err error

func TestSearchIssuesQuery(t *testing.T) {
	if err := searchIssuesQueryTest("GET", t_IssuesQueryURL); err != nil {
		t.Errorf("error: %v", err.Error())
	}
}

func TestSearchIssuesAddr(t *testing.T) {
	if err := searchIssuesAddrTest("GET", t_IssuesAddrURL); err != nil {
		t.Errorf("error: %v", err.Error())
	}
}

func BenchmarkSearchIssuesQuery(b *testing.B) {
	if err = searchIssuesQueryTest("GET", t_IssuesQueryURL); err != nil {
		Log.Printf("error: %v", err.Error())
	}
}

func BenchmarkSearchIssuesAddr(b *testing.B) {
	if err = searchIssuesAddrTest("GET", t_IssuesAddrURL); err != nil {
		Log.Printf("error: %v", err.Error())
	}
}

// SearchIssues queries the GitHub issue tracker.
func searchIssuesQueryTest(HTTP, URL string) error {

	// Genereate request.
	req, err := http.NewRequest(HTTP, URL, nil)
	if err != nil {
		Log.Printf("error: %v", err.Error())
		return err
	}

	// Add header to request.
	req.Header.Set(
		"Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		Log.Printf("error: %v", err.Error())
		return err
	}

	// Close resp.Body.
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		Log.Printf("error: %v : status: %v", err.Error(), resp.Status)
		return err
	}

	// decode json from within the file.
	var result IssuesSearchResult

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		Log.Printf("error: %v", err.Error())
		return err
	}
	resp.Body.Close()
	return nil
}

// SearchIssues queries the GitHub issue tracker.
func searchIssuesAddrTest(HTTP, URL string) error {

	// Genereate request.
	req, err := http.NewRequest(HTTP, URL, nil)
	if err != nil {
		_, _, line, _ := runtime.Caller(0)
		fmt.Fprintf(os.Stderr, "error: %v: %s\n", line, err.Error())
		return err
	}

	// Add header to request.
	req.Header.Set(
		"Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		_, _, line, _ := runtime.Caller(0)
		fmt.Fprintf(os.Stderr, "error: %v: %s\n", line, err.Error())
		return err
	}

	// Close resp.Body.
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		_, _, line, _ := runtime.Caller(0)
		fmt.Fprintf(os.Stderr, "error: %v: %s\n", line, resp.Status)
		return err
	}

	// decode json from within the file.
	var result []*Issue

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		_, _, line, _ := runtime.Caller(0)
		fmt.Fprintf(os.Stderr, "error: %v: %s\n", line, err.Error())
		return err
	}
	resp.Body.Close()
	return nil
}
