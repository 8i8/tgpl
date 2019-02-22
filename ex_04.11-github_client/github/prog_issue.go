package github

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// Compose issues for the designated repo.
func composeIssue(conf Config) (*bytes.Buffer, error) {

	var err error
	sc := bufio.NewScanner(os.Stdin)

	// Insist that a title be entered.
	var title string
	for title == "" && err == nil {
		fmt.Print("title: ")
		sc.Scan()
		title = sc.Text()
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
		body, err = openInEditor(conf, body)
	} else {
		fmt.Println("An empty line after the message will end the message.")
		for body == "" {
			fmt.Print("message: ")
			for sc.Scan() {
				line := sc.Text()
				if len(line) == 0 {
					break
				}
				body += line + "\n"
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

	// Marshal into json format.
	json, err := issueToJSON(title, body)

	return json, err
}

// editIssue opens and edits and existing issue.
func editIssue(conf Config, reply Reply) (*bytes.Buffer, error) {

	var err error
	issue := reply.Msg.(Issue)
	// Retreive data from issue.
	currentTitle, currentBody := issue.Title, issue.Body

	// Display the current title and offer the occasion to change it.
	sc := bufio.NewScanner(os.Stdin)
	var title string
	for title == "" {
		fmt.Println("leave empty and press enter to keep the current title.")
		fmt.Print("Title: ", currentTitle, "\nNew title: ")
		sc.Scan()
		title = sc.Text()
		title = strings.TrimSpace(title)
		if title == "" {
			title = currentTitle
		}
	}
	if err != nil {
		return nil, fmt.Errorf("title scanner: %+v", err)
	}

	// Message body, open an editor if specified.
	var body string
	body, err = openInEditor(conf, currentBody)
	if err != nil {
		return nil, fmt.Errorf("openInEditor: %+v", err)
	}

	// Marshal into json format.
	json, err := issueToJSON(title, body)

	return json, nil
}

// Get user input by way of the flag designated editor.
func openInEditor(conf Config, text string) (string, error) {

	// Create a temporary file for text editing.
	tmpfile, err := ioutil.TempFile("", "issue.*.md")
	if err != nil {
		return text, fmt.Errorf("TempFile: %v", err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	// Write message text to file.
	if len(text) > 0 {
		_, err = tmpfile.Write([]byte(text))
		if err != nil {
			return text, fmt.Errorf("File write: %v", err)
		}
	}

	// Open temporary file in the given editor.
	cmd := exec.Command(conf.Editor, tmpfile.Name())
	_, err = cmd.Output()
	if err != nil {
		return text, fmt.Errorf("Editor Command: %v", err)
	}

	// Read temp file back in.
	content, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		return text, fmt.Errorf("File read failed: %+v", err.Error())
	}
	text = strings.TrimSpace(string(content))

	return text, err
}

// getPass requests that the user enter their password and then returns it
// encoded as a base64 string.
func getPass(c Config) (string, error) {

	// User input.
	sc := bufio.NewScanner(os.Stdin)
	var pass string
	var err error
	for pass == "" {
		fmt.Printf("Enter host password for user '%v': ", c.User)
		sc.Scan()
		pass = sc.Text()
	}
	if err != nil {
		return pass, fmt.Errorf("pass scanner: %+v", err)
	}
	// Encode to base64.
	pass = base64.StdEncoding.EncodeToString([]byte(pass))

	return pass, nil
}
