package main

import (
	"flag"

	"tgpl/ex_04.12-xkcd/xkcd"
)

func init() {
	flag.BoolVar(&xkcd.VERBOSE, "v", false, "")
	flag.BoolVar(&xkcd.UPDATE, "u", false, "")
	flag.UintVar(&xkcd.TESTRUN, "t", 0, "")
	flag.UintVar(&xkcd.RECORD, "n", 0, "")
}

func main() {

	// Command line input.
	xkcd.SetupFlags(flag.CommandLine)
	flag.Parse()

	xkcd.Init()
}
