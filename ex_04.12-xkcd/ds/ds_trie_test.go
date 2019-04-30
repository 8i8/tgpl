package ds

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

var serialised string

func nodeList() node {
	n := node{}
	n.list = append(n.list, 1234)
	n.list = append(n.list, 234)
	n.list = append(n.list, 9534)
	n.list = append(n.list, 45)
	n.list = append(n.list, 859)
	return n
}

var m = make(MList)

func genList() []int {
	rand.Seed(time.Now().UnixNano())
	var list []int
	l := rand.Intn(4)
	l++
	for i := 0; i < l; i++ {
		list = append(list, rand.Intn(10000))
	}
	return list
}

var s0 = []string{
	"aardvark",
	"ardvarook",
	"hello",
	"help",
	"rdva",
	"rlock",
	"this",
	"world",
	"worlock"}

func makeTrie() Trie {

	for _, word := range s0 {
		m[word] = genList()
	}

	trie := Trie{}
	trie.BuildFromMap(m)

	return trie
}

func Test_ExpandTrie(t *testing.T) {

	trie := makeTrie()
	s1 := trie.ExpandTrie()

	if len(s1) != trie.wc {
		t.Errorf("ExpandTrie: expected %v recieved %v", trie.wc, s1)
	}
	for i, word := range s1 {
		if strings.Compare(word, s0[i]) != 0 {
			t.Errorf("ExpandTrie: %d: expected %v recieved %v", i, s0[i], word)
		}
	}
}

func Test_SearchTrie(t *testing.T) {
	trie := makeTrie()
	res := trie.SearchTrie([]string{"or"})
	if !res[0].found {
		t.Errorf("SearchTrie: `%v` wanted `%v`", res, true)
	}
}
