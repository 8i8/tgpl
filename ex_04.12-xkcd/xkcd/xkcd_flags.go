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
	xkcdget

SYNOPSIS
	xkcdget [flag]

	xkcdget is a comand line application for indexing and searching xkcd
	cartoons by their textual descriptions.
FLAGS
	xkcd	Displays the latest cartoon number and address.
	-v	Verbose mode.
	-u	Update comic database.
	-s	Search database.
	-t	Generate test database of size 'n'.
	-n	Display comic 'n' from database.
	-w	Display comic 'n' from web.
`
