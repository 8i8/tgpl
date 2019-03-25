package gitish

import (
	"fmt"
)

// DisplayIssue displays a list of issues in the terminal.
func DisplayIssue(c Config) error {

	// Run with current configuration.
	reply, err := makeRequest(c, nil)
	if err != nil {
		return fmt.Errorf("makeRequest: %v", err)
	}

	// Print to terminal.
	err = outputResponse(c, reply)
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

	// Run with current configuration.
	_, err = makeRequest(c, json)
	if err != nil {
		return fmt.Errorf("makeRequest: %v", err)
	}
	return nil
}

// EditIssue edits an existing issue.
func EditIssue(c Config) error {

	// Set state for read to aquire the targeted issue.
	f &^= cEDIT
	f &^= cAUTH
	f |= cREAD
	reportState("EditIssue READ")

	// Run with current configuration.
	reply, err := makeRequest(c, nil)
	if err != nil {
		return fmt.Errorf("makeRequest: %v", err)
	}

	// If edits are to be been made, make them.
	json, err := editIssue(c, reply)
	if err != nil {
		return fmt.Errorf("editIssue: %v", err)
	}

	// Set to authentication for updating edited issue.
	f |= cEDIT
	f |= cAUTH
	f &^= cREAD
	reportState("EditIssue EDIT")

	// Post issue.
	_, err = makeRequest(c, json)
	if err != nil {
		return fmt.Errorf("makeRequest: json: %v", err)
	}
	return nil
}

// LockUnlockIssue locks a new issue.
func LockUnlockIssue(c Config) error {

	var json []byte
	var err error

	// Marshal into json format.
	if f&cREASON > 0 {
		// Write specifics json data to lock the issue.
		// TODO this does not appear to be working.
		json, err = lockReasonJSON(c.Reason)
		if err != nil {
			return fmt.Errorf("lockIssue: %v", err)
		}
	}

	// Run with current configuration.
	_, err = makeRequest(c, json)
	if err != nil {
		return fmt.Errorf("makeRequest: %v", err)
	}
	return nil
}

// Run is the main programs entry point.
func Run(c Config) error {

	switch {
	case f&cREAD > 0:
		return DisplayIssue(c)
	case f&cLIST > 0:
		return DisplayIssue(c)
	case f&cEDIT > 0:
		return EditIssue(c)
	case f&cRAISE > 0:
		return RaiseIssue(c)
	case f&cLOCK > 0:
		return LockUnlockIssue(c)
	case f&cUNLOCK > 0:
		return LockUnlockIssue(c)
	default:
		str := "Run: c.Mode error hit end of switch statment"
		return fmt.Errorf(str)
	}
}
