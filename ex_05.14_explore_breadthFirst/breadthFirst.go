package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type node struct {
	file     os.FileInfo
	path     string
	parent   *node
	children *[]node
}

func getPath(n node) string {
	var stack []string
	var l int
	v := &n
	for v.parent != nil {
		v = v.parent
		stack = append(stack, v.file.Name())
		l += len(v.file.Name()) + 1
	}
	path := make([]byte, 0, l)
	for i := len(stack); i > 0; i-- {
		path = append(path, stack[i-1]+"/"...)
	}
	return string(path)
}

var errNilPointer = errors.New("nil pointer exception")

func filesystem(n node) []node {
	fname := "filesystem"
	if n.file == nil {
		log.Fatalf(fname+": %s", errNilPointer)
	}
	path := getPath(n)
	fmt.Printf("%s%s\n", path, n.file.Name())
	if !n.file.IsDir() {
		return nil
	}
	files, err := ioutil.ReadDir(path + n.file.Name())
	if err != nil {
		log.Fatalf(fname+": %s", err)
	}

	nodes := make([]node, 0, len(files))
	for _, f := range files {
		nodes = append(nodes, node{file: f, parent: &n})
	}
	n.children = &nodes
	return nodes
}

func breadthFirst(fn func(node) []node, worklist []node) {
	seen := make(map[os.FileInfo]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item.file] {
				worklist = append(worklist, fn(item)...)
				seen[item.file] = true
			}
		}
	}
}

func main() {
	f, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	n := make([]node, 0, len(f))
	for i := range f {
		n = append(n, node{file: f[i]})
	}
	breadthFirst(filesystem, n)
}
