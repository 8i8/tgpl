package main

import (
	"flag"
	"fmt"

	"tgpl/ex_04.11-github_client/github"
)

var (
	mode                         string
	Repo, Queries, Token, Editor *string
)

func init() {
	const def = "list"
	flag.StringVar(&mode, "mode", def,
		`Set the running mode of the program, requires an option argument.
	'list' a list of active issues, following the given search creiteria.
	'read' a designated issue, followed by the specific issue number.
	'edit' an existing issue.
	'raise' a new issue.
	'resolved' set the issue status to resolved.`)
	flag.StringVar(&mode, "m", def,
		"Raise a new issue (shorthand) requires an option argument.")

	Repo = flag.String("r", "golang/go", "set repo address")
	Token = flag.String("t", "", "set token")
	Editor = flag.String("e", "", "set editor")
	Queries = flag.String("q", "json decoder", "set a query")
}

func main() {

	flag.Parse()
	conf := github.Config{
		github.User{},
		*Token,
		*Editor,
		github.Request{github.User{}, *Repo, *Queries}}
	fmt.Println(conf)
	if err := github.LoadConfig(conf); err != nil {
		panic(err)
	}

	switch mode {
	case "list":
		results, _ := github.ListIssues(conf.Strings())
		github.PrintIssues(results)
	case "raise":
		issue := github.SetIssue()
		github.RaiseIssue(issue, conf.Token)
	case "read":
	case "edit":
		github.EditIssue()
	case "resolved":
		fmt.Println(mode)
	}
}
