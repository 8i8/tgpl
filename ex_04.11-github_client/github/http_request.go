package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Header contains header request key paires.
type Header struct {
	Key, Value string
}

// Status contains an http responce status, both the text and the code.
type Status struct {
	Code    int
	Message string
}

// Address contains the request tupe and URL of a request.
type Address struct {
	HTTP, URL string
	header    []Header
	Status
}

// sendRequest compiles and sends the appropriate predefined request.
func sendRequest(conf Config, addr Address, json io.Reader) (*http.Response, error) {

	// Get a new request object.
	req, err := http.NewRequest(addr.HTTP, addr.URL, json)
	if err != nil {
		return nil, fmt.Errorf("NewRequest: %v", err)
	}

	// Load the key value paires into the request.
	for _, h := range addr.header {
		req.Header.Set(h.Key, h.Value)
	}

	// Make the request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("DefaultClient: %v", err)
	}
	return resp, err
}

// getStatus fills a Status struct with an http responce status data and
// verifies that responce data agains the expected responce, raising an error
// when required.
func getStatus(resp *http.Response) (Status, error) {

	var s Status
	s.Code = resp.StatusCode
	s.Message = http.StatusText(resp.StatusCode)

	if (rState == rLone || rState == rMany) && s.Code != http.StatusOK {
		// If data recieved.
		resp.Body.Close()
		return s, fmt.Errorf("response: %v %v", s.Code, s.Message)
		// IfRecord was created.
	} else if rState == rNone && s.Code != http.StatusCreated {
		resp.Body.Close()
		return s, fmt.Errorf("response: %v %v", s.Code, s.Message)
	}

	return s, nil
}

// respDecode decodes an http responce dependant on the expected responce
// state, into a single issue or an array of issues as required.
func respDecode(c Config, resp *http.Response) (interface{}, error) {

	var err error

	if c.Verbose {
		fmt.Println("respDecode: attempting decode")
	}

	// Decode into either an issue struct or an array of issue structs.
	if rState == rLone {

		// Single issue.
		var issue Issue
		if err = json.NewDecoder(resp.Body).Decode(&issue); err != nil {
			return nil, fmt.Errorf("json decoder failed: %v", err)
		}
		if c.Verbose {
			fmt.Println("respDecode: single issue decode")
		}
		return issue, err

	} else if rState == rMany {

		// Array of issues.
		var issue IssuesSearchResult
		if err = json.NewDecoder(resp.Body).Decode(&issue); err != nil {
			return nil, fmt.Errorf("json decoder failed: %v", err)
		}
		if c.Verbose {
			fmt.Println("respDecode: multiple issue decode")
		}
		result := issue.Items
		return result, err
	}

	return nil, nil
}

// makeRequest orchestrates an http request.
func makeRequest(conf Config, json io.Reader) (interface{}, error) {

	// Set the correct url for the request.
	addr, err := setURL(conf)
	if err != nil {
		return nil, fmt.Errorf("setURL failed: %v", err)
	}

	// Compose an array of header key value paires.
	addr.header, err = composeHeader(conf)
	if err != nil {
		return nil, fmt.Errorf("composeHeader: %v", err)
	}

	// Make and send the request.
	resp, err := sendRequest(conf, addr, json)
	if err != nil {
		return nil, fmt.Errorf("sendRequest: %v", err)
	}

	// Record the responce return status.
	addr.Status, err = getStatus(resp)
	if err != nil {
		return nil, fmt.Errorf("getStatus: %v", err)
	}

	// Decode the responce.
	result, err := respDecode(conf, resp)
	if err != nil {
		return nil, fmt.Errorf("respDecode: %v", err)
	}

	resp.Body.Close()
	return result, err
}
