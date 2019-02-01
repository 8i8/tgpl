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
	"testing"
)

const t_IssuesAddrURL = "https://api.github.com/repos/8i8/test/issues"
const t_IssuesQueryURL = "https://api.github.com/search/issues?q=repo:8i8/test"

func TestSearchIssuesQuery(t *testing.T) {
	searchIssuesQueryTest(t, "GET", t_IssuesQueryURL)
}

func TestSearchIssuesAddr(t *testing.T) {
	searchIssuesAddrTest(t, "GET", t_IssuesAddrURL)
}

// SearchIssues queries the GitHub issue tracker.
func searchIssuesQueryTest(t *testing.T, HTTP, URL string) error {

	// Genereate request.
	req, err := http.NewRequest(HTTP, URL, nil)
	if err != nil {
		t.Errorf("error: %v\n", err.Error())
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		return err
	}

	// Add header to request.
	req.Header.Set(
		//"Accept", "application/vnd.github.v3.text-match+json")
		"Accept", "application/vnd.github.v3+json")
	//"Accept", "application/json")
	//req.Header.Set("Authorization", "token "+conf.Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("error: %v\n", err.Error())
		return err
	}

	// Close resp.Body.
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		t.Errorf("error: %v", resp.Status)
		return err
	}

	// decode json from within the file.
	var result IssuesSearchResult

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		t.Errorf("error: %v\n", err.Error())
		return err
	}
	resp.Body.Close()
	return nil
}

// SearchIssues queries the GitHub issue tracker.
func searchIssuesAddrTest(t *testing.T, HTTP, URL string) error {

	// Genereate request.
	req, err := http.NewRequest(HTTP, URL, nil)
	if err != nil {
		t.Errorf("error: %v\n", err.Error())
		return err
	}

	// Add header to request.
	req.Header.Set(
		//"Accept", "application/vnd.github.v3.text-match+json")
		"Accept", "application/vnd.github.v3+json")
	//"Accept", "application/json")
	//req.Header.Set("Authorization", "token "+conf.Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("error: %v\n", err.Error())
		return err
	}

	// Close resp.Body.
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		t.Errorf("error: %v", resp.Status)
		return err
	}

	// decode json from within the file.
	var result []*Issue

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		t.Errorf("error: %v\n", err.Error())
		return err
	}
	resp.Body.Close()
	return nil
}

func BenchmarkSearchIssuesQuery(b *testing.B) {
	for i := 0; i < 3; i++ {
		searchIssuesQueryTest(nil, "GET", t_IssuesQueryURL)
	}
}

func BenchmarkSearchIssuesAddr(b *testing.B) {
	for i := 0; i < 3; i++ {
		searchIssuesAddrTest(nil, "GET", t_IssuesAddrURL)
	}
}
