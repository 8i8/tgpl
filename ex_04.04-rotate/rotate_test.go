package main

import (
	"fmt"
	"strings"
	"testing"
)

var s []int = []int{0, 1, 2, 3, 4, 5}
var v [LEN]int = [LEN]int{0, 1, 2, 3, 4, 5}

func TestRotate(t *testing.T) {

	s1 := rotate(s[:])
	str := fmt.Sprintf("%v", s1)
	s2 := "[2 3 4 5 0 1]"
	if strings.Compare(str, s2) != 0 {
		t.Errorf("error: recieved %v expected %v.", str, s2)
	}
}

func TestRotateSinglePass(t *testing.T) {

	s1 := rotateSinglePass(v)
	str := fmt.Sprintf("%v", s1)
	s2 := "[2 3 4 5 0 1]"
	if strings.Compare(str, s2) != 0 {
		t.Errorf("error: recieved %v expected %v.", str, s2)
	}
}

func BenchmarkRotate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotate(s)
	}
}

func BenchmarkRotateSinglePass(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotateSinglePass(v)
	}
}
