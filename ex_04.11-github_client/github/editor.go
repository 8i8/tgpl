package github

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

// Get user input using the users editor of choice.
func editorWrite(conf Config) (string, error) {

	// Open the tempory file in the given editor.
	cmd := exec.Command(conf.Editor, "temp.md")
	_, err := cmd.Output()
	if err != nil {
		Log.Fatal(err)
	}

	// Read that file back in.
	f, err := os.Open("temp.md")
	input := bufio.NewScanner(f)
	var body string
	for input.Scan() {
		fmt.Printf("read `%v`\n", input.Text())
		body = body + input.Text()
	}
	f.Close()
	err = os.Remove("temp.md")
	if err != nil {
		Log.Print("Temp file removal failed.")
	}
	return body, err
}
