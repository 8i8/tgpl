package github

import (
	"fmt"
)

// The run state of the program, interpreted by command line flags. This
// variable is set as in integer within the configuration struct, by the
// function SetState(c Config) at program start.
type Mode int

const (
	mList Mode = iota
	mRead
	mRais
	mEdit
	mLock
	mRaw
)

type Resp int

// Mode of the expected http response type.
const (
	rMany Resp = iota
	rLone
	rNone
	rRaw
)

// The Programs main running state.
var rState Resp

var mState [6]string
var rStateName [4]string

func init() {
	// mode
	mState[mList] = "mList"
	mState[mRead] = "mRead"
	mState[mEdit] = "mEdit"
	mState[mRais] = "mRais"
	mState[mLock] = "mLock"
	mState[mRaw] = "mRaw"
	// response
	rStateName[rNone] = "rNone"
	rStateName[rLone] = "rLone"
	rStateName[rMany] = "rMany"
	rStateName[rRaw] = "rRaw"
}

// isFullAddress checks if the requirements have been met to enter rLone
// mode.
func isFullAddress(c Config) bool {
	return (len(c.Author) > 0 || len(c.User) > 0 || len(c.Org) > 0) &&
		len(c.Repo) > 0
}

// checkModeValid verifies that there are not two contradicting flags set.
func checkModeValid(c Config) error {
	tally := 0
	if c.Edit {
		tally++
	}
	if c.Lock {
		tally++
	}
	if c.Raise {
		tally++
	}
	if tally > 1 {
		str := "Please define either -x -e or -l"
		return fmt.Errorf(str)
	}
	return nil
}

// setRunMode sets the program run mode from the given flags.
func setRunMode(c *Config) {

	// Set edit or lock state before checking for the presence of a number.
	if c.Edit {
		c.Mode = mEdit
	} else if c.Lock {
		c.Mode = mEdit
	} else if c.Raise {
		c.Mode = mRais
	} else if c.Raw {
		c.Mode = mRaw
	}

	// In the case where a number is explicitly provided and the required
	// parameters exist for a direct HTTP access then do so, else add the
	// number to the query, listing as a search parameter and then expect
	// multiple results.
	if len(c.Number) > 0 && !c.Edit && !c.Lock {
		if isFullAddress(*c) {
			c.Mode = mRead
		} else {
			c.Queries = append(c.Queries, c.Number)
			c.Mode = mList
		}
	}
}

// setRespExp sets the program HTTP response expectation from the previously
// defined running mode.
func setRespExp(c *Config) {

	if c.Mode == mList {
		rState = rMany
	} else if c.Mode == mRead {
		rState = rLone
	} else if c.Mode == mRais {
		rState = rNone
	} else if c.Mode == mEdit {
		rState = rLone
	} else if c.Mode == mLock {
		rState = rNone
	} else if c.Mode == mRaw {
		rState = rRaw
	}
}

// SetState defines the state in which to run the program, set by the
// configuration of the users flags.
func SetState(c *Config) error {

	// Error check flags for state contradiction.
	err := checkModeValid(*c)
	if err != nil {
		return fmt.Errorf("SetState: %v", err)
	}

	// 1) set one of five program modes.
	setRunMode(c)
	// 2) set one of three expected response modes.
	setRespExp(c)

	// Output state in verbose mode.
	if c.Verbose {
		fmt.Printf("SetState run mode: %v\n", mState[c.Mode])
		fmt.Printf("SetState response mode: %v\n", rStateName[rState])
	}

	return nil
}
