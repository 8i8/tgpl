// Exercise 6.05: The type of each word used by IntSet is uint64, but 64
// bit arithmatic may be inefficient on a 32 bit platform. Modify the
// program to use the uint type, which is the most efficien unsigned
// integer type for the platform. Instead of dividing by 64 define a
// constant holding the effective size of uint in bits, 32 or 64. You
// can use the perhaps too-clever expression 32 << (^uint(0) >> 63) for
// this purpose.
package main

import (
	"bytes"
	"fmt"
)

// usize is either 32 or 64 depending upon the system, 32 or 64 bit.
const usize = 32 << (^uint(0) >> 63)

// An IntSet is a set of small non-negative integers. Its zero value
// represents the empty set.
type IntSet struct {
	words []uint
}

// Hes reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/usize, uint(x%usize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/usize, uint(x%usize)
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
		for j := 0; j < usize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", usize*i+j)
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
		for i := 0; i < usize; i++ {
			if word&(1<<uint(i)) != 0 {
				count++
			}
		}
	}
	return
}

// Remove removes the given integer from the set.
func (s *IntSet) Remove(x int) {
	word, bit := x/usize, uint(x%usize)
	s.words[word] ^= uint(1 << bit)
}

// Clear clears the set.
func (s *IntSet) Clear() {
	s.words = s.words[:0]
}

// Copy makes and returns a copy of the current set.
func (s *IntSet) Copy() *IntSet {
	return &IntSet{words: s.words}
}

// AddAll add all of the posative intagers given to the set.
func (s *IntSet) AddAll(elems ...int) {
	for _, x := range elems {
		s.Add(x)
	}
}

// IntersectWith sets s to the intersection between s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			break
		}
	}
	s.words = s.words[:len(t.words)]
}

// DifferenceWith sets s to the difference between s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, word := range s.words {
		if i < len(t.words) {
			s.words[i] = word &^ t.words[i]
		}
	}
}

// SymmetricDifference sets s to the symmetric difference between s and
// t.
func (s *IntSet) SymmetricDifferenceWith(t *IntSet) {
	if len(t.words) > len(s.words) {
		s.words, t.words = t.words, s.words
	}
	for i, tword := range t.words {
		s.words[i] ^= tword
	}
}

// Elems returns an a slice containing the content of the set in
// increasing order of magnitude.
func (s *IntSet) Elems() (set []int) {
	l := s.Len()
	if cap(set) < l {
		set = append(set, make([]int, 0, l)...)
	}
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < usize; j++ {
			if word&(1<<uint(j)) != 0 {
				set = append(set, usize*i+j)
			}
		}
	}
	return
}

func main() {
}
