package main

import (
	"flag"
	"fmt"

	"tgpl/ex_04.11-github_client/github"
)

var (
	conf                                                   github.Config
	Login, Owner, Author, Org, Repo, Token, Editor, Number *string
)

func init() {
	const def = "def"
	flag.StringVar(&conf.Mode, "mode", def,
		`Set the running mode of the program, requires an option argument.
	'list' a list of active issues, following the given search creiteria.
	'read' a designated issue, followed by the specific issue number.
	'edit' an existing issue.
	'raise' a new issue.
	'raw' test raw input.
	'resolved' set the issue status to resolved.`)
	flag.StringVar(&conf.Mode, "m", def,
		"Raise a new issue (shorthand) requires an option argument.")

	Login = flag.String("l", "", "set user name")
	Owner = flag.String("u", "", "set user name")
	Author = flag.String("a", "", "set user name")
	Org = flag.String("o", "", "set organisation name")
	Repo = flag.String("r", "", "set repo name")
	Token = flag.String("t", "", "set token")
	Editor = flag.String("e", "", "set editor")
	Number = flag.String("n", "", "set issue number")
}

// Store command line arguments in the config struct.
func setFlags() {
	flag.Parse()
	conf.Login = *Login
	conf.Owner = *Owner
	conf.Author = *Author
	conf.Org = *Org
	conf.Repo = *Repo
	conf.Token = *Token
	conf.Editor = *Editor
	conf.Number = *Number
	conf.Queries = flag.Args()
}

func main() {

	// Command line input.
	setFlags()

	// Setup programming for selected mode, in some cases the program mode
	// is altered here.
	github.SetState(&conf)

	switch conf.Mode {
	case "list":
		github.ListIssues(conf)
	case "read":
		github.ReadIssue(conf)
	case "raise":
		// TODO 2 implement writing issues.
		github.RaiseIssueOld(conf)
	case "edit":
		// TODO 1 set the correct URL.
		// TODO 2 implement writing issues.
		github.EditIssue()
	case "resolved":
		// TODO 1 set the correct URL.
		// TODO 2 implement writing issues.
		fmt.Println(conf.Mode)
	case "raw":
		// TODO 1
		github.ListIssues(conf)
	}
}
