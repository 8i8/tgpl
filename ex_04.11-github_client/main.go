package main

import (
	"flag"
	"fmt"

	"tgpl/ex_04.11-github_client/github"
)

var conf github.Config

func init() {
	const def = "def"
	flag.StringVar(&conf.Mode, "mode", def,
		`Set the running mode of the program, requires an option argument.
	'list' a list of active issues, following the given search creiteria.
	'read' a designated issue, followed by the specific issue number.
	'edit' an existing issue.
	'raise' a new issue.
	'raw' test raw input.
	'lock' set the issue status to resolved.`)

	flag.StringVar(&conf.Mode, "m", def,
		"Raise a new issue (shorthand), requires an optional argument.")

	flag.StringVar(&conf.Lock, "lock", "",
		`Set the current issue's lock status, requires an optional argument.
	'topic' a list of active issue's, following the given search creiteria.
	'heated' a designated issue, followed by the specific issue number.
	'resolved' an existing issue.
	'spam' a new issue.`)

	flag.StringVar(&conf.Lock, "k", "",
		"Set the current issue's lock status (shorthand), requires an optional argument.")

	flag.StringVar(&conf.Login, "l", "", "set user name")
	flag.StringVar(&conf.Owner, "u", "", "set user name")
	flag.StringVar(&conf.Author, "a", "", "set user name")
	flag.StringVar(&conf.Org, "o", "", "set organisation name")
	flag.StringVar(&conf.Repo, "r", "", "set repo name")
	flag.StringVar(&conf.Token, "t", "", "set token")
	flag.StringVar(&conf.Editor, "e", "", "set editor")
	flag.StringVar(&conf.Number, "n", "", "set issue number")

	flag.BoolVar(&conf.Verbose, "v", false, "Verbose mode")
}

func main() {

	// Command line input.
	flag.Parse()

	// Setup programming for selected mode, in some cases the program mode
	// is altered here.
	err := github.SetState(&conf)

	switch conf.Mode {
	case "list":
		err = github.ListIssues(conf)
	case "read":
		err = github.ReadIssue(conf)
	case "raise":
		err = github.RaiseIssue(conf)
	case "edit":
		err = github.EditIssue(conf)
	case "lock":
		// TODO 1 set the correct URL.
		// TODO 2 implement writing issues.
		err = github.ReadIssue(conf)
	default:
		fmt.Println("Run with -h flag for user instructions.")
	}

	// Signal any program failure.
	if err != nil {
		fmt.Println(err)
	}
}
