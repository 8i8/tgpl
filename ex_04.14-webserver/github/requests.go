package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Bucket struct {
	Issues []Issue
	Page   int
	Err    error
	Count  int
}

var baseURL = "https://api.github.com/repos/"

// GetAllIssues returns a struct containing all of issues for the repo.
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

// GoGetAllIssues returns a struct containing all of issues for the repo.
func GoGetAllIssues(name, repo, token string, bucket Bucket, ch chan<- Bucket) {

	if VERBOSE {
		fmt.Printf("                                                                                \r")
		fmt.Printf(" GetAllIssues: requesting data from github\r")
	}

	// Set the url and header, make request.
	url := baseURL + name + "/" + repo + "/issues?state=all&page=" +
		strconv.Itoa(bucket.Page) + "&per_page=100"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		bucket.Err = err
		fmt.Println("error: http.NewRequest")
		ch <- bucket
		return
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.Header.Set("Authorization", "token "+token)
	resp, err := http.DefaultClient.Do(req)

	// Print status before dealing with errors.
	if VERBOSE {
		fmt.Printf("                                                                                \r")
		// if resp.StatusCode == 200 {
		// 	fmt.Printf(" StatusCode: %v\r", resp.Status)
		// } else {
		fmt.Printf(" StatusCode: %v: page: %d tries: %d\n", resp.Status, bucket.Page, bucket.Count)
		// }
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		if resp.StatusCode > 200 && resp.StatusCode < 400 {
			bucket.Err = err
			ch <- bucket
			return
		}
		bucket.Err = err
		fmt.Printf("error: resp.StatusCode: %v", resp.Status)
		ch <- bucket
		return
	}
	if err != nil {
		bucket.Err = err
		ch <- bucket
		return
	}

	// Load structs from buckets
	if err := json.NewDecoder(resp.Body).Decode(&bucket.Issues); err != nil {
		resp.Body.Close()
		fmt.Println("error: json.NewDecoder")
		ch <- bucket
		return
	}
	resp.Body.Close()

	if VERBOSE {
		fmt.Printf("                                                                                \r")
		fmt.Printf(" GetAllIssues: data downloaded\r")
	}

	ch <- bucket
	return
}
