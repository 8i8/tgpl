package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var baseURL = "https://api.github.com/repos/"

// GetAllIssues returns a struct containg all of issues for the repo.
func GetAllIssues(name, repo, token string, page int) ([]Issue, error) {

	if VERBOSE {
		fmt.Printf("GetAllIssues: requesting data from github\n")
	}

	// Set the url and make a new html get request,
	url := baseURL + name + "/" + repo + "/issues?state=all&page=" + strconv.Itoa(page) + "&per_page=100"
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
		if resp.StatusCode > 200 && resp.StatusCode < 400 {
			if VERBOSE {
				fmt.Printf("StatusCode: %v\n", resp.Status)
			}
			return nil, nil
		}
		return nil, fmt.Errorf("StatusCode: %v", resp.Status)
	}

	if VERBOSE {
		fmt.Printf("StatusCode: %v\n", resp.Status)
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

// GoGetAllIssues returns a struct containg all of issues for the repo.
func GoGetAllIssues(name, repo, token string, page int, ch chan<- []Issue) error {

	if VERBOSE {
		fmt.Printf("                                                                                \r")
		fmt.Printf(" GetAllIssues: requesting data from github\r")
	}

	var result []Issue

	// Set the url and make a new html get request,
	url := baseURL + name + "/" + repo + "/issues?state=all&page=" +
		strconv.Itoa(page) + "&per_page=100&since=2000-01-01T00:00:00+00:00"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		ch <- result
		return fmt.Errorf("NewRequest: %v", err)
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.Header.Set("Authorization", "token "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ch <- result
		return fmt.Errorf("DefaultClient.Do: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		if resp.StatusCode > 200 && resp.StatusCode < 400 {
			if VERBOSE {
				fmt.Printf("StatusCode: %v\n", resp.Status)
			}
			ch <- result
			return nil
		}
		ch <- result
		return fmt.Errorf("StatusCode: %v\n", resp.Status)
	}

	if VERBOSE {
		fmt.Printf("                                                                                \r")
		fmt.Printf(" StatusCode: %v\r", resp.Status)
	}

	// Load structs from results
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		ch <- result
		return fmt.Errorf("Decode: %v", err)
	}
	resp.Body.Close()

	if VERBOSE {
		fmt.Printf("                                                                                \r")
		fmt.Printf(" GetAllIssues: data downloaded\r")
	}

	ch <- result
	return nil
}
