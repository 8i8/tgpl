package github

import (
	"fmt"
)

// The run state of the program, interpreted by command line flags. This
// variable is set as in integer within the configuration struct, by the
// function SetState(c Config) at program start.
type Mode int

const (
	mLIST Mode = 1 << iota
	mREAD
	mRAISE
	mEDIT
	mLOCK
	mRAW
)

type Resp int

// Mode of the expected http response type.
const (
	rMANY Resp = 1 << iota
	rLONE
	rNONE
	rRAW
)

// The Programs main running state.
var rState Resp
var mStateName = make(map[Mode]string)
var rStateName = make(map[Resp]string)

func init() {

	// mode
	mStateName[mLIST] = "mLIST"
	mStateName[mREAD] = "mREAD"
	mStateName[mEDIT] = "mEDIT"
	mStateName[mRAISE] = "mRAISE"
	mStateName[mLOCK] = "mLOCK"
	mStateName[mRAW] = "mRAW"

	// response
	rStateName[rNONE] = "rNONE"
	rStateName[rLONE] = "rLONE"
	rStateName[rMANY] = "rMANY"
	rStateName[rRAW] = "rRAW"
}

// isFullAddress checks if the requirements have been met to enter rLONE
// mode.
func isFullAddress(c Config) bool {
	return (len(c.Author) > 0 || len(c.User) > 0 || len(c.Org) > 0) &&
		len(c.Repo) > 0
}

// checkModeValid verifies that there are not two contradicting flags set.
func checkModeValid(c Config) error {
	count := 0
	if c.Edit {
		count++
	}
	if c.Lock {
		count++
	}
	if c.Raise {
		count++
	}
	if count > 1 {
		str := "Please define either -x -e or -l"
		return fmt.Errorf(str)
	}
	return nil
}

//setDefault sets the default state of the program.
func setDefaults(c *Config) {
	rState = rMANY
	c.Mode = mLIST
}

// setRunMode sets the program run mode from the given flags.
func setRunMode(c *Config) {

	setDefaults(c)

	// Set edit or lock state before checking for the presence of a number.
	if c.Edit {
		c.Mode = mEDIT
	} else if c.Lock {
		c.Mode = mEDIT
	} else if c.Raise {
		c.Mode = mRAISE
	} else if c.Raw {
		c.Mode = mRAW
	}

	// In the case where a number is explicitly provided and the required
	// parameters exist for a direct HTTP access then do so, else add the
	// number to the query, listing as a search parameter and in
	// consiquence expect multiple results.
	if len(c.Number) > 0 && !c.Edit && !c.Lock {
		if isFullAddress(*c) {
			c.Mode = mREAD
		} else {
			c.Queries = append(c.Queries, c.Number)
			c.Mode = mLIST
		}
	}
}

// setRespExp sets the program HTTP response expectation from the previously
// defined running mode.
func setRespExp(c *Config) {

	if c.Mode == mLIST {
		rState = rMANY
	} else if c.Mode == mREAD {
		rState = rLONE
	} else if c.Mode == mRAISE {
		rState = rNONE
	} else if c.Mode == mEDIT {
		rState = rLONE
	} else if c.Mode == mLOCK {
		rState = rNONE
	} else if c.Mode == mRAW {
		rState = rRAW
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

	// 1) set program modes.
	setRunMode(c)
	// 2) set expected response modes.
	setRespExp(c)

	// Output state in verbose mode.
	if c.Verbose {
		fmt.Printf("SetState run mode: %v\n", mStateName[c.Mode])
		fmt.Printf("SetState response mode: %v\n", rStateName[rState])
	}

	return nil
}
