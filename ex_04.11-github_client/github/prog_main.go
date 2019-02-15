package github

import (
	"fmt"
)

// DisplayIssue displays a result set of issues in the terminal.
func DisplayIssue(c Config) error {

	// Run with defined configuration.
	result, err := makeRequest(c, nil)
	if err != nil {
		return fmt.Errorf("makeRequest: %v", err)
	}

	// If an mode other than mNone print to terminal.
	err = OutputResponce(c, result)
	if err != nil {
		return fmt.Errorf("OutputResponce: %v", err)
	}
	return nil
}

// RaiseIssue raises a new issue.
func RaiseIssue(c Config) error {

	// Compose a new issue.
	json, err := composeIssue(c)
	if err != nil {
		return fmt.Errorf("composeIssue: %v", err)
	}

	// Run with defined configuration.
	_, err = makeRequest(c, json)
	if err != nil {
		return fmt.Errorf("makeRequest: %v", err)
	}
	return nil
}

// EditIssue edits an existing issue.
func EditIssue(c Config) error {

	// Run with defined configuration.
	result, err := makeRequest(c, nil)
	if err != nil {
		return fmt.Errorf("makeRequest: %v", err)
	}

	// If edits are to be been made, edit and then post them to the server.
	json, err := editIssue(c, result.(Issue))
	if err != nil {
		return fmt.Errorf("editIssue: %v", err)
	}

	// Post the newly edited issue.
	_, err = makeRequest(c, json)
	if err != nil {
		return fmt.Errorf("makeRequest: json: %v", err)
	}
	return nil
}

// LockIssue locks a new issue.
func LockIssue(c Config) error {

	// Run with defined configuration.
	_, err := makeRequest(c, nil)
	if err != nil {
		return fmt.Errorf("makeRequest: %v", err)
	}
	return nil
}

// Program entry point as commandline client.
func Run(c Config) error {

	switch c.Mode {
	case mList:
		return DisplayIssue(c)
	case mRead:
		return DisplayIssue(c)
	case mRais:
		return RaiseIssue(c)
	case mEdit:
		return EditIssue(c)
	case mLock:
		return LockIssue(c)
	default:
		return fmt.Errorf("Run: c.Mode error hit end of switch statment")
	}
	return nil
}