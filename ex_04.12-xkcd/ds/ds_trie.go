package ds

import (
	"fmt"
	"sort"
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
	list  []int
}

type Data struct {
	found bool
	word  string
	List  []int
}

type debug = int

var d debug

const (
	VERBOSE debug = 1 << iota
)

func init() {
	// d |= VERBOSE
}

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  Build
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// addRuneToMap adds a rune to the hashmap adding a new hashmap if required.
func (n *node) addRuneToMap(r rune) {
	if n.next == nil {
		n.next = make(hash)
	}
	if n.next[r] == nil {
		n.next[r] = new(node)
	}
}

func (t *Trie) Add(s string, list []int) {

	// If first use, initialise.
	if t.start == nil {
		t.start = new(node)
	}
	n := t.start

	// For each rune in the string, add to the trie, create nodes where required.
	for _, r := range s {
		r = unicode.ToLower(r)
		n.addRuneToMap(r)
		n = n.next[r]
	}

	// Add data to the last node of the word.
	// If the word is not already present count it, set the
	// word end by augmenting the count.
	if n.count == 0 {
		t.wc++
	}
	n.count++
	// Add the nodes index list.
	n.list = list
}

func (t *Trie) BuildFromMap(m MList) {

	for word, indices := range m {
		t.Add(word, indices)
	}
}

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  Trie Search
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// addData combines two data structs.
func addData(data, newdata *Data) *Data {

	if newdata == nil {
		return data
	}

	if newdata.found {
		data.found = true
	}

	for _, l := range newdata.List {
		data.List = append(data.List, l)
	}
	return data
}

// goGetAllWords recursive fuction for searching for all word termination
// points in the trie branch.
func (n *node) goGetAllWords(ch chan<- []int) {

	ch1 := make(chan []int)
	var list []int

	// If the current node has linked data, store it.
	if n.list != nil {
		for _, l := range n.list {
			list = append(list, l)
		}
	}

	// For every rune in the hashmap move forwards and look for word
	// endings add any new found btrees to the list.
	for _, next := range n.next {
		go next.goGetAllWords(ch1)
	}
	for range n.next {
		for _, newlist := range <-ch1 {
			list = append(list, newlist)
		}
	}

	ch <- list
	return
}

// getAllWords starts a multi goroutine trie search for all following words.
func getAllWords(n *node, data Data) Data {

	// If the currend node has linked data, store it.
	if n.list != nil {
		for _, l := range n.list {
			data.List = append(data.List, l)
		}
	}

	ch := make(chan []int)

	// For every rune in the map, call goGetAllWords.
	for _, n1 := range n.next {
		go n1.goGetAllWords(ch)
	}

	// For every returned btree array, add the btrees to the data struct.
	for range n.next {
		list := <-ch
		for _, l := range list {
			data.List = append(data.List, l)
		}
	}

	return data
}

// goCheckWord verifies using multiple goroutines if the given string exists in
// the trie as a complete word, from the given node.
func (n *node) goCheckWord(d Data, s string, ch chan<- Data) {

	// Keep the address of the first node, this is required when calling
	// getAllWords, so as to retrieve the linked ntree data that relates to
	// the current word.
	for _, r := range s {
		if n.next[r] == nil {
			ch <- d
			return
		}
		n = n.next[r]
	}
	d.found = true

	d = getAllWords(n, d)
	ch <- d
	return
}

// goSearchForFirstRune itterates over and launches a new word search upon
// finding the first rune of the given string in the trie branches.
func (n *node) goSearchForFirstRune(s string, ch chan<- Data) {

	data := Data{}
	data.word = s

	if n == nil {
		ch <- data
		return
	}

	// For each character in the hashmap check the first rune against the
	// rune key of the hashmap if it matches, run goCheckWord to search for
	// the given string.
	ch1 := make(chan Data)
	for r, next := range n.next {
		if r == rune(s[0]) {
			// remove the first rune from the string, it has already been noted.
			go next.goCheckWord(data, s[1:], ch1)
		}
	}

	// If found go get all following words.
	ch2 := make(chan Data)
	for r, next := range n.next {
		if r == rune(s[0]) {
			data = <-ch1
		}
		// Moving forwards, check again.
		go next.goSearchForFirstRune(s, ch2)
	}
	for range n.next {
		newdata := <-ch2
		addData(&data, &newdata)
	}

	ch <- data
}

// SearchTrie starts a multi goroutine trie search for the given strings.
func (t *Trie) SearchTrie(s []string) []Data {

	n := t.start
	ch := make(chan Data)
	results := []Data{}

	// Send a goroutine for every word in the list.
	for _, word := range s {
		go n.goSearchForFirstRune(word, ch)
	}
	for i, _ := range s {
		data := <-ch
		results = append(results, data)
		if d&VERBOSE > 0 {
			printState(results[i])
		}
	}

	if d&VERBOSE > 0 {
		for _, data := range results {
			fmt.Printf("%s: ", data.word)
			printList(data.List)
			fmt.Println()
		}
	}

	return results
}

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  Expand Trie
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// goExpandTrie is a recursive helper function for ExpandTrie.
func goExpandTrie(n *node, suffix string, ch1 chan<- []string) {

	// To receive all suffix that are to be returned.
	var wordlist []string

	// This is the end of a word, return the stem.
	if len(n.list) > 0 {
		wordlist = append(wordlist, suffix)
	}

	// This is a leaf node, stop here.
	if n.next == nil {
		ch1 <- wordlist
		return
	}

	// Launch a goroutine for every hash bucket in the map.
	ch2 := make(chan []string)
	for r, _ := range n.next {
		// Add the rune to the current suffix and reiterate.
		go goExpandTrie(n.next[r], suffix+string(r), ch2)
	}
	// Concatenate any returned stems as suffix to the current word.
	for range n.next {
		stemlist := <-ch2
		for _, stem := range stemlist {
			wordlist = append(wordlist, stem)
		}
	}

	// Send the composite stems to be added as suffix to the root stems.
	ch1 <- wordlist
}

// // ExpandTrie returns a sorted list of all words included in the trie.
func (t *Trie) ExpandTrie() []string {

	// Retrieve first node.
	n := t.start

	// Prepare to receive the entire list of words stored in the trie.
	var results []string

	// For each character in the hashmap run the goPrint command.
	ch := make(chan []string)
	for r, _ := range n.next {
		// Add the first rune to the function to start reconstruction.
		go goExpandTrie(n.next[r], string(r), ch)
	}

	// Compile a list of all returned words.
	for range n.next {
		wordlist := <-ch
		for _, word := range wordlist {
			results = append(results, word)
		}
	}
	// Sort the results.
	sort.Strings(results)

	return results
}
