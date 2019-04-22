package main

import (
	"flag"

	"tgpl/ex_04.12-xkcd/xkcd"
)

func init() {
	flag.BoolVar(&xkcd.VERBOSE, "v", false, "") // Verbose mode.
	flag.BoolVar(&xkcd.UPDATE, "u", false, "")  // Update database.
	flag.BoolVar(&xkcd.SEARCH, "s", false, "")  // Search for <args>
	flag.BoolVar(&xkcd.TITLE, "l", false, "")   // Print a list of title number and address.
	flag.UintVar(&xkcd.TESTRUN, "test", 0, "")  // Test database.
	flag.UintVar(&xkcd.DBGET, "n", 0, "")       // Display comic 'n'.
	flag.UintVar(&xkcd.WEBGET, "w", 0, "")      // Display comic 'n' http.
}

func main() {
	xkcd.SetupFlags(flag.CommandLine)
	flag.Parse()

	xkcd.Run(flag.Args())
}
