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
	case rMany:
		// Decode multiple issues, place into the envelope structs interface.
		var issue IssuesSearchResult
		if err := json.Unmarshal(msg, &issue); err != nil {
			return reply, fmt.Errorf("json decoder Msg failed: %v", err)
		}
		reply.Msg = issue
	case rLone:
		// Decode a single issue, place	into the envelope struct interface.
		var issue Issue
		if err := json.Unmarshal(msg, &issue); err != nil {
			return reply, fmt.Errorf("json decoder Msg failed: %v", err)
		}
		reply.Msg = issue
	case rRaw:
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
			reply.Type = rLone
			break
		}
		reply.Msg = issues
		reply.Type = rMany
	}

	if c.Verbose {
		fmt.Println("respDecode: single issue decode")
	}

	return reply, err
}

// // respDecode decodes an http responce dependant upon the expected responce
// // state, into a single issue or an array of issues as required.
// func respDecode(c Config, resp *http.Response) (interface{}, error) {

// 	var err error
// 	if c.Verbose {
// 		fmt.Println("respDecode: attempting decode")
// 	}

// 	// Decode into either an issue struct or an array of issue structs.
// 	if rState == rLone {

// 		// Single issue.
// 		var issue Issue
// 		if err = json.NewDecoder(resp.Body).Decode(&issue); err != nil {
// 			return nil, fmt.Errorf("json decoder failed: %v", err)
// 		}
// 		if c.Verbose {
// 			fmt.Println("respDecode: single issue decode")
// 		}
// 		return issue, err

// 	} else if rState == rMany {

// 		// Array of issues.
// 		var issue IssuesSearchResult
// 		if err = json.NewDecoder(resp.Body).Decode(&issue); err != nil {
// 			return nil, fmt.Errorf("json decoder failed: %v", err)
// 		}
// 		if c.Verbose {
// 			fmt.Println("respDecode: multiple issue decode")
// 		}
// 		result := issue.Items
// 		return result, err
// 	}

// 	return nil, nil
// }

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
