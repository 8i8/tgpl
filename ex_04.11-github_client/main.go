package main

import (
	"flag"
	"fmt"

	"tgpl/ex_04.11-github_client/github"
)

var conf github.Config
var flags github.FlagsIn

func init() {

	// Mode
	read := `Read mode:
	Set read mode to read an issue, if the editor flag is also supplied
	then the issue is opened in the designated editor for reading.`
	list := `List mode:
	Set list mode to display current issues in the given group.`
	edit := `Edit mode:
	Set edit mode to edit an issue, if the editor flag is also supplied
	then the issue is opened in the designated editor for reading.`
	raise := `Raise mode:
	Set raise mode to raise a new issue, requires a full repsitory address
	and Oauth2 authorisation.`
	lock := `Lock mode:
	Set lock mode to lock or to alter the current lock status of an issue,
	if no reason is given with the flag then the default 'resolved' is set.`
	unlock := `Unlock mode:
	Set unlock mode to unlock a previously close issue, requires autentication.`
	set := `Set mode:
	Set mode to define default editor and username.`
	verbose := `Verbose mode:
	Print information where available, explicitly describes the programs
	current state of operation.`

	flag.BoolVar(&flags.Read, "read", false, read)
	flag.BoolVar(&flags.List, "list", false, list)
	flag.BoolVar(&flags.Edit, "edit", false, edit)
	flag.BoolVar(&flags.Raise, "raise", false, raise)
	flag.StringVar(&conf.Lock, "lock", "", lock)
	flag.BoolVar(&flags.Unlock, "unlock", false, unlock)
	flag.BoolVar(&flags.Set, "set", false, set)
	flag.BoolVar(&flags.Verbose, "v", false, verbose)

	// Data
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

	flag.StringVar(&conf.User, "u", "", user)
	flag.StringVar(&conf.Author, "a", "", author)
	flag.StringVar(&conf.Org, "o", "", org)
	flag.StringVar(&conf.Repo, "r", "", repo)
	flag.StringVar(&conf.Number, "n", "", number)
	flag.StringVar(&conf.Token, "t", "", token)
	flag.StringVar(&conf.Editor, "d", "", editor)
}

func main() {

	// Command line input.
	flag.Parse()
	conf.Queries = flag.Args()

	// Setup programming for selected mode, in some cases the program mode
	// is altered here, as such we pass in a pointer.
	err := github.SetState(conf, flags)
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
