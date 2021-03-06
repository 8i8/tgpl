// Excercise 6.01: Impllement these additional methods:
// func (*IntSet) Len() int      // return the number of elements.
// func (*IntSet) Remove(x int)  // remove x from the set.
// func (*IntSet) Clear()        // remove all elements from the set.
// func (*IntSet) Copy() *IntSet // return a copy of the set.
package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers. Its zero value
// represents the empty set.
type IntSet struct {
	words []uint64
}

// Hes reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len returns the number of elements within the set.
func (s *IntSet) Len() (count int) {
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for i := 0; i < 64; i++ {
			if word&(1<<uint(i)) != 0 {
				count++
			}
		}
	}
	return
}

// Remove removes the gived integer from the set.
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] ^= uint64(1 << bit)
}

// Clear clears the set.
func (s *IntSet) Clear() {
	s.words = s.words[:0]
}

// Copy makes and returns a copy of the current set.
func (s *IntSet) Copy() *IntSet {
	return &IntSet{words: s.words}
}

func main() {
}
