package gitish

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Address struct {
	Http, Url string
}

// Header contains header request key paires.
type Header struct {
	Key, Value string
}

// Status contains an http responce status, both the text and the code.
type Status struct {
	Code    int
	Message string
}

func sendRequest(conf Config, addr Address) (*http.Response, error) {

	req, err := http.NewRequest(addr.Http, addr.Url, nil)
	if err != nil {
		return nil, fmt.Errorf("NewRequest: %v", err)
	}

	headers, err := composeHeader(conf)
	if err != nil {
		return nil, fmt.Errorf("composeHeader: %v", err)
	}

	for _, h := range headers {
		req.Header.Set(h.Key, h.Value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("DefaultClient: %v", err)
	}
	return resp, err
}

func getStatus(resp *http.Response) Status {

	var s Status
	s.Code = resp.StatusCode
	s.Reason = http.StatusText(resp.StatusCode)
	return s
}

func respDecode(resp *http.Response) ([]*Issue, error) {

	// Decode reply urlAddr for a direct http request and urlSear using the
	// API search function.
	var result []*Issue
	if state == respLone {
		var issue Issue
		if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
			return nil, fmt.Errorf("json decoder failed: %v", err)
		}
		result = append(result, &issue)

	} else if state == respMult {
		var issue IssuesSearchResult
		if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
			return nil, fmt.Errorf("json decoder failed: %v", err)
		}
		result = issue.Items
	}

	return result, nil
}

// MakeRequest orchistrates an http request.
func MakeRequest(conf Config) error {

	addr, err := setURL(conf)
	if err != nil {
		return fmt.Errorf("serUrl failed: %v", err)
	}

	resp, err := sendRequest(conf, addr)
	if err != nil {
		return err
	}

	stat := getStatus(resp)
	if stat.Code > 300 {
		return fmt.Errorf("error: %v", stat.Message)
	}

	resp.Body.Close()

	// Add functionality for counting the length of the returned array so
	// as to decide wheter or not to print an individual issue or a list.
	ListIssues(conf)

	return err
}
