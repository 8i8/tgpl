// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

package github

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func RaiseIssue() {

	str := `{"title": "The Go Issue","body": "This is the thing you see, it came from a Go application."}`
	data := bytes.NewBufferString(str)

	// Formulate post request
	req, err := http.NewRequest("POST", IssuesPostURL, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "POST failed : %s\n", err)
	}

	// Set header.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "token "+Config.Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "HEAD failed : %s\n", resp.Status)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "Issue creation failed: %s\n", resp.Status)
	}

	resp.Body.Close()
}
