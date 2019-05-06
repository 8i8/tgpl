package github

import (
	"flag"
	"fmt"
)

// Flag that is set to signal help is to be printed.
var helpflag bool

// SetupFlags initiates custom flag usage.
func SetupFlags(flag *flag.FlagSet) {

	flag.Usage = func() {
		fmt.Println(help)
	}
}

// Help: the text output when the -help or -h flags are raised.
var help = `
NAME
	github

SYNOPSIS
	github [mode] [u|a|o] <value> [r] <value> [options]

MODES
	github -[mode]

	-read	Read an existing issue.
	-list	List all issues for a specific repo or user.
	-edit	Edit an existing issue.
	-raise	Raise a new issue.
	-lock	Lock an issue.
	-unlock	Unlock an issue.

FLAGS
	github	-[flag]

	-v	Verbose mode, gives detailed description of the programs actions.
	-h
	-help	Prints out the programs help file.

	github	-[flag] [value]

	-u	User name.
	-a	Author.
	-o	Organisation name.
	-r	Repository name.
	-n	Issue number, requires that author and repository also be defined.
	-t	OAuth2 token.
	-d	External editor launch command.
	-l	Provide a reason for locking.

	[lock rational]

		* off-topic
		* too heated
		* resolved
		* spam
`
