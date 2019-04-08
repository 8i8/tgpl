package trie

import (
	"fmt"
	"unicode"
)

type hash map[rune]*node

type Trie struct {
	wc    int // Word count.
	start *node
}

type node struct {
	stop bool // End of a word.
	next hash
}

// Add inserts given string into the trie.
func (t *Trie) Add(s string) {

	if t.start == nil {
		t.start = new(node)
	}
	n := t.start

	for _, r := range s {
		r = unicode.ToLower(r)
		if n.next == nil {
			n.next = make(hash)
		}
		if n.next[r] == nil {
			n.next[r] = new(node)
		}
		n = n.next[r]
	}

	// If the word is not already present, count it.
	if !n.stop {
		t.wc++
		n.stop = true
	}
}

// Check verifies using a single routine if given string exists in the trie as
// a complete word from the given node.
func (n *node) Check(s string) bool {

	l := 0
	for _, r := range s {
		l++
		if n.next[r] == nil {
			break
		}
		n = n.next[r]
	}
	if l != len(s) {
		return false
	}
	return true
}

func getAllFollowing(t *Trie) {
}

// Check verifies using multiple goroutines if the given string exists in the
// trie as a complete word, from the given node.
func (n *node) GopherCheck(s string, ch chan<- bool) {

	// remove the first rune it has already been noted, it was the key of
	// the map that calls this function.
	s = string([]rune(s)[1:])

	for _, r := range s {
		if n.next[r] == nil {
			ch <- false
			return
		}
		n = n.next[r]
	}
	ch <- true
	return
}

// SubWord checks whether the given string exists as a sub word within the
// trie.
func (n *node) subWord(s string, ch1 chan<- bool) {

	ch := make(chan bool)

	if n == nil {
		ch1 <- false
		return
	}

	// For each character in the hashmap check the first rune against the
	// rune hey of the hashmap if it matches, run CheckGopher to search for
	// the given string.
	for r, c := range n.next {
		if r == rune(s[0]) {
			go c.GopherCheck(s, ch)
		}
	}
	for r, _ := range n.next {
		if r == rune(s[0]) {
			out := <-ch
			if out {
				ch1 <- out
				return
			}
		}
	}

	// Move forwards one character and repeat for every character in the
	// map.
	for _, n1 := range n.next {
		go n1.subWord(s, ch)
	}
	for range n.next {
		out := <-ch
		if out {
			ch1 <- out
			return
		}
	}

	ch1 <- false
	return
}

// SubWord starts a multi goroutine trie search for the given string.
func (t *Trie) SubWord(s string) {

	n := t.start
	ch := make(chan bool)
	var out bool

	// Send a goroutine.
	go n.subWord(s, ch)
	out = <-ch
	printState(s, out)
}

func printState(s string, out bool) {

	if out {
		fmt.Printf("string '%s' found.\n", s)
	} else {
		fmt.Printf("string '%s' not found.\n", s)
	}
}

func main() {

	trie := Trie{}

	trie.Add("hello")
	trie.SubWord("l")
	trie.Add("world")
	trie.SubWord("l")
}
