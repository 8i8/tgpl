package gitish

import (
	"fmt"
)

// Header contains header request key paires.
type Header struct {
	Key, Value string
}

// Status contains an http responce status, both the text and the code.
type Status struct {
	Code   int
	Reason string
}

// MakeRequest orchistrates an http request.
func MakeRequest(conf Config) error {

	var err error
	switch conf.Mode {
	case mList:
		err = ListIssues(conf)
	case mRead:
		err = ReadIssue(conf)
	case mEdit:
		err = EditIssue(conf)
	case mLock:
		err = ReadIssue(conf)
	case mRaise:
		err = RaiseIssue(conf)
	default:
		fmt.Println("Run with -h flag for user instructions.")
	}

	_, _, err = setURL(conf)

	return err
}

// func WriteResponce() {
// }
