package gitish

import (
	"flag"
)

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
   Configuration
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// The Programs main running state.
type flags uint

// Bitfield state.
var f flags

// Conf is the main program configuration data.
var Conf Config

// FlagsIn contains the user set flags from the command line.
var FlagsIn FlagsInStruct

// Config contains program operation specific data.
type Config struct {
	User   string // Repository owner,
	Token  string // Oauth token.
	Editor string // External editor.
	Reason string // Reason for locking.
	Req           // Request data.
}

// Req contains all of the data of a particular request.
type Req struct {
	Author  string   // Author name.
	Org     string   // Organisation name.
	Repo    string   // Repository name.
	Number  string   // Issue number.
	Queries []string // Queries that have been retrieved from the Args[] array.
}

// FlagsInStruct flags set from the user command line arguments.
type FlagsInStruct struct {
	Read    bool
	List    bool
	Edit    bool
	Raise   bool
	Lock    bool
	Unlock  bool
	Set     bool
	Verbose bool
}

// Initialise flags, used in main at program start.
func init() {

	// Mode
	flag.BoolVar(&FlagsIn.Read, "read", false, "")
	flag.BoolVar(&FlagsIn.List, "list", false, "")
	flag.BoolVar(&FlagsIn.Edit, "edit", false, "")
	flag.BoolVar(&FlagsIn.Raise, "raise", false, "")
	flag.BoolVar(&FlagsIn.Lock, "lock", false, "")
	flag.BoolVar(&FlagsIn.Unlock, "unlock", false, "")
	flag.BoolVar(&FlagsIn.Set, "set", false, "")

	// Flags
	flag.BoolVar(&FlagsIn.Verbose, "v", false, "")

	// Help
	flag.BoolVar(&helpflag, "h", false, "")
	flag.BoolVar(&helpflag, "help", false, "")

	// Values
	flag.StringVar(&Conf.User, "u", "", "")
	flag.StringVar(&Conf.Author, "a", "", "")
	flag.StringVar(&Conf.Org, "o", "", "")
	flag.StringVar(&Conf.Repo, "r", "", "")
	flag.StringVar(&Conf.Number, "n", "", "")
	flag.StringVar(&Conf.Token, "t", "", "")
	flag.StringVar(&Conf.Editor, "d", "", "")
	flag.StringVar(&Conf.Reason, "l", "", "")
}

// All flags.
const (
	// Program mode.
	cREAD   flags = 1 << iota // Read mode expects a single issue response.
	cLIST                     // List mode expects multiple issue response.
	cEDIT                     // Edit a single issue.
	cRAISE                    // Raise a new issue.
	cLOCK                     // Lock an issue.
	cUNLOCK                   // Unlock a locked issue.
	cSET                      // Set default global editor and user name.

	// Values
	cUSER   // Every values flag set indicates that the named value has been
	cAUTHOR // provided.
	cORG
	cREPO
	cNUMBER
	cTOKEN
	cEDITOR
	cREASON

	// Flags
	cVERBOSE // Print out extra detail to the command line.

	// Internally defined.
	cNAME // Used as an indicator to signal that either a user name an
	// author or and organisation have been provided.
	cAUTH // Used as an indicator to signal that authentication is required.

)

// cADDRESS stores the combined flag set that defined the state requirement for
// a direct address HTTP access rather than a search, used by cREAD and other
// similarly single response or address function.
var cADDRESS flags

func init() {
	cADDRESS |= (cNAME | cREPO | cNUMBER)
}

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  Map for outputting bitfield flags by name.
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// mState is a hashmap for printing out flags by name, used by the verbose
// flag.
var mState = make(map[flags]string)

// Load map with flag names.
func init() {

	// mode
	mState[cREAD] = "cREAD"
	mState[cLIST] = "cLIST"
	mState[cEDIT] = "cEDIT"
	mState[cRAISE] = "cRAISE"
	mState[cLOCK] = "cLOCK"
	mState[cUNLOCK] = "cUNLOCK"
	mState[cSET] = "cSET"

	// Values
	mState[cUSER] = "cUSER"
	mState[cAUTHOR] = "cAUTHOR"
	mState[cORG] = "cORG"
	mState[cREPO] = "cREPO"
	mState[cNUMBER] = "cNUMBER"
	mState[cTOKEN] = "cTOKEN"
	mState[cEDITOR] = "cEDITOR"
	mState[cREASON] = "cREASON"

	// Flags
	mState[cVERBOSE] = "cVERBOSE"

	// Internally defined.
	mState[cNAME] = "cNAME"
	mState[cAUTH] = "cAUTH"
}
