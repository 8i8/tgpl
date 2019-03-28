package main

import (
	"fmt"

	"tgpl/ex_04.12-xkcd/xkcd"
)

func main() {

	comics, err := xkcd.LoadDatabase()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Printf("%+v\n", comics.Print(0))
}
