package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path"
	"strings"
	"syscall"
)

func hasPrefix(prefixes []string, url string) bool {
	for i := range prefixes {
		if strings.HasPrefix(url, prefixes[i]) {
			return true
		}
	}
	return false
}

func breadthFirst(fn func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	prefix := worklist
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				// Only record pages within this domain.
				if hasPrefix(prefix, item) {
					worklist = append(worklist, fn(item)...)
					seen[item] = true
				}
			}
		}
	}
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

func writePage(url, dir, filename string) {
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
	tgpl.PrettyPrint(w, url)
	w.Flush()
}

func crawl(url string) []string {
	list, err := tgpl.Extract(url)
	if err != nil {
		log.Print(err)
	}
	dir, file := createFolders(url)
	writePage(url, dir, file)
	return list
}

func main() {
	// Add http if not already present.
	for i, url := range os.Args[1:] {
		os.Args[i+1] = tgpl.CheckPrefix(url, "https://")
	}
	breadthFirst(crawl, os.Args[1:])
}
