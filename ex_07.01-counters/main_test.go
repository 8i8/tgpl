package main

import "testing"

var phrase = []byte("The quick brown\nfox jumped over\nthe lazy dog.")

func TestByteCounter(t *testing.T) {
	const fname = "TestByteCounter"
	var b ByteCounter
	p, err := b.Write(phrase)
	if err != nil {
		t.Errorf("%s: want nil got %q", fname, err)
	}
	exp := 45
	if p != exp {
		t.Errorf("%s: want %d got %d", fname, exp, p)
	}
	exp = 45
	if b != ByteCounter(exp) {
		t.Errorf("%s: want %d got %d", fname, b, p)
	}
}

func TestWordCounter(t *testing.T) {
	const fname = "TestWordCounter"
	var w WordCounter
	p, err := w.Write(phrase)
	if err != nil {
		t.Errorf("%s: want nil got %q", fname, err)
	}
	exp := 45
	if p != exp {
		t.Errorf("%s: want %d got %d", fname, exp, p)
	}
	exp = 9
	if w != WordCounter(exp) {
		t.Errorf("%s: want %d got %d", fname, w, p)
	}
}

func TestLineCounter(t *testing.T) {
	const fname = "TestLineCounter"
	var l LineCounter
	p, err := l.Write(phrase)
	if err != nil {
		t.Errorf("%s: want nil got %q", fname, err)
	}
	exp := 45
	if p != exp {
		t.Errorf("%s: want %d got %d", fname, exp, p)
	}
	exp = 3
	if l != LineCounter(exp) {
		t.Errorf("%s: want %d got %d", fname, l, p)
	}
}
