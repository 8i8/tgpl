package main

import (
	"bufio"
	"fmt"
	"os"
)

type myMap struct {
	c int
	m map[string]int
}

func main() {
	counts := make(map[string]*myMap)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "empty", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts)
			f.Close()
		}
	}
	for line, files := range counts {
		for file, n := range files.m {
			if n >= 1 {
				fmt.Printf("%d of %d \t%s in %s\n", n, files.c, line, file)
			}
		}
	}
}

func countLines(f *os.File, name string, counts map[string]*myMap) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			m := myMap{}
			m.m = make(map[string]int)
			counts[input.Text()] = &m
		}
		counts[input.Text()].m[name]++
		counts[input.Text()].c++
	}
	// NOTE: ignoring potential errors from input.Err()
}
