package main

import (
	"fmt"
	"strings"
	"testing"
)

var a []int
var d [LEN]int = [LEN]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
var s2 string = "[12 11 10 9 8 7 6 5 4 3 2 1 0]"

func init() {
	a = append(a, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
}

func TestReverseSlice(t *testing.T) {
	c := a
	c = reverseSlice(c)
	s1 := fmt.Sprintf("%v", c)
	if strings.Compare(s1, s2) != 0 {
		t.Errorf("error: got %v expected %v", s1, s2)
	}
}

func TestReverseArray(t *testing.T) {
	e := d
	reverseArray(&e)
	s1 := fmt.Sprintf("%v", e)
	if strings.Compare(s1, s2) != 0 {
		t.Errorf("error: got %v expected %v", s1, s2)
	}
}

func TestReverseArrayUnwound(t *testing.T) {
	e := d
	reverseArrayUnwound(&e)
	s1 := fmt.Sprintf("%v", e)
	if strings.Compare(s1, s2) != 0 {
		t.Errorf("error: got %v expected %v", s1, s2)
	}
}

func TestReverseArrayCopyUnwound(t *testing.T) {
	e := d
	e = reverseArrayCopyUnwound(e)
	s1 := fmt.Sprintf("%v", e)
	if strings.Compare(s1, s2) != 0 {
		t.Errorf("error: got %v expected %v", s1, s2)
	}
}

func BenchmarkReverseSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseSlice(a[:])
	}
}

func BenchmarkReverseArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseArray(&d)
	}
}

func BenchmarkReverseArrayUnwound(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseArrayUnwound(&d)
	}
}

func BenchmarkReverseArrayCopyUnwound(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseArrayCopyUnwound(d)
	}
}
