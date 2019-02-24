package github

import (
	"fmt"
)

// DisplayIssue displays a result set of issues in the terminal.
func DisplayIssue(c Config) error {

	// Run with defined configuration.
	reply, err := makeRequest(c, nil)
	if err != nil {
		return fmt.Errorf("makeRequest: %v", err)
	}

	// Print to terminal.
	err = OutputResponse(c, reply)
	if err != nil {
		return fmt.Errorf("OutputResponse: %v", err)
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
// TODO NOW state is being altered here, should it be?
func EditIssue(c Config) error {

	// Set default state to use GET
	f |= cLONE
	f |= cREAD

	// Run with defined configuration.
	reply, err := makeRequest(c, nil)
	if err != nil {
		return fmt.Errorf("makeRequest: %v", err)
	}

	// If edits are to be been made, edit and then post them to the server.
	json, err := editIssue(c, reply)
	if err != nil {
		return fmt.Errorf("editIssue: %v", err)
	}

	// Set state to use authentication.
	f |= cNONE
	f |= cEDIT

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

	switch {
	case f&cLIST > 0:
		return DisplayIssue(c)
	case f&cREAD > 0:
		return DisplayIssue(c)
	case f&cRAISE > 0:
		return RaiseIssue(c)
	case f&cEDIT > 0:
		return EditIssue(c)
	case f&cLOCK > 0:
		return LockIssue(c)
	case f&cRAW > 0:
		return DisplayIssue(c)
	default:
		str := "Run: c.Mode error hit end of switch statment"
		return fmt.Errorf(str)
	}
	return nil
}
