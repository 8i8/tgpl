package trie

type hash map[rune]*Trie

type Trie struct {
	r    rune
	stop bool
	next hash
}

// Add inserts the given string into the trie.
func (t *Trie) Add(s string) {

	for _, r := range s {
		if t.next == nil {
			t.next = make(hash)
		}
		if t.next[r] == nil {
			t.next[r] = new(Trie)
			t.next[r].r = r
		}
		t = t.next[r]
	}
	t.stop = true
}

// Check verifies if the given string exists in the trie as a complete word.
func (t *Trie) Check(s string) bool {

	for _, r := range s {
		if t.next[r] == nil {
			break
		}
		t = t.next[r]
	}
	return true
}

// SubWord checks if the given string exists as a sub word with the trie.
func (t *Trie) SubWord(s string) bool {

	var out bool
	for _, h := range t.next {
		out = h.Check(s)
	}
	for _, h := range t.next {
		u := h.next
		r := h.r
		out = u[r].SubWord(s)
	}
	return out
}
