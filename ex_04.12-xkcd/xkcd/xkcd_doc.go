/*
Package xkcd - Command line client and index for the xkcd comic website.

SYNOPSIS
	xkcd [flag][value][args]

DESCRIPTION
	xkcd is a command line client and search index for the online xkcd
	comic. Once a database of the repetoir has been made by scanning the
	site, comics can be located using the clients search function or
	browsed by number.

FLAGS
	xkcd -s hello world

	-s	Search for <args> amongst the comic descriptions in the local database.

	xkcd -n 571

	-n	Display the description of comic 'n' from the database index.
	-w	Download and display the description of comic 'n' from the web.

	xkcd -u -s hello world

	-u	Update first with the latest comics descriptions.
	-v	Verbose mode, for a detailed output of the programs actions.
	-help	Prints out the programs help file.
	-h

HTTP REQUESTS

	Information of the sites API
		https://xkcd.com/json.html

	http://xkcd.com/info.0.json (current comic)
	http://xkcd.com/614/info.0.json (comic #614)

*/

package xkcd

import (
	"flag"
	"fmt"
)

// helpFlag is a flag that is set to signal help is to be printed.
var helpFlag bool

// SetupFlags provides custom help documentation.
func SetupFlags(flag *flag.FlagSet) {
	flag.Usage = func() {
		fmt.Println(help)
	}
}

var help = `
NAME
	xkcd

SYNOPSIS
	xkcd [flag][value][args]

DESCRIPTION
	xkcd is a command line client and search index for the online xkcd
	comic. Once a database of the repetoir has been made by scanning the
	site, comics can be located using the clients search function or
	browsed by number.

FLAGS
	xkcd -s hello world

	-s	Search for <args> amongst the comic descriptions in the local database.

	xkcd -n 571

	-n	Display the description of comic 'n' from the database index.
	-w	Download and display the description of comic 'n' from the web.

	xkcd -u -s hello world

	-u	Update first with the latest comics descriptions.
	-v	Verbose mode, for a detailed output of the programs actions.
	-help	Prints out the programs help file.
	-h

`
