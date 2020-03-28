package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello, World!"
	s = expand(s, bigWorld)
	fmt.Println(s)
}

func bigWorld(s string) string {
	foo := "World"
	return strings.ReplaceAll(s, foo, strings.ToUpper(foo))
}

func expand(s string, f func(s string) string) string {
	return f(s)
}
