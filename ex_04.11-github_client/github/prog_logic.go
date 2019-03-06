package github

import (
	"fmt"
)

// The run state of the program, interpreted by command line flags. This
// variable is set as in integer within the configuration struct, by the
// function SetState(c Config) at program start.
type Mode int

// hasFullAddress checks if the requirements are met to use a direct address.
func hasFullAddress() bool {
	return (f&cAUTHOR > 0 || f&cUSER > 0 || f&c.ORG > 0) && f&cREPO > 0
}

// checkModeValid verifies that there are not two contradicting flags set.
func checkModeValid() error {
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
		str := "Please define only one of either -x -e or -l"
		return fmt.Errorf(str)
	}
	return nil
}

// SetState defines the state in which to run the program, set by the
// configuration of the users flags.
func SetState(c *Config, fl FlagsIn) error {

	// Set state variable from user input, flags and settings.
	getFlags(fl)
	getConfig(c)

	// Error check flags that no conradicting states exist.
	err := checkModeValid()
	if err != nil {
		return fmt.Errorf("SetState: %v", err)
	}

	// Set default state from input flags covers all base cases.
	f |= READ_LIST

	// In the case where a number is explicitly provided and the required
	// parameters exist for a direct HTTP access then do so, else add the
	// number to the query, listing as a search parameter and in
	// consiquence expect multiple results.
	if f&cNUMBER > 0 && f&cEDIT == 0 && f&cLOCK == 0 {
		if hasFullAddress() {
			f |= READ_RECORD
		} else {
			c.Queries = append(c.Queries, c.Number)
			f |= READ_LIST
		}
	}

	// Set programs http responce expectation.

	// Output state in verbose mode.
	// if f&cVERBOSE > 0 {
	// 	fmt.Printf("SetState run mode: %v\n", mStateName[b^a])
	// 	fmt.Printf("SetState response mode: %v\n", mStateName[f^b])
	// }

	return nil
}
