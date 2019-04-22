package ds

import "fmt"

var mark bool
var prevList []uint

func printTrieList(n *node) {

	// For all unicode char in the map.
	for r, next := range n.next {

		fmt.Printf("%c", r)

		if mark {
			printList(prevList)
			mark = false
		}

		// If the end of a branch, print a new line char.
		if next.next == nil {
			fmt.Print("\n")
		}

		// Keep track of current word and print it after a new line,
		// that the new branch may continue with the word stem postfixed.
		printTrieList(next)
	}
}

func printList(list []uint) {
	for _, l := range list {
		fmt.Print("[", l, "]")
	}
}

func (t *Trie) PrintList() {
	printTrieList(t.start)
}
func Verbose() {
	d |= VERBOSE
}

func printState(data Data) {

	if data.found {
		fmt.Printf("ds: trie: string '%s' found.\n", data.word)
	} else {
		fmt.Printf("ds: trie: string '%s' not found.\n", data.word)
	}
}
