package ds

import (
	"bytes"
	"io"
	"math/rand"
	"strings"
	"testing"
	"time"

	"tgpl/ex_04.12-xkcd/msg"
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

func Test_writeList(t *testing.T) {
	var buf bytes.Buffer
	w := msg.NewErrWriter(&buf)
	n := nodeList()

	writeList(&n, w)

	s1 := buf.String()
	s2 := " 1234,234,9534,45,859, "
	if strings.Compare(s1, s2) != 0 {
		t.Errorf("result: Test_writeList: `%v` wanted `%v`", s1, s2)
	}
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

func simplTrie() Trie {
	trie := Trie{}
	trie.Add("hello", []int{9, 8, 40, 32})
	trie.Add("help", []int{5, 4, 6, 396})
	trie.Add("world", []int{95, 94, 8, 783})
	trie.Add("worlock", []int{3, 9534, 385, 8343})
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
	res := trie.SearchTrie([]string{"elp"})
	if !res[0].found {
		t.Errorf("SearchTrie: `%v` wanted `%v`", res, true)
	}
}

func Test_GetIds(t *testing.T) {
	buf := bytes.NewBuffer([]byte("9,8,40,32 "))
	re := msg.NewErrReader(buf)
	n := new(node)

	_, err := getIds(n, re)
	if err != nil && err != io.EOF {
		t.Errorf("GetIds: `%v`", err)
	}

	l1 := n.list
	l2 := []int{9, 8, 40, 32}

	if len(l1) != len(l2) {
		t.Errorf("GetIds: `%v` wanted `%v`", l1, l2)
	} else {
		for i, v := range l2 {
			if l1[i] != v {
				t.Errorf("GetIds: %d:`%d` wanted `%d`", i, l1[i], v)
			}
		}
	}
}

func Test_SerialiseDeserialiseTrie(t *testing.T) {
	var buf bytes.Buffer
	w := msg.NewErrWriter(&buf)
	trie := makeTrie()
	s1 := trie.ExpandTrie()

	err := trie.SerialiseTrie(w)
	if err != nil {
		t.Errorf("SerialiseTrie: error: %v", err)
	}
	serialised = buf.String()

	buf2 := bytes.NewBuffer([]byte(serialised))
	re := msg.NewErrReader(buf2)

	trie2 := Trie{}
	err = trie2.DeserialiseTrie(re)
	if err != nil && err != io.EOF {
		t.Errorf("DeserialiseTrie: error: %v", err)
	}
	s2 := trie2.ExpandTrie()

	if len(s1) != len(s2) {
		t.Errorf("SerialiseDeserialiseTrie: recieved `%v` wanted `%v`", s2, s1)
	}

	for i, word := range s2 {
		if strings.Compare(word, s1[i]) != 0 {
			t.Errorf("SerialiseDeserialiseTrie: recieved `%v` wanted `%v`", word, s1[i])
		}
	}

	s3 := []string{"hello", "world"}
	res := trie.SearchTrie(s3)
	if len(res) != 2 {
		t.Errorf("DeserialiseTrie: SearchTrie: `%v` wanted `%v`", len(res), 4)
	}
}

// func Test_DeserialiseFile(t *testing.T) {

// 	// Open cache file for reading.
// 	file, err := os.Open("../data/cache.test.data")
// 	if err != nil {
// 		t.Errorf("DeserialiseTrie: file open error: %v", err)
// 	}

// 	// Perform deserialisation.
// 	rw := msg.NewErrReader(file)
// 	trie := Trie{}
// 	err = trie.DeserialiseTrie(rw)
// 	if err != nil && err != io.EOF {
// 		t.Errorf("DeserialiseTrie: error: %v", err)
// 	}

// 	s1 := []string{"hello", "world"}
// 	res := trie.SearchTrie(s1)

// 	for i, _ := range res {
// 		if !res[i].found {
// 			t.Errorf("DeserialiseFile: `%v` wanted `%v`", res[i], true)
// 		}
// 	}

// 	// Finish up.
// 	err = file.Close()
// 	if err != nil {
// 		t.Errorf("DeserialiseTrie: file close error: %v", err)
// 	}
// }
