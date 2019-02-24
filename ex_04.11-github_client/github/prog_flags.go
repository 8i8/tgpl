package github

// The Programs main running state.
var f Flags

const (
	cLIST Flags = 1 << iota
	cREAD
	cRAISE
	cEDIT
	cLOCK
	cRAW
	cMANY
	cLONE
	cNONE
	cVERBOSE
)

var mStateName = make(map[Flags]string)

func init() {

	// mode
	mStateName[cLIST] = "cLIST"
	mStateName[cREAD] = "cREAD"
	mStateName[cEDIT] = "cEDIT"
	mStateName[cRAISE] = "cRAISE"
	mStateName[cLOCK] = "cLOCK"
	mStateName[cRAW] = "cRAW"

	// response
	mStateName[cNONE] = "cNONE"
	mStateName[cLONE] = "cLONE"
	mStateName[cMANY] = "cMANY"
	mStateName[cRAW] = "cRAW"
	mStateName[cVERBOSE] = "cVERBOSE"
}

type FlagsIn struct {
	Edit    bool // Signal request to edit an issue.
	Lock    bool // Lock a repository.
	Raise   bool // Raise a new issue.
	Raw     bool // Raw request input.
	Verbose bool // Signals the program print out extra detail.
}

func getFlags(in FlagsIn) {

	// Set state for program global.
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
