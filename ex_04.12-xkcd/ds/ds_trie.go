package ds

import (
	"fmt"
	"unicode"
)

type hash map[rune]*node

type Trie struct {
	wc    int // Word count.
	nc    int // Node count.
	start *node
}

type node struct {
	count int // The number of times that the word has been added.
	next  hash
	list  []uint
}

type Data struct {
	//id        int
	found bool
	word  string
	Links []uint
}

type debug = uint

var d debug

const (
	VERBOSE debug = 1 << iota
	ADD
	SEARCHWORDS
	GOSEARCHFIRTSRUNE
	GETALLWORDS
	GOGETALLWORDS
	BTREEADD
	BTREECOUNT
)

func init() {
	//d |= VERBOSE
	//d |= ADD
	//d |= SEARCHWORDS
	//d |= GOSEARCHFIRTSRUNE
	//d |= GETALLWORDS
	//d |= GOGETALLWORDS
}

func (t *Trie) Add(s string, list []uint) {

	// If first use, initialise.
	if t.start == nil {
		t.start = new(node)
	}
	n := t.start

	// For each rune in the string, add to the trie, create nodes where required.
	for _, r := range s {
		r = unicode.ToLower(r)
		if n.next == nil {
			n.next = make(hash)
		}
		if n.next[r] == nil {
			n.next[r] = new(node)
			t.nc++
		}
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

// addData combines two data structs.
func addData(data, newdata *Data) *Data {

	if newdata == nil {
		return data
	}

	if newdata.found {
		data.found = true
	}

	for _, l := range newdata.Links {
		data.Links = append(data.Links, l)
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

// goGetAllWords
func (n *node) goGetAllWords(ch chan<- []uint) {

	ch1 := make(chan []uint)
	var list []uint

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
			data.Links = append(data.Links, l)
		}
	}

	ch := make(chan []uint)

	// For every rune in the map, call goGetAllWords.
	for _, n1 := range n.next {
		go n1.goGetAllWords(ch)
	}

	// For every returned btree array, add the btrees to the data struct.
	for range n.next {
		list := <-ch
		for _, l := range list {
			data.Links = append(data.Links, l)
		}
	}

	return data
}

func (n *node) goSearchFirstRune(s string, ch chan<- Data) {

	data := Data{}
	data.word = s

	if n == nil {
		ch <- data
		return
	}

	// For each character in the hashmap check the first rune against the
	// rune key of the hashmap if it matches, run CheckGopher to search for
	// the given string.
	ch1 := make(chan Data)
	for r, next := range n.next {
		if r == rune(s[0]) {
			// remove the first rune fron the string, it has already been noted.
			go next.goCheckWord(data, s[1:], ch1)
		}
	}

	// If found go get all following words.
	ch2 := make(chan Data)
	for r, next := range n.next {
		if r == rune(s[0]) {
			data = <-ch1
		}
		// Moveing forwards, check again.
		go next.goSearchFirstRune(s, ch2)
	}
	for range n.next {
		newdata := <-ch2
		addData(&data, &newdata)
	}

	ch <- data
}

// SearchWords starts a multi goroutine trie search for the given strings.
func (t *Trie) SearchWords(s []string) []Data {

	n := t.start
	ch := make(chan Data)
	results := []Data{}

	// Send a goroutine for every word in the list.
	for _, word := range s {
		go n.goSearchFirstRune(word, ch)
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
			printList(data.Links)
			fmt.Println()
		}
	}

	return results
}
