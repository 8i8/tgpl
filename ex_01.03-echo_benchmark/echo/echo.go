package echo

import (
	"strings"
)

func echo1(args []string) string {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return s
}

func echo2(args []string) string {
	s, sep := "", ""
	for _, arg := range args[:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func echo3(args []string) string {
	return strings.Join(args[:], " ")
}
