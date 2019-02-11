package gitish

import "fmt"

// ReadIssue formulates an http request using that POST method on the specified
// git repository.
func ReadIssue(conf Config) error {

	// Serch for requested issue.
	issue, err := searchIssues(conf)
	if err != nil {
		return fmt.Errorf("read: %v", err)
	}

	// Display.
	printIssue(*issue[0])

	return err
}
