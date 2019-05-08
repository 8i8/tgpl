package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// respDecode decodes an http response, in accord with program state.
func respDecode(c Config, resp *http.Response) (Reply, error) {

	var reply Reply
	var msg json.RawMessage
	var err error
	var log string

	if f&cVERBOSE > 0 {
		fmt.Println("respDecode: decoding http response")
	}

	// Decode response into a raw holding variable, stored as a map of key
	// value pairs.
	if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
		return reply, fmt.Errorf("json decoder body failed: %v", err)
	}

	// Apply the decoding accorging to state.
	switch {
	case f&cLIST > 0:
		// Decode multiple issues, place in interface.
		var issue IssuesSearchResult
		if err := json.Unmarshal(msg, &issue); err != nil {
			return reply, fmt.Errorf("json decoder Msg failed: %v", err)
		}
		reply.Msg = issue
		log = "multiple"
	case f&cREAD > 0:
		// Decode a single issue, place	in interface.
		var issue Issue
		if err := json.Unmarshal(msg, &issue); err != nil {
			return reply, fmt.Errorf("json decoder Msg failed: %v", err)
		}
		reply.Msg = issue
		log = "single"
	default:
		log = "empty"
	}

	if f&cVERBOSE > 0 {
		fmt.Println("respDecode: response successfuly read", log)
	}

	return reply, err
}

// issueToJSON marshals data into json format and returns it as a byte slice.
func issueToJSON(title, body string) ([]byte, error) {

	// Write data into a struct.
	var issue Issue
	issue.Title = title
	issue.Body = body

	// Marshal the struct
	json, err := json.Marshal(issue)
	if err != nil {
		return nil, fmt.Errorf("Marshal: %v", err)
	}

	return json, err
}

// lockReasonJSON prepares the json to lock an issue and also add the given
// reason for locking.
func lockReasonJSON(reason string) ([]byte, error) {

	// Reduced struct for locking process.
	type localIssue struct {
		Locked bool   `json:"locked"`
		Reason string `json:"active_lock_reason"`
	}

	// Write data into a struct.
	var issue localIssue
	issue.Locked = true
	issue.Reason = reason

	// Marshal the struct
	json, err := json.Marshal(issue)
	if err != nil {
		return nil, fmt.Errorf("Marshal: %v", err)
	}

	return json, err
}
