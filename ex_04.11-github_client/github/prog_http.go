package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Header struct {
	Key, Value string
}

type Status struct {
	Code   int
	Reason string
}

func WriteResponce() {
}

// SearchIssues queries the GitHub issue tracker.
func searchIssues(conf Config) ([]*Issue, error) {

	// Set the appropriate URL.
	HTTP, URL, err := setURL(conf)
	if err != nil {
		return nil, fmt.Errorf("serUrl failed: %v", err)
	}

	// Generate request.
	req, err := http.NewRequest(HTTP, URL, nil)
	if err != nil {
		return nil, fmt.Errorf("http NewRequest failed: %v", err)
	}

	// Add header to request.
	req.Header.Set(
		"Accept", "application/vnd.github.v3.text-match+json")
	if conf.Token != "" {
		req.Header.Set("Authorization", "token "+conf.Token)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http responce failed: %v", err)
	}

	// Close without decoding if not ok.
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("http response: %d %v", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	// Decode reply urlAddr for a direct http request and urlSear using the
	// API search function.
	var result []*Issue
	if state == urlAddr {
		var issue Issue
		if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
			return nil, fmt.Errorf("json decoder failed: %v", err)
		}
		result = append(result, &issue)
	} else if state == urlSear {
		var issue IssuesSearchResult
		if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
			return nil, fmt.Errorf("json decoder failed: %v", err)
		}
		result = issue.Items
	}

	resp.Body.Close()
	return result, err
}

// Generate a new issue.
func raiseIssue(conf Config, json *bytes.Buffer) error {

	// Set the appropriate URL.
	HTTP, URL, err := setURL(conf)
	if err != nil {
		return fmt.Errorf("raiseIssue : %v", err)
	}

	// Formulate post request
	req, err := http.NewRequest(HTTP, URL, json)
	if err != nil {
		return fmt.Errorf("raiseIssue: %v", err)
	}

	// Set header.
	req.Header.Set("Accept", "application/vnd.github.v3.json")
	req.Header.Set("Authorization", "token "+conf.Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("raiseIssue: %v", err)
	}

	// If response not successful report it.
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("http response: %d %v", resp.StatusCode,
			http.StatusText(resp.StatusCode))
	}

	resp.Body.Close()
	return err
}

// overwrite the existing issue.
func writeIssue(conf Config, json *bytes.Buffer) error {

	// Set the appropriate URL.
	HTTP, URL, err := setURL(conf)
	if err != nil {
		return fmt.Errorf("setURL: %v", err)
	}

	// Formulate post request
	req, err := http.NewRequest(HTTP, URL, json)
	if err != nil {
		return fmt.Errorf("NewRequest: %v", err)
	}

	// Set header.
	req.Header.Set("Accept", "application/vnd.github.v3.json")
	req.Header.Set("Authorization", "token "+conf.Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("http request: %v", err)
	}

	// If response not successful report it.
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("http response: %d %v", resp.StatusCode,
			http.StatusText(resp.StatusCode))
	}

	resp.Body.Close()
	return err
}
