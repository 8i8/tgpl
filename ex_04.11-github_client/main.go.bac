package main

import (
	"flag"
	"fmt"

	"tgpl/ex_04.11-github_client/gitish"
)

func main() {

	// Command line input.
	gitish.SetupFlags(flag.CommandLine)
	flag.Parse()
	gitish.Conf.Queries = flag.Args()

	// Setup programming for selected mode, in some cases the program mode
	// is altered here, as such we pass in a pointer.
	err := gitish.SetBitState(gitish.Conf, gitish.FlagsIn)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Run the program with given configuration.
	err = gitish.Run(gitish.Conf)
	if err != nil {
		fmt.Println(err)
		return
	}
}
