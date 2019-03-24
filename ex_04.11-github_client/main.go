package main

import (
	"flag"
	"fmt"

	"tgpl/ex_04.11-github_client/gitish"
)

// Initiate custom flag usage call.
func setupFlags(flag *flag.FlagSet) {

	flag.Usage = func() {
		fmt.Println(gitish.Help)
	}
}

func main() {

	// Command line input.
	setupFlags(flag.CommandLine)
	flag.Parse()
	if gitish.Helpflag {
		println(gitish.Help)
		return
	}
	gitish.Conf.Queries = flag.Args()

	// Setup programming for selected mode, in some cases the program mode
	// is altered here, as such we pass in a pointer.
	err := gitish.SetState(gitish.Conf, gitish.FlagsIn)
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
