package main

import (
	"flag"
	"fmt"

	"tgpl/ex_04.11-github_client/github"
)

var (
	conf                            github.Config
	Login, Org, Repo, Token, Editor *string
)

func init() {
	const def = "list"
	flag.StringVar(&conf.Mode, "mode", def,
		`Set the running mode of the program, requires an option argument.
	'list' a list of active issues, following the given search creiteria.
	'read' a designated issue, followed by the specific issue number.
	'edit' an existing issue.
	'raise' a new issue.
	'resolved' set the issue status to resolved.`)
	flag.StringVar(&conf.Mode, "m", def,
		"Raise a new issue (shorthand) requires an option argument.")

	Login = flag.String("u", "", "set user name")
	Org = flag.String("o", "", "set organisation name")
	Repo = flag.String("r", "", "set repo name")
	Token = flag.String("t", "", "set token")
	Editor = flag.String("e", "", "set editor")
}

// Store command line arguments in the config struct.
func setFlags() {
	flag.Parse()
	conf.Login = *Login
	conf.Org = *Org
	conf.Repo = *Repo
	conf.Token = *Token
	conf.Editor = *Editor
	conf.Queries = flag.Args()
}

func main() {

	setFlags()
	//if err := github.LoadConfig(conf); err != nil {
	//	panic(err)
	//}

	switch conf.Mode {
	case "list":
		// TODO 1 set the correct URL.
		results, _ := github.ListIssues(conf)
		github.PrintIssues(results)
	case "read":
		// TODO 1 set the correct URL.
	case "raise":
		// TODO 2 impliment writing issues.
		issue := []github.Issue{}
		github.RaiseIssue(issue, conf)
	case "edit":
		// TODO 1 set the correct URL.
		// TODO 2 impliment writing issues.
		github.EditIssue()
	case "resolved":
		// TODO 1 set the correct URL.
		// TODO 2 impliment writing issues.
		fmt.Println(conf.Mode)
	}
}
