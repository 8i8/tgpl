package xkcd

import (
	"flag"
	"fmt"
)

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
	-u	Updates the database.
	-d	Displays the latest cartoon's description.
	-r	Remakes the list of cartoons.
	-s	Search the xkcd database.
`
