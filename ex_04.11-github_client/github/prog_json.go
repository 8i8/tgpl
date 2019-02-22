package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// respDecode decodes an http response dependant upon the expected response
// state, into a single issue or an array of issues as required.
func respDecode(c Config, resp *http.Response) (Reply, error) {

	var reply Reply
	var msg json.RawMessage
	var err error
	var log string

	if c.Verbose {
		fmt.Println("respDecode: attempting decode")
	}

	// Decode the response into a raw holding variable, here it is stored
	// as a map of key value pairs.
	if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
		return reply, fmt.Errorf("json decoder body failed: %v", err)
	}

	// Set the Type of struct
	reply.Type = rState

	// Set the final decoding type dependant on the program state.
	switch rState {
	case rMANY:
		// Decode multiple issues, place into the envelope structs interface.
		var issue IssuesSearchResult
		if err := json.Unmarshal(msg, &issue); err != nil {
			return reply, fmt.Errorf("json decoder Msg failed: %v", err)
		}
		reply.Msg = issue
		log = "multiple"
	case rLONE:
		// Decode a single issue, place	into the envelope struct interface.
		var issue Issue
		if err := json.Unmarshal(msg, &issue); err != nil {
			return reply, fmt.Errorf("json decoder Msg failed: %v", err)
		}
		reply.Msg = issue
		log = "single"
	case rRAW:
		// Decode multiple issues, place into the envelope structs interface.
		var issues IssuesSearchResult
		err := json.Unmarshal(msg, &issues)
		if err != nil {
			// Decode a single issue, place	into the envelope struct interface.
			var issue Issue
			if err := json.Unmarshal(msg, &issue); err != nil {
				return reply, fmt.Errorf("json decoder Raw failed: %v", err)
			}
			reply.Msg = issue
			reply.Type = rLONE
			log = "raw single"
			break
		}
		reply.Msg = issues
		reply.Type = rMANY
		log = "raw multiple"
	default:
		log = "empty"
	}

	if c.Verbose {
		fmt.Println("respDecode: attempt", log)
	}

	return reply, err
}

// issueToJSON marshals data into json format and returns it in a bytes buffer.
func issueToJSON(title, body string) (*bytes.Buffer, error) {

	// Write data into a struct.
	var issue Issue
	issue.Title = title
	issue.Body = body

	// Marshal the struct
	json, err := json.Marshal(issue)
	if err != nil {
		return nil, fmt.Errorf("Marshal: %v", err)
	}

	// Write into a byte buffer.
	var b bytes.Buffer
	b.Write(json)

	return &b, err
}
