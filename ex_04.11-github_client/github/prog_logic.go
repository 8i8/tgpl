package github

import (
	"fmt"
)

// The run state of the program, interpreted by command line flags. This
// variable is set as in integer within the configuration struct, by the
// function SetState(c Config) at program start.
type Mode int

// isFullAddress checks if the requirements have been met to enter cLONE
// mode.
func isFullAddress(c Config) bool {
	return (len(c.Author) > 0 || len(c.User) > 0 || len(c.Org) > 0) &&
		len(c.Repo) > 0
}

// checkModeValid verifies that there are not two contradicting flags set.
func checkModeValid(c Config) error {
	count := 0
	if f&cEDIT > 0 {
		count++
	}
	if f&cLOCK > 0 {
		count++
	}
	if f&cRAISE > 0 {
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
	f |= (cMANY | cLIST)
}

// setRunMode sets the program run mode from the given flags.
func setRunMode(c *Config) {

	setDefaults(c)

	// flag.StringVar(&conf.User, "u", "", user)
	// flag.StringVar(&conf.Author, "a", "", author)
	// flag.StringVar(&conf.Org, "o", "", org)
	// flag.StringVar(&conf.Repo, "r", "", repo)
	// flag.StringVar(&conf.Number, "n", "", number)
	// flag.StringVar(&conf.Token, "t", "", token)
	// flag.StringVar(&conf.Editor, "d", "", editor)

	// In the case where a number is explicitly provided and the required
	// parameters exist for a direct HTTP access then do so, else add the
	// number to the query, listing as a search parameter and in
	// consiquence expect multiple results.
	if len(c.Number) > 0 && f&cEDIT > 0 && f&cLOCK > 0 {
		if isFullAddress(*c) {
			f |= cREAD
		} else {
			c.Queries = append(c.Queries, c.Number)
			f |= cLIST
		}
	}
}

// setRespExp sets the program HTTP response expectation from the previously
// defined running mode.
func setRespExp(c *Config) {

	if f&cLIST > 0 {
		f |= cMANY
	} else if f&cREAD > 0 {
		f |= cLONE
	} else if f&cRAISE > 0 {
		f |= cNONE
	} else if f&cEDIT > 0 {
		f |= cLONE
	} else if f&cLOCK > 0 {
		f |= cNONE
	} else if f&cRAW > 0 {
		f |= cRAW
	}
}

// SetState defines the state in which to run the program, set by the
// configuration of the users flags.
func SetState(c *Config, fl FlagsIn) error {

	// Set state from input flags.
	getFlags(fl)

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
	// if f&cVERBOSE > 0 {
	// 	fmt.Printf("SetState run mode: %v\n", mStateName[b^a])
	// 	fmt.Printf("SetState response mode: %v\n", mStateName[f^b])
	// }

	return nil
}
