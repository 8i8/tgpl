package github

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

// Raise a new issue on a github repository.
func RaiseIssue(conf Config) error {

	// Fill io.Buffer with json data.
	var json *bytes.Buffer
	var err error
	json, err = composeIssue(conf)
	if err != nil {
		return fmt.Errorf("composeIssue: %v", err)
	}

	// Make http request.
	err = raiseIssue(conf, json)
	return err
}

// Compose issues for the designated repo.
func composeIssue(conf Config) (*bytes.Buffer, error) {

	var err error
	sc := bufio.NewScanner(os.Stdin)

	// Insist that a title be entered.
	var title string
	for title == "" && err == nil {
		fmt.Print("title: ")
		sc.Scan()
		title += sc.Text()
		title = strings.TrimSpace(title)
		err = sc.Err()
		if title == "" && err == nil {
			fmt.Print("Please provide a title.\n")
		}
	}
	if err != nil {
		return nil, fmt.Errorf("title scanner: %+v", err)
	}

	// Require that a message body be entered, open an editor if desired.
	var body string
	if len(conf.Editor) > 0 {
		body, err = openEditor(conf)
	} else {
		fmt.Println("An empty line after the message will end the message.")
		for body == "" {
			fmt.Print("message: ")
			for sc.Scan() {
				line := sc.Text()
				if len(line) == 0 {
					break
				}
				body += line
			}
			body = strings.TrimSpace(body)
			err = sc.Err()
			if body == "" && err == nil {
				fmt.Println("Please provide some details.")
			}
		}
	}
	if err != nil {
		return nil, fmt.Errorf("message scanner: %+v", err)
	}

	str := `{"title":"` + title + `","body":"` + body + `"}`
	json := bytes.NewBufferString(str)

	return json, err
}
