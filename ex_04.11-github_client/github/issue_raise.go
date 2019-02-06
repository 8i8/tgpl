package github

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Raise a new issue on a github repository.
func RaiseIssue(conf Config) {

	// Get user input.
	json, err := writeIssue(conf)
	if err != nil {
		Log.Printf("error: %v", err.Error())
		return
	}

	// Make http request.
	raiseIssue(conf, json)
}

// Compose issues for the designated repo.
func writeIssue(conf Config) (*bytes.Buffer, error) {

	var err error
	reader := bufio.NewReader(os.Stdin)

	// Get insist that title be entered.
	var title string
	for title == "" {
		fmt.Print("title: ")
		title, _ = reader.ReadString('\n')
		title = strings.Replace(title, "\n", "", -1)
		if title == "" {
			fmt.Print("Please provide a title.\n")
		}
	}

	// Require that a message body be entered, open an editor if desired.
	var body string
	if len(conf.Editor) > 0 {
		body, err = editorWrite(conf)
		if err != nil {
			return nil, err
		}
	} else {
		for body == "" {
			fmt.Print("message: ")
			body, _ = reader.ReadString('\n')
			body = strings.Replace(body, "\n", "", -1)
			if body == "" {
				fmt.Print("Please provide some details.\n")
			}
		}
	}

	// If either field are empty, refuse to generate.
	if len(title) == 0 || len(body) == 0 {
		err = errors.New("Issue fields are empty, refusing to send request.")
		return nil, err
	}

	str := `{"title":"` + title + `","body":"` + body + `"}`
	json := bytes.NewBufferString(str)

	return json, err
}
