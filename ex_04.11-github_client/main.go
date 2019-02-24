package main

import (
	"flag"
	"fmt"

	"tgpl/ex_04.11-github_client/github"
)

var conf github.Config
var flags github.FlagsIn

func init() {

	user := `Login user name:
	The name used when logging into the github API, searches and requests
	made that do not have the "author" specified will use this value in the
	search.`
	author := `Author's name: 
	The author's of the requested repository's login name.`
	org := `Organisation name:
	The name of the organisation in which the search for the issue will be made.`
	repo := `Repository name:
	The name of the repository in which to search for the requested issue.`
	number := `Issue number:
	The identification number of the requested repository issue.`
	token := `Oauth2 token:
	Specify the oauth token to obtain access privileges for editing issues.`
	editor := `Designated editor:
	Specify an editor to use for your issue editing request.`
	lock := `Lock mode:
	Set lock mode, edit the lock status of an issue, requires user authentication.`
	raise := `Raise mode:
	Set raise mode to raise a new issue, requires a full repsitory address
	and Oauth2 authorisation.`
	verbose := `Verbose mode:
	Print information where available, explicitly describes the programs
	current state of operation.`
	edit := `Edit mode ~ edit an existing issue.
	A paragraph of text whos sole purpose is that of filling space, as
	such; It is far less of a concern to me what it says as the actual
	amount of space that it occupies.`

	flag.StringVar(&conf.User, "u", "", user)
	flag.StringVar(&conf.Author, "a", "", author)
	flag.StringVar(&conf.Org, "o", "", org)
	flag.StringVar(&conf.Repo, "r", "", repo)
	flag.StringVar(&conf.Number, "n", "", number)
	flag.StringVar(&conf.Token, "t", "", token)
	flag.StringVar(&conf.Editor, "d", "", editor)
	flag.BoolVar(&flags.Lock, "k", false, lock)
	flag.BoolVar(&flags.Edit, "e", false, edit)
	flag.BoolVar(&flags.Raise, "x", false, raise)
	flag.BoolVar(&flags.Raw, "w", false, lock)
	flag.BoolVar(&flags.Verbose, "v", false, verbose)
}

func main() {

	flag.Parse()
	// Command line input.
	conf.Queries = flag.Args()

	// Setup programming for selected mode, in some cases the program mode
	// is altered here, as such we pass in a pointer.
	err := github.SetState(&conf, flags)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Run the program with given configuration.
	err = github.Run(conf)
	if err != nil {
		fmt.Println(err)
		return
	}
}
