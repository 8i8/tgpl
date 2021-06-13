package main

import (
	"fmt"
	"os"
)

func panicRecovery() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, r)
		}
	}()
	panic("help")
}

func main() {
	panicRecovery()
	fmt.Println("still running")
}
