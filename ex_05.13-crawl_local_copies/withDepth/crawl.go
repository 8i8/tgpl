package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"syscall"
	"tgpl/tgpl/tgpl"
)

type node struct {
	url      string
	path     string
	depth    int
	parent   *node
	children *[]node
}

func getPath(n node) node {
	var stack []string
	var l int
	v := &n
	_, file := path.Split(n.url)
	for v.parent != nil {
		v = v.parent
		stack = append(stack, file)
		l += len(file) + 1
		n.depth++
	}
	path := make([]byte, 0, l)
	for i := len(stack); i > 0; i-- {
		path = append(path, stack[i-1]+"/"...)
	}
	n.path = string(path)
	fmt.Println("depth:", n.depth, "path:", n.path)
	return n
}

func hasPrefix(prefixes []node, url string) bool {
	for i := range prefixes {
		if strings.HasPrefix(url, prefixes[i].url) {
			return true
		}
	}
	return false
}

func createFolders(url string) (string, string) {
	fname := "createFolders"
	dir, file := path.Split(tgpl.RemoveHttpPrefix(url))
	if dir == "" {
		dir = file + "/"
	}
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatalf(fname+": MkdirAll: %w", err)
		}
	} else if errors.Is(err, syscall.ENOTDIR) {
		return dir, file
	} else if err != nil {
		log.Fatal(fname+": ", err)
	}
	return dir, file
}

func writePage(url, dir, filename string, depth int) {
	fname := "writePage"
	if dir == "" {
		dir = filename + "/"
	}
	f, err := os.OpenFile(dir+filename, os.O_CREATE|os.O_WRONLY, 0644)
	if errors.Is(err, syscall.EISDIR) || errors.Is(err, syscall.ENOTDIR) {
		return
	}
	if err != nil {
		log.Fatal(fname+": opening file: ", err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	tgpl.PrettyPrintDoc(w, url, depth)
	w.Flush()
}

func crawl(n node) []node {
	fname := "crawl"
	n = getPath(n)
	list, err := tgpl.Extract(n.url)
	if err != nil {
		log.Printf(fname+": %s", err)
	}

	dir, file := createFolders(n.url)
	writePage(n.url, dir, file, n.depth)
	nodes := make([]node, len(list))
	for i := range list {
		nodes[i] = node{url: list[i], parent: &n}
	}
	n.children = &nodes
	return nodes
}

func breadthFirst(fn func(item node) []node, worklist []node) {
	seen := make(map[string]bool)
	prefix := worklist
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item.url] {
				// Only record pages within this domain.
				if hasPrefix(prefix, item.url) {
					worklist = append(worklist, fn(item)...)
					seen[item.url] = true
				}
			}
		}
	}
}

func main() {
	n := make([]node, len(os.Args[1:]))
	// Add http if not already present.
	for i, url := range os.Args[1:] {
		os.Args[i+1] = tgpl.CheckPrefix(url, "https://")
		n[i] = node{url: url}
	}
	breadthFirst(crawl, n)
}
