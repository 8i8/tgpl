package main

import (
	"bytes"
	"testing"
)

func TestReverseBytes(t *testing.T) {
	var s1 []byte = []byte("desrever eb dluohs sihT")
	var s2 []byte = []byte("This should be reversed")

	s3 := ReverseBytes(s1)
	if bytes.Compare(s3, s2) != 0 {
		t.Errorf("error: received `%v` wanted `%v`.\n", string(s3), string(s2))
	}
}

func TestReverseUtf8(t *testing.T) {
	var s1 []byte = []byte("desrever eb dluohs sihT")
	var s2 []byte = []byte("This should be reversed")

	s3 := ReverseUtf8(s1)
	if bytes.Compare(s3, s2) != 0 {
		t.Errorf("error: received `%v` wanted `%v`.\n", string(s3), string(s2))
	}

	var s4 []byte = []byte("界世 ,olleH")
	var s5 []byte = []byte("Hello, 世界")

	s6 := ReverseUtf8(s4)
	if bytes.Compare(s6, s5) != 0 {
		t.Errorf("error: received `%v` wanted `%v`.\n", string(s6), string(s5))
	}

	var s7 []byte = []byte("Hello, 世界")
	var s8 []byte = []byte("界世 ,olleH")

	s9 := ReverseUtf8(s7)
	if bytes.Compare(s9, s8) != 0 {
		t.Errorf("error: received `%v` wanted `%v`.\n", string(s9), string(s8))
	}
}
