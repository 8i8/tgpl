package ds

import (
	"fmt"
	"os"
	"strconv"
	"unicode"

	"tgpl/ex_04.12-xkcd/msg"
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
	d |= VERBOSE
}

// buildTrieFromMap constructs a search trie from a search map.
func buildTrieFromMap(t *Trie, m MList) {

	for word, indices := range m {
		t.Add(word, indices)
	}
}

func InitaliseTrie(m MList, addr string) (*Trie, error) {

	t := new(Trie)
	if _, err := os.Stat(addr); err == nil {
		err = deserialiseTrieFromFile(t, addr)
		if err != nil {
			return t, fmt.Errorf("t.Deserialise: %v", err)
		}
	} else {
		buildTrieFromMap(t, m)
		serialiseTrieToFile(t, addr)
	}

	return t, nil
}

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
 *  Trie Serialisation
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// writeList writes the data list from a node to the given bufio buffer.
func writeList(n *node, w *msg.ErrWriter) {

	w.WriteRune(' ')
	for i, id := range n.list {
		w.WriteInt(id)
		// Add a comma between all integers, except for the last.
		if i != len(n.list)-1 {
			w.WriteRune(',')
		}
	}
	w.WriteRune(' ')

	// If there is a writer error, indicate the function.
	if w.Err != nil {
		w.Err = fmt.Errorf("serialise: %v", w.Err)
	}
}

// serialiseTrie is a helper function for SerialiseTrie.
func serialiseTrie(n *node, w *msg.ErrWriter) {

	// Denote a leaf node in the serialisation.
	if n.next == nil {
		w.WriteRune('¬')
		return
	}

	// If there is associated data, write it.
	if len(n.list) > 0 {
		writeList(n, w)
	}

	// Go recursively to all other nodes in the trie.
	i := 0
	for r, next := range n.next {
		i++
		w.WriteRune(r)
		serialiseTrie(next, w)
	}

	// If there is a writer error, indicate the function.
	if w.Err != nil {
		w.Err = fmt.Errorf("serialise: %v", w.Err)
	}
}

// SerialiseTrie writes the content of the trie structure to disk, used as a cache
// for the trie data structure.
func (t *Trie) SerialiseTrie(w *msg.ErrWriter) error {

	if d&VERBOSE > 0 {
		fmt.Printf("ds: trie serialisation started ...\n")
	}

	// Begin serialisation of trie data structure.
	serialiseTrie(t.start, w)
	if w.Err != nil {
		return fmt.Errorf("serialise: %v", w.Err)
	}

	if d&VERBOSE > 0 {
		fmt.Printf("ds: ... trie serialisation done\n")
	}

	return nil
}

// gerIds read until the next \t is reached, record all integer id's until
// then.
func getIds(n *node, rw *msg.ErrReader) (rune, *msg.ErrReader, error) {

	var err error
	var str []rune
	var r rune
	i := 0

	// read until the next \t.
	for r != ' ' {
		r, _, err = rw.ReadRune()
		// The integers are space separated.
		for r != ',' {
			str = append(str, r)
			r, _, err = rw.ReadRune()
			if err != nil {
				return r, rw, err
			}
		}
		if err != nil {
			return r, rw, err
		}
		// Change to an integer type and store in the nodes list.
		i, err = strconv.Atoi(string(str))
		n.list = append(n.list, i)
		str = str[:0]
	}
	// The last one.
	i, err = strconv.Atoi(string(str))
	n.list = append(n.list, i)
	if err != nil {
		return r, rw, fmt.Errorf("strconv.Atoi: %v", err)
	}

	r, _, err = rw.ReadRune()
	if err != nil {
		return r, rw, err
	}

	return r, rw, nil
}

// deserialiseTrie is a helper function for Deserialise.
func deserialiseTrie(n *node, rw *msg.ErrReader) (rune, error) {

	r, _, err := rw.ReadRune()
	if err != nil {
		return r, err
	}

	// If the leaf marker has been read, return and continue onto the next
	// branch.
	if r == '\n' {
		print('\n')
		return r, err
	}

	// If the node requires linked information, add it.
	if r == ' ' {
		r, rw, err = getIds(n, rw)
	}

	// Until a leaf node is reached, keep putting the next rune in a node
	// and moving to it.
	for r != '\n' && r != ' ' {
		n.addRuneToMap(r)
		r, err = deserialiseTrie(n.next[r], rw)
	}

	// // Keep putting the next rune into a node and moving to that node.
	// for _, r := range buf {
	// 	n.addRuneToMap(r)
	// 	print(r)
	// 	err = deserialise(n.next[r], rw)
	// }

	return r, err
}
