package main

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	const fname = "TestString"
	tr := &tree{}
	add(tr, 9)
	add(tr, 35)
	add(tr, 7)
	add(tr, 27)
	add(tr, 4)
	exp := "0 4 7 9 27 35"
	str := fmt.Sprint(tr)
	if str != exp {
		t.Errorf("%s: want %q got %q", fname, exp, str)
	}
}
