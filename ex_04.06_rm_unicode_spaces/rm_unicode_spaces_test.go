package main

import (
	"bytes"
	"testing"
)

var s2 []byte = []byte(" This is a unicode whitespace free sentance. ")

func TestAdjDuplicates(t *testing.T) {

	var s1 []byte = []byte("  This is a     unicode  whitespace  free sentance.  ")
	s := adjDuplicates(s1)
	if bytes.Compare(s, s2) != 0 {
		t.Errorf("error: recieved `%v` expected `%v`", string(s), string(s2))
	}
}
