package main

import (
	"os"

	"tgpl/ex_05.01-recursive_findlinks/fl"
)

func main() {
	fl.Findlinks(os.Stdin)
	//fl.FindlinksOrig(os.Stdin)
}
