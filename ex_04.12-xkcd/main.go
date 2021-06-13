package main

import (
	"flag"
	"xkcd/xkcd"
)

func init() {
	flag.BoolVar(&xkcd.VERBOSE, "v", false, "") // Verbose mode.
	flag.BoolVar(&xkcd.UPDATE, "u", false, "")  // Update database.
	flag.BoolVar(&xkcd.TITLE, "l", false, "")   // Print a list of title number and address.
	flag.IntVar(&xkcd.TESTRUN, "test", 0, "")   // Test database.
	flag.IntVar(&xkcd.DBGET, "n", -1, "")       // Display comic 'n'.
	flag.IntVar(&xkcd.WEBGET, "w", 0, "")       // Display comic 'n' by way of http.
}

func main() {
	xkcd.SetupFlags(flag.CommandLine)
	flag.Parse()
	xkcd.Run(flag.Args())
}
