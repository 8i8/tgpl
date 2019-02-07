package github

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// Get user input by way of the flag designated editor.
func openEditor(conf Config) (string, error) {

	// Creeate a tempory file for text editing.
	tmpfile, err := ioutil.TempFile("", "issue.*.md")
	if err != nil {
		return "", fmt.Errorf("TempFIle: %v", err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	// Open tempory file in the given editor.
	cmd := exec.Command(conf.Editor, tmpfile.Name())
	_, err = cmd.Output()
	if err != nil {
		Log.Fatal(err)
	}

	// Read temp file back in.
	var body string
	content, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		return body, fmt.Errorf("File read failed: %+v", err.Error())
	}
	body = string(content)
	body = strings.TrimSpace(body)

	return body, err
}

func openEditorBody(conf Config, body string) (string, error) {
	var err error
	return body, err
}
