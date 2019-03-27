package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func echo1(args []string) {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + strconv.Itoa(i) + " " + args[i]
		sep = " \n"
	}
	fmt.Println(s)
}

func echo2(args []string) {
	s, sep := "", ""
	for i, arg := range args[:] {
		s += sep + strconv.Itoa(i) + arg
		sep = " \n"
	}
	fmt.Println(s)
}

func echo3(args []string) {
	fmt.Println(strings.Join(args[:], " \n"))
}

func main() {
	echo1(os.Args)
	echo2(os.Args)
	echo3(os.Args)
}
