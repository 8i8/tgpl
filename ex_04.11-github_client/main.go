package main

import (
	"flag"
	"fmt"

	"tgpl/ex_04.11-github_client/github"
)

var conf github.Config

func init() {
	flag.StringVar(&conf.User, "u", "", "set user name")
	flag.StringVar(&conf.Author, "a", "", "set author's name")
	flag.StringVar(&conf.Org, "o", "", "set organisation name")
	flag.StringVar(&conf.Repo, "r", "", "set repo name")
	flag.StringVar(&conf.Number, "n", "", "set issue number")
	flag.StringVar(&conf.Token, "t", "", "provide Oauth token")
	flag.StringVar(&conf.Editor, "d", "", "set prefered editor")

	flag.BoolVar(&conf.Lock, "k", false, "Lock/unlock mode")
	flag.BoolVar(&conf.Verbose, "v", false, "Verbose mode")
	flag.BoolVar(&conf.Edit, "e", false, "Edit mode")
}

func main() {

	// Command line input.
	flag.Parse()

	// Setup programming for selected mode, in some cases the program mode
	// is altered here.
	err := github.InitState(&conf)

	switch conf.Mode {
	case github.MoList:
		err = github.ListIssues(conf)
	case github.MoRead:
		err = github.ReadIssue(conf)
	case github.MoEdit:
		err = github.EditIssue(conf)
	case github.MoLock:
		err = github.ReadIssue(conf)
	case github.MoRaise:
		err = github.RaiseIssue(conf)
	default:
		fmt.Println("Run with -h flag for user instructions.")
	}

	// Signal any program failure.
	if err != nil {
		fmt.Println(err)
	}
}
