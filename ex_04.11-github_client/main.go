package main

import (
	"flag"
	"fmt"

	"tgpl/ex_04.11-github_client/github"
)

func main() {

	// Command line input.
	github.SetupFlags(flag.CommandLine)
	flag.Parse()
	github.Conf.Queries = flag.Args()

	// Setup programming for selected mode, in some cases the program mode
	// is altered here, as such we pass in a pointer.
	err := github.SetBitState(github.Conf, github.FlagsIn)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Run the program with given configuration.
	err = github.Run(github.Conf)
	if err != nil {
		fmt.Println(err)
		return
	}
}
