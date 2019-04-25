package ds

import (
	"bytes"
	"strings"
	"testing"

	"tgpl/ex_04.12-xkcd/msg"
)

func nodeList() node {
	n := node{}
	n.list = append(n.list, 1234)
	n.list = append(n.list, 234)
	n.list = append(n.list, 9534)
	n.list = append(n.list, 45)
	n.list = append(n.list, 859)
	return n
}

func Test_writeList(t *testing.T) {
	var buf bytes.Buffer
	w := msg.NewErrWriter(&buf)
	n := nodeList()

	writeList(&n, w)

	s1 := buf.String()
	s2 := " 1234,234,9534,45,859 "
	if strings.Compare(s1, s2) != 0 {
		t.Errorf("result: Test_writeList: `%v` wanted `%v`", s1, s2)
	}
}

func makeTrie() Trie {
	trie := Trie{}
	trie.Add("hello", []int{9, 8, 40, 32})
	trie.Add("world", []int{95, 94, 8, 783})
	trie.Add("help", []int{5, 4, 6, 396})
	trie.Add("worlock", []int{3, 9534, 385, 8343})
	return trie
}

func Test_SearchTrie(t *testing.T) {
	trie := makeTrie()
	res := trie.SearchTrie([]string{"elp"})
	if len(res[0].List) != 4 {
		for _, l := range res {
			t.Errorf("list: %v", l)
		}
		t.Errorf("SearchTrie: `%v` wanted `%v`", len(res), 4)
	}
}

func Test_SerialiseTrie(t *testing.T) {
	var buf bytes.Buffer
	w := msg.NewErrWriter(&buf)

	trie := makeTrie()
	err := trie.SerialiseTrie(w)
	s1 := buf.String()
	s2 := "Hello"

	if err != nil {
		t.Errorf("SerialiseTrie: error: %v", err)
	}

	if strings.Compare(s1, s2) != 0 {
		t.Errorf("result: Test_writeList: `%v` wanted `%v`", s1, s2)
	}
}
