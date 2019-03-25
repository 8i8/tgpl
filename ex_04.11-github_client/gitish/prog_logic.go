package gitish

import (
	"bufio"
	"fmt"
	"os"
)

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  Set bitfield for program state
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// setBitAuth sets a flag in the bitfield if authentication is required for
// the http request that is needed.
func setBitAuth() error {

	if f&cVERBOSE > 0 {
		fmt.Printf("setBitAuth: testing for authentication requirements\n")
	}
	if f&(cEDIT|cRAISE|cLOCK|cUNLOCK) > 0 {
		f |= cAUTH
		return ckAuth()
	}
	return nil
}

// setBitMode set a bitfield flag from the input argument that defines the
// program running mode.
func setBitMode(in FlagsInStruct) error {

	// Ascertain program mode, assure the use of one only.
	if in.Read {
		f |= cREAD
		return ckRead()
	} else if in.List {
		f |= cLIST
		return ckList()
	} else if in.Edit {
		f |= cEDIT
		return ckRead()
	} else if in.Raise {
		f |= cRAISE
		return ckList()
	} else if in.Lock {
		f |= cLOCK
		return ckRead()
	} else if in.Unlock {
		f |= cUNLOCK
		return ckRead()
	} else if in.Set {
		f |= cSET
		return ckSet()
	}

	// None defined, set a default running mode.
	f |= cLIST
	return ckList()
}

// setBitValues sets a bitfield from input argument values.
func setBitValues(c Config) {

	// Set only one name.
	if len(c.Org) > 0 {
		f |= cORG
		f |= cNAME
	} else if len(c.Author) > 0 {
		f |= cAUTHOR
		f |= cNAME
	} else if len(c.User) > 0 {
		f |= cUSER
		f |= cNAME
	}

	if len(c.Repo) > 0 {
		f |= cREPO
	}
	if len(c.Number) > 0 {
		f |= cNUMBER
	}
	if len(c.Token) > 0 {
		f |= cTOKEN
	}
	if len(c.Editor) > 0 {
		f |= cEDITOR
	}
	if len(c.Reason) > 0 {
		f |= cREASON
	}
}

// SetBitState defines the state in which to run the program, set by the
// configuration of the users flags.
func SetBitState(c Config, fl FlagsInStruct) error {

	if fl.Verbose {
		f |= cVERBOSE
	}

	// If help has been requested, stop here and print out the help
	// information.
	if helpflag {
		println(help)
		return nil
	}

	// Set flags to denote existing values.
	setBitValues(c)

	// Set main running mode from input arguments, check logical coherence.
	err := setBitMode(fl)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	// If the current running mode requires authentication, set the flag
	// and test.
	err = setBitAuth()
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	if f&cVERBOSE > 0 {
		reportState("SetBitState")
	}

	return nil
}

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  Logical checks of program state.
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// ckAuth verify that the requirements for authentication are met.
func ckAuth() error {
	if f&cVERBOSE > 0 {
		fmt.Printf("ckAuth: testing authentication requirements\n")
	}
	if f&cTOKEN > 0 || f&cUSER > 0 {
		return nil
	}
	if f&cVERBOSE > 0 {
		reportState("ckAuth")
	}
	err := fmt.Errorf("please provide a user name or an OAuth2 token")
	return err
}

// ckRead verify that the requirements for Read mode are met.
func ckRead() error {
	if f&cVERBOSE > 0 {
		fmt.Printf("ckRead: testing for cNAME cREPO cNUMBER\n")
	}
	if f&cNAME > 0 && f&cREPO > 0 && f&cNUMBER > 0 {
		return nil
	}
	if f&cVERBOSE > 0 {
		reportState("ckRead")
	}
	err := fmt.Errorf("name repo and issue number are required")
	return err
}

// ckList verify that the requirements for the List mode are met.
func ckList() error {
	if f&cVERBOSE > 0 {
		fmt.Printf("ckList: testing for cNAME cREPO\n")
	}
	if f&cNAME > 0 || f&cREPO > 0 {
		return nil
	}
	if f&cVERBOSE > 0 {
		reportState("ckList")
	}
	err := fmt.Errorf("either a user or the repo name are required")
	return err
}

// ckSet Verifies that the requirements for a default configuration are met.
func ckSet() error {
	if f&cVERBOSE > 0 {
		fmt.Printf("ckSet: testing for cNAME cEDITOR\n")
	}
	if f&cNAME > 0 && f&cEDITOR > 0 {
		return nil
	}
	if f&cVERBOSE > 0 {
		reportState("ckSet")
	}
	err := fmt.Errorf("either a name or an editor command are required")
	return err
}

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  General use
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// reportState printout all set bitfield flags by name.
func reportState(context string) {

	w := bufio.NewWriter(os.Stdout)
	fmt.Fprintf(w, "%v", context)
	for i, s := range mState {
		if f&i > 0 {
			fmt.Fprintf(w, ": %v", s)
		}
	}
	fmt.Fprintf(w, "\n")
	w.Flush()
}
