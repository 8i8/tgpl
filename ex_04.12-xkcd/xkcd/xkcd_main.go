package xkcd

import (
	"fmt"

	"tgpl/ex_04.12-xkcd/quest"
)

const cLastURL = "https://xkcd.com/info.0.json"
const cBaseURL = "https://xkcd.com/"
const cTailURL = "/info.0.json"

// Database file name.
var cNAME = "xkcd.json"

// Verbouse program output whilst running.
var (
	VERBOSE bool
	UPDATE  bool
	SEARCH  bool
	LIST    bool
	DBGET   uint
	WEBGET  uint
	TESTRUN uint
)

// setConfig sets required state variables for desired program run mode.
func setConfig() {

	// sets quest package to verbose.
	if VERBOSE {
		quest.VERBOSE = true
	}

	// Sets program to generate a test database.
	if TESTRUN > 0 {
		cNAME = "test.json"
	}
}

// Run is the xkcd main program routine.
func Run(args []string) {

	// Program state.
	setConfig()

	if WEBGET > 0 {
		WebGet(WEBGET)
		return
	}

	comics, err := xkcdInit()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	if DBGET > 0 {
		comics.DbGet(DBGET - 1)
		return
	}

	if SEARCH {
		comics.Search(args)
		return
	}

	if UPDATE {
		comics.Update()
	}
}
