package github

import (
	"bufio"
	"fmt"
	"os"
)

// The Programs main running state.
var f Flags

const (
	// Program mode.
	cREAD    Flags = 1 << iota // Request read mode.
	cLIST                      // Request list mode.
	cEDIT                      // Request to edit an issue.
	cRAISE                     // Raise a new issue.
	cLOCK                      // Lock a repository.
	cUNLOCK                    // Unlock an existing locked issue.
	cSET                       // Set default global editor and user name.
	cVERBOSE                   // Signals the program print out extra detail.

	// Data.
	cUSER   // User name given.
	cAUTHOR // Authour name given.
	cORG    // Organisation name given.
	cREPO   // Repoitory name given.
	cNUMBER // Issue number given.
	cTOKEN  // Oauth2 token given.
	cEDITOR // Editor defined.
	cNAME   // Used as an indicator to signal that either a user name an
	// author or and organisation have been provided.
)

var mState = make(map[Flags]string)

func init() {

	// mode
	mState[cREAD] = "cREAD"
	mState[cLIST] = "cLIST"
	mState[cEDIT] = "cEDIT"
	mState[cRAISE] = "cRAISE"
	mState[cLOCK] = "cLOCK"
	mState[cUNLOCK] = "cUNLOCK"
	mState[cSET] = "cSET"
	mState[cVERBOSE] = "cVERBOSE"

	// response
	mState[cUSER] = "cUSER"
	mState[cAUTHOR] = "cAUTHOR"
	mState[cORG] = "cORG"
	mState[cREPO] = "cREPO"
	mState[cNUMBER] = "cNUMBER"
	mState[cTOKEN] = "cTOKEN"
	mState[cEDITOR] = "cEDITOR"
}

// FlagsIn is the strut to pass user command line settings into the program.
type FlagsIn struct {
	Read    bool
	List    bool
	Edit    bool
	Raise   bool
	Unlock  bool
	Set     bool
	Verbose bool
}

// assesInput sets flags for all input values provided, used in preferance over
// the len() function to switch the programs control flow.
func assesInput(c Config) {

	// Set only one name as the address name in the case that more than one
	// have been provided.
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

	if len(c.Lock) > 0 {
		f |= cLOCK
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
}

// ckRead verify that the requirments for Read mode have been met.
func ckRead() error {
	if f&cVERBOSE > 0 {
		fmt.Printf("ckRead: testing for cNAME|cREPO|cNUMBER\n")
		reportState("ckRead")
	}
	if f&cNAME > 0 && f&cREPO > 0 && f&cNUMBER > 0 {
		return nil
	}
	err := fmt.Errorf("name repo and issue number are required")
	return err
}

// ckList verifys that the requirments for the List mode have been met.
func ckList() error {
	if f&cVERBOSE > 0 {
		fmt.Printf("ckList: testing for cNAME cREPO\n")
		reportState("ckList")
	}
	if f&cNAME > 0 && f&cREPO > 0 {
		return nil
	}
	err := fmt.Errorf("at the very least a user name or the repo are required")
	return err
}

// ckAll verifys that the requirments have been met for a direct address
// autorised acces to an issue, needed by the raise edit lock and unlock
// functions.
func ckAll() error {
	if f&cVERBOSE > 0 {
		fmt.Printf("ckAll: testing for cNAME cREPO cNUMBER cTOKEN\n")
		reportState("ckAll")
	}
	if f&cNAME > 0 && f&cREPO > 0 && f&cNUMBER > 0 && f&cTOKEN > 0 {
		return nil
	}
	err := fmt.Errorf("name, repo, number and an Oauth2 token all required")
	return err
}

// ckSet Verifies that the requrments have been met to set a default
// configuration.
func ckSet() error {
	if f&cVERBOSE > 0 {
		fmt.Printf("ckSet: testing for cNAME cEDITOR\n")
		reportState("ckSet")
	}
	if f&cNAME > 0 && f&cEDITOR > 0 {
		return nil
	}
	err := fmt.Errorf("either a name or an editor command are required")
	return err
}

// setMode translates the user input flags into the programs main state
// variable.
func setMode(in FlagsIn) error {

	// Assertain program mode, assure the use of one only.
	if in.Read {
		f |= cREAD
		return ckRead()
	} else if in.List {
		f |= cLIST
		return ckList()
	} else if in.Edit {
		f |= cEDIT
		return ckAll()
	} else if in.Raise {
		f |= cRAISE
		return ckList()
	} else if f&cLOCK > 0 {
		return ckAll()
	} else if in.Unlock {
		f |= cUNLOCK
		return ckAll()
	} else if in.Set {
		f |= cSET
		return ckSet()
	}

	// None set, providea a default running mode.
	f |= cLIST
	return ckList()
}

// SetState defines the state in which to run the program, set by the
// configuration of the users flags.
func SetState(c Config, fl FlagsIn) error {

	// Indipendant flags.
	if fl.Verbose {
		f |= cVERBOSE
	}

	// Set booleans to mirror any flag values that may have been input.
	assesInput(c)

	// Set state variable from user input, flags and settings.
	err := setMode(fl)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	if f&cVERBOSE > 0 {
		reportState("SetState")
	}

	return nil
}

// reportState outputs the name of all booleans that are set.
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
