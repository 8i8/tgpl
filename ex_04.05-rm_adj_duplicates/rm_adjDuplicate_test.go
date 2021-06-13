package main

import (
	"fmt"
	"strings"
	"testing"
)

func getData() []string {
	var s []string = []string{
		"Please",
		"",
		"Get",
		"",
		"All",
		"",
		"The",
		"",
		"Spaces",
		"",
		"Outa",
		"",
		"Here",
	}
	return s
}

func getData2() []string {
	var s []string = []string{
		"Please",
		"Get",
		"All",
		"All",
		"The",
		"Doubles",
		"Doubles",
		"Doubles",
		"Outa",
		"Here",
	}
	return s
}

func TestNonempty(t *testing.T) {
	s1 := getData()
	s2 := "[Please Get All The Spaces Outa Here]"
	s1 = nonempty(s1)
	s3 := fmt.Sprintf("%v", s1)
	if strings.Compare(s2, s3) != 0 {
		t.Errorf("error: recieved %v wanted %v.", s2, s3)
	}
}

func TestNonempty2(t *testing.T) {
	s1 := getData()
	s2 := "[Please Get All The Spaces Outa Here]"
	s1 = nonempty2(s1)
	s3 := fmt.Sprintf("%v", s1)
	if strings.Compare(s2, s3) != 0 {
		t.Errorf("error: recieved %v wanted %v.", s2, s3)
	}
}

func TestAdjDuplicates(t *testing.T) {
	s1 := getData2()
	s2 := "[Please Get All The Doubles Outa Here]"
	s1 = adjDuplicates(s1)
	s3 := fmt.Sprintf("%v", s1)
	if strings.Compare(s2, s3) != 0 {
		t.Errorf("error: recieved %v wanted %v.", s2, s3)
	}
}

func BenchmarkNonempty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nonempty(getData())
	}
}

func BenchmarkNonempty2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nonempty2(getData())
	}
}

func BenchmarkAdjDoubles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		adjDuplicates(getData2())
	}
}
