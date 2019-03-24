package main

import (
	"flag"

	"tgpl/ex_04.11-github_client/github"
)

func main() {

	// Command line input.
	flag.Parse()
	if github.Helpflag {
		println(github.Help)
		return
	}
	github.Conf.Queries = flag.Args()

	// Setup programming for selected mode, in some cases the program mode
	// is altered here, as such we pass in a pointer.
	err := github.SetState(github.Conf, github.FlagsIn)
	if err != nil {
		println(err)
		return
	}

	// Run the program with given configuration.
	err = github.Run(github.Conf)
	if err != nil {
		println(err)
		return
	}
}
