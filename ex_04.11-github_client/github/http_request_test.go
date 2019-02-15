package github

import (
	"encoding/json"
	"net/http"
	"testing"
)

const tIssuesAddrURL = "https://api.github.com/repos/8i8/test/issues"
const tIssuesQueryURL = "https://api.github.com/search/issues?q=repo:8i8/test"

var err error

func TestSearchIssuesQuery(t *testing.T) {
	if err := searchIssuesQueryTest("GET", tIssuesQueryURL); err != nil {
		t.Errorf("error: %v", err.Error())
	}
}

func TestSearchIssuesAddr(t *testing.T) {
	if err := searchIssuesAddrTest("GET", tIssuesAddrURL); err != nil {
		t.Errorf("error: %v", err.Error())
	}
}

func BenchmarkSearchIssuesQuery(b *testing.B) {
	if err = searchIssuesQueryTest("GET", tIssuesQueryURL); err != nil {
		errlog.Printf("error: %v", err.Error())
	}
}

func BenchmarkSearchIssuesAddr(b *testing.B) {
	if err = searchIssuesAddrTest("GET", tIssuesAddrURL); err != nil {
		errlog.Printf("error: %v", err.Error())
	}
}

// SearchIssues queries the GitHub issue tracker.
func searchIssuesQueryTest(HTTP, URL string) error {

	// Genereate request.
	req, err := http.NewRequest(HTTP, URL, nil)
	if err != nil {
		errlog.Printf("error: %v", err.Error())
		return err
	}

	// Add header to request.
	req.Header.Set(
		"Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		errlog.Printf("error: %v", err.Error())
		return err
	}

	// Close resp.Body.
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		errlog.Printf("error: %v", resp.Status)
		return err
	}

	// decode json from within the file.
	var result IssuesSearchResult

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		errlog.Printf("error: %v", err.Error())
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
		errlog.Printf("error: %v", err.Error())
		return err
	}

	// Add header to request.
	req.Header.Set(
		"Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		errlog.Printf("error: %v", err.Error())
		return err
	}

	// Close resp.Body.
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		errlog.Printf("error: %v", resp.Status)
		return err
	}

	// decode json from within the file.
	var result []*Issue

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		errlog.Printf("error: %v", err.Error())
		return err
	}
	resp.Body.Close()
	return nil
}
