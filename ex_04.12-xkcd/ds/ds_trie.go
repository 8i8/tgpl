package ds

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
	count int // The number of times that the word has been added.
	next  hash
	data  *BtreeNode
}

func add(t *Trie, s string, d BtreeNode) {

	// If first use, initialise.
	if t.start == nil {
		t.start = new(node)
	}
	n := t.start

	// For each rune in the string, add to the trie creating nodes when and
	// if they are required.
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

	// If the word is not already present, the word end flag false count it
	// and set the words end.
	if n.count == 0 {
		t.wc++
	}
	n.count++
	BtreeAdd(n.data, d)
}

// Add inserts given string into the trie also counting occurances of that
// string and storing the index of its origin.
func (t *Trie) Add(s string, index uint) *Trie {
	data := BtreeNode{}
	data.index = index
	add(t, s, data)
	return t
}

// Check verifies using multiple goroutines if the given string exists in the
// trie as a complete word, from the given node.
func (n *node) GopherCheck(s string, ch chan<- *BtreeNode) {

	// remove the first rune it has already been noted, it was the key of
	// the map that calls this function.
	s = string([]rune(s)[1:])
	empty := &BtreeNode{}

	for _, r := range s {
		if n.next[r] == nil {
			ch <- empty
			return
		}
		n = n.next[r]
	}
	ch <- n.data
	return
}

// subWord is a helper function for SubWordSearch.
func (n *node) subWord(s string, ch1 chan<- *BtreeNode) {

	ch := make(chan *BtreeNode)
	btree := &BtreeNode{}

	if n == nil {
		ch1 <- btree
		return
	}

	// For each character in the hashmap check the first rune against the
	// rune key of the hashmap if it matches, run CheckGopher to search for
	// the given string.
	for r, c := range n.next {
		if r == rune(s[0]) {
			go c.GopherCheck(s, ch)
		}
	}
	for r, _ := range n.next {
		if r == rune(s[0]) {
			btree := <-ch
			if btree != nil && btree.index > 0 {
				ch1 <- btree
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
		btree := <-ch
		if btree != nil && btree.index > 0 {
			ch1 <- btree
			return
		}
	}

	ch1 <- btree
}

// SubWordSearch starts a multi goroutine trie search for the given string.
func (t *Trie) SubWordSearch(s []string) []*BtreeNode {

	n := t.start
	ch := make(chan *BtreeNode)
	var results []*BtreeNode

	// Send a goroutine.
	for _, word := range s {
		go n.subWord(word, ch)
	}
	for range s {
		btree := <-ch
		fmt.Println(btree)
		results = append(results, btree)
	}

	return results
}
