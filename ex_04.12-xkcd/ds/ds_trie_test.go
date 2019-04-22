package ds

import "tgpl/ex_04.12-xkcd/ds"

func Test() {

	trie := ds.Trie{}

	trie.Add("test", 9)
	trie.Add("ee", 20)
	trie.Add("all", 9)
	trie.Add("allestinet", 15)
	trie.Add("allestinet", 1)
	trie.Add("allet", 12)
	trie.Add("allet", 14)
	trie.Add("alletne", 13)
	trie.Add("alloet", 17)
	trie.Add("alloetium", 18)
	trie.Add("hello", 4)
	trie.Add("hellonetworld", 2)
	trie.Add("helloworld", 1)
	trie.Add("helloworld", 11)
	trie.Add("net", 4)
	trie.Add("net", 16)
	trie.Add("world", 9)
	trie.Add("world", 5)
	trie.Add("world", 6)
	trie.Add("world", 7)
	trie.Add("world", 8)
	trie.Add("yenhellon", 3)

	var words []string
	// words = append(words, "test")
	// words = append(words, "e")
	// words = append(words, "ello")
	// words = append(words, "net")
	// words = append(words, "nothere")
	// words = append(words, "allestinet")
	words = append(words, "hello")
	words = append(words, "world")

	trie.SearchWords(words)
	//trie.Print()
}
