package github

// The Programs main running state.
var f Flags

const (
	// Boolean flags.
	cEDIT Flags = 1 << iota
	cLOCK
	cRAISE
	cRAW
	cVERBOSE
	// Text settings.
	cBASE  // First indices of address.
	cUSER
	cAUTHOR
	cORG
	// Data details.
	cREPO  // Second indices of address.
	cNUMBER // Third indices of address.
	cTOKEN
	cEDITOR
	// Consequent program mode.
	cAUTH
	cRESP
	cCREATE
)

var (
	READ_RECORD = Flags
	READ_LIST
)

var mStateName = make(map[Flags]string)

func init() {

	// mode
	mStateName[cEDIT] = "cEDIT"
	mStateName[cLOCK] = "cLOCK"
	mStateName[cRAISE] = "cRAISE"
	mStateName[cRAW] = "cRAW"
	mStateName[cVERBOSE] = "cVERBOSE"

	// response
	mStateName[cBASE] = "cBASE"
	mStateName[cUSER] = "cUSER"
	mStateName[cAUTHOR] = "cAUTHOR"
	mStateName[cORG] = "cORG"

	mStateName[cREPO] = "cREPO"
	mStateName[cNUMBER] = "cNUMBER"
	mStateName[cTOKEN] = "cTOKEN"
	mStateName[cEDITOR] = "cEDITOR"

	mStateName[cAUTH] = "cAUTH"
	mStateName[cRESP] = "cRESP"
	mStateName[cCREATE] = "cCREATE"
}

// FlagsIn is the strut to pass user command line settings into the program.
type FlagsIn struct {
	Edit    bool // Signal request to edit an issue.
	Lock    bool // Lock a repository.
	Raise   bool // Raise a new issue.
	Raw     bool // Raw request input.
	Verbose bool // Signals the program print out extra detail.
}

// getConfig sets user input configuration details in the programmes main state
// variable.
func getConfig(c Config) {
	if len(c.User) > 0 {
		f |= (cUSER | cBASE)
	}
	if len(c.Author) > 0 {
		f |= (cAUTHOR | cBASE)
	}
	if len(c.Org) > 0 {
		f |= (cORG | cBASE)
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
	if len(c.EDITOR) > 0 {
		f |= cEDITOR
	}
}

// getFlags translates the user input flags into the programs main state
// variable.
func getFlags(in FlagsIn) {

	if in.Edit {
		f |= cEDIT
	}
	if in.Lock {
		f |= cLOCK
	}
	if in.Raise {
		f |= cRAISE
	}
	if in.Raw {
		f |= cRAW
	}
	if in.Verbose {
		f |= cVERBOSE
	}
}
