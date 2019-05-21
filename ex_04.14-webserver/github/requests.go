package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var baseURL = "https://api.github.com/repos/"

// GetAllIssues returns a struct containg all of issues for the repo.
func GetAllIssues(name, repo, token string) ([]Issue, error) {

	if VERBOSE {
		fmt.Printf("GetAllIssues: requesting data from github\n")
	}

	// Set the url and make a new html get request,
	url := baseURL + name + "/" + repo + "/issues?state=all&page=10&per_page=100"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("NewRequest: %v", err)
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.Header.Set("Authorization", "token "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("DefaultClient.Do: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("StatusCode: %v", resp.Status)
	}

	if VERBOSE {
		fmt.Printf("StatusCode: %d %v\n", resp.StatusCode, resp.Status)
	}

	// Load structs from results
	var result []Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, fmt.Errorf("Decode: %v", err)
	}
	resp.Body.Close()

	if VERBOSE {
		fmt.Printf("GetAllIssues: data downloaded\n")
	}

	return result, nil
}
