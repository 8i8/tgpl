package main

import (
	"os"
	tgpl "tgpl/tgpl/pretty_printer"
)

func main() {
	tgpl.PrettyPrintDoc(os.Stdout, "https://8i8.fr", 4)
}
