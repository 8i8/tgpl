package gitish

import (
	"bytes"
	"fmt"
)

// RaiseIssue raise a new issue on a github repository.
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
