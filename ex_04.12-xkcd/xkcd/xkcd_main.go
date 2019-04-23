package xkcd

import (
	"fmt"
	"math/rand"
	"time"

	"tgpl/ex_04.12-xkcd/ds"
	"tgpl/ex_04.12-xkcd/quest"
)

const cLastURL = "https://xkcd.com/info.0.json"
const cBaseURL = "https://xkcd.com/"
const cTailURL = "/info.0.json"

// Database file name.
var cNAME = "xkcd.json"
var cADDRESS = "data/"

// Verbouse program output whilst running.
var (
	VERBOSE bool
	UPDATE  bool
	SEARCH  bool
	TITLE   bool
	DBGET   int
	WEBGET  int
	TESTRUN int
)

// setConfig sets required state variables for desired program run mode.
func setConfig() {

	// sets quest package to verbose.
	if VERBOSE {
		quest.VERBOSE = true
		ds.Verbose()
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
		fmt.Printf("error: xkcdInit: %v\n", err)
		return
	}

	if DBGET > 0 {
		comics.DbGet(DBGET)
		return
	}

	// Prints the latest comic
	if DBGET == 0 {
		comics.DbGet(0)
		return
	}

	// Print out the latest comic if the data base has been updated.
	if UPDATE {
		if comics.Update() {
			comics.DbGet(0)
		}
		return
	}

	if len(args) > 0 {
		comics.Search(args)
		return
	}

	// When the program is run with no arguments print out a random comic
	// from the database.
	rand.Seed(time.Now().UnixNano())
	comics.DbGet(uint(rand.Intn(int(comics.Len))))

}
