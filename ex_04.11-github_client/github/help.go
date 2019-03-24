package github

import (
	"flag"
)

var Conf Config
var FlagsIn FlagsInStruct
var Helpflag bool

var Help = `
NAME
	github

SYNOPSIS
	github [mode] [name|repo|number] [options]

MODES
	github -[mode]

	-read	Read an existing issue.
	-list	List all issues for a specific repo or user.
	-edit	Edit an existing issue.
	-raise	Raise a new issue.
	-lock	Lock an issue.
	-unlock	Unlock an issue.
	-set	Set default user name and editor.

	github -[mode] [value]

	-lock	Lock an issue.

FLAGS
	github	-[flag]

	-v	Verbose mode, gives detailed description of the programs actions.
	-h
	-help	Print out the programs help file.

	github	-[flag] [value]

	-u	User name.
	-a	Author.
	-o	Organisation name.
	-r	Repository name.
	-n	Issue number, requires that author and repository also be defined.
	-t	OAuth2 token.
	-d	External editor launch command.
	-l	Provide a reason for locking.
`

func init() {

	// Mode
	flag.BoolVar(&FlagsIn.Read, "read", false, "")
	flag.BoolVar(&FlagsIn.List, "list", false, "")
	flag.BoolVar(&FlagsIn.Edit, "edit", false, "")
	flag.BoolVar(&FlagsIn.Raise, "raise", false, "")
	flag.BoolVar(&FlagsIn.Lock, "lock", false, "")
	flag.BoolVar(&FlagsIn.Unlock, "unlock", false, "")
	flag.BoolVar(&FlagsIn.Set, "set", false, "")
	flag.BoolVar(&FlagsIn.Verbose, "v", false, "")
	flag.BoolVar(&Helpflag, "h", false, "")
	flag.BoolVar(&Helpflag, "help", false, "")

	// Data
	flag.StringVar(&Conf.User, "u", "", "")
	flag.StringVar(&Conf.Author, "a", "", "")
	flag.StringVar(&Conf.Org, "o", "", "")
	flag.StringVar(&Conf.Repo, "r", "", "")
	flag.StringVar(&Conf.Number, "n", "", "")
	flag.StringVar(&Conf.Token, "t", "", "")
	flag.StringVar(&Conf.Editor, "d", "", "")
	flag.StringVar(&Conf.Reason, "l", "", "")
}
