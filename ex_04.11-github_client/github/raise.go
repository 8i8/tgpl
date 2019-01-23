package github

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
)

func EditIssue() {

	// Make patch request
	req, err := http.NewRequest("PATCH", IssuesPostURL, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "POST failed : %s\n", err)
	}
}

// RaiseIssue raises an issue on the given github repo.
func RaiseIssue() {

	str := `{"title": "The Go Issue","body": "This is the thing you see, it came from a Go application."}`
	data := bytes.NewBufferString(str)

	// Make post request
	req, err := http.NewRequest("POST", IssuesPostURL, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "POST failed : %s\n", err)
	}

	req.Header.Set("Accept", "application/json")

	// token connect
	//req.Header.Set("Authorization", "token c537f1922712dc29a8fcc96cbf1b6119a7f4cb25")

	// standard password.
	encoded := base64.StdEncoding.EncodeToString([]byte("8i8:9b0c17e4c8821c9c508859c28fd30064"))
	req.Header.Set("Authorization", "basic "+encoded)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "HEAD failed : %s\n", resp.Status)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Fprintf(os.Stdout, "Return status: %s\n", resp.Status)
	}

	resp.Body.Close()
}
