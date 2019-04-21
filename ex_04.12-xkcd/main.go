package main

import (
	"flag"

	"tgpl/ex_04.12-xkcd/xkcd"
)

func init() {
	flag.BoolVar(&xkcd.VERBOSE, "v", false, "")   // Verbose mode.
	flag.BoolVar(&xkcd.UPDATE, "u", false, "")    // Update database.
	flag.BoolVar(&xkcd.SEARCH, "s", false, "")    // Search for <args>
	flag.BoolVar(&xkcd.BTREE, "b", false, "")     // Use the list instead of btree.
	flag.BoolVar(&xkcd.BTREESYNC, "y", false, "") // Use the sync map and goroutines.
	flag.BoolVar(&xkcd.LISTSYNC, "ls", false, "") // Use the sync map and goroutines.
	flag.UintVar(&xkcd.TESTRUN, "t", 0, "")       // Test database.
	flag.UintVar(&xkcd.DBGET, "n", 0, "")         // Display comic 'n'.
	flag.UintVar(&xkcd.WEBGET, "w", 0, "")        // Display comic 'n' http.
}

func xkcdProgram() {
	// Command line input.
	xkcd.SetupFlags(flag.CommandLine)
	flag.Parse()

	xkcd.Run(flag.Args())
}

func main() {

	xkcdProgram()
	//ds.Test()
}
