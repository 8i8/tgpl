package main

import (
	"bytes"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	const fname = "TestCountingWriter"
	phrase := []byte("Hello, World!")
	buf, c := CountingWriter(&bytes.Buffer{})
	n, err := buf.Write(phrase)
	if err != nil {
		t.Errorf("%s: want nil got %q", fname, err)
	}
	n1 := len(phrase)
	if n != n1 {
		t.Errorf("%s: want %d got %d", fname, n1, n)
	}
	exp := int64(len(phrase))
	if *c != exp {
		t.Errorf("%s: want %d got %d", fname, exp, c)
	}
	n, err = buf.Write(phrase)
	if err != nil {
		t.Errorf("%s: want nil got %q", fname, err)
	}
	n1 = len(phrase)
	if n != n1 {
		t.Errorf("%s: want %d got %d", fname, n1, n)
	}
	exp = 2 * int64(len(phrase))
	if *c != exp {
		t.Errorf("%s: want %d got %d", fname, exp, c)
	}
}
