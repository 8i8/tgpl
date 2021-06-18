// Exercise 7.10: The sort.Interface type can be adapted to other uses.
// Write a function IsPalindrome(s sort.Interface) bool that reports
// whether the sequence s is a palindrome, in other words, reversing the
// sequence would not change it. Assume that the elements at indices i
// and j are equal if !s.Less(i, j) && !s.Less(j, i).
package main

import (
	"fmt"
	"sort"
)

type runes []rune

func (r runes) Len() int           { return len(r) }
func (r runes) Less(i, j int) bool { return r[i] < r[j] }
func (r runes) Swap(i, j int)      { r.swap(i, j) }

func (r runes) swap(i, j int) {
	l := len(r) >> 1
	if i < l {
		j = len(r) - 1 - i
		r[i], r[j] = r[j], r[i]
	}
}

func IsPalindrome(s sort.Interface) bool {
	l := s.Len()
	h := l >> 1 // >> 1 decrements by a power of two, giving half
	// with no remainder.
	l-- // -1 to offset being an index and not the value of a count.
	for i := 0; i < h; i++ {
		j := l - i
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

func main() {
	str := []rune("123456789987654321")
	if IsPalindrome(runes(str)) {
		fmt.Println(string(str), "is a palindrome")
	} else {
		fmt.Println(string(str), "is not a palindrome")
	}

	str = []rune("1234567891987654321")
	if IsPalindrome(runes(str)) {
		fmt.Println(string(str), "is a palindrome")
	} else {
		fmt.Println(string(str), "is not a palindrome")
	}
	str = []rune("193456789987654321")
	if IsPalindrome(runes(str)) {
		fmt.Println(string(str), "is a palindrome")
	} else {
		fmt.Println(string(str), "is not a palindrome")
	}

	str = []rune("1234767891987654321")
	if IsPalindrome(runes(str)) {
		fmt.Println(string(str), "is a palindrome")
	} else {
		fmt.Println(string(str), "is not a palindrome")
	}
}
