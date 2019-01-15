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

	s1 = []byte("界世 ,olleH")
	s2 = []byte("Hello, 世界")

	s3 = ReverseUtf8(s1)
	if bytes.Compare(s3, s2) != 0 {
		t.Errorf("error: received `%v` wanted `%v`.\n", string(s3), string(s2))
	}

	s1 = []byte("Hello, 世界")
	s2 = []byte("界世 ,olleH")

	s3 = ReverseUtf8(s1)
	if bytes.Compare(s3, s2) != 0 {
		t.Errorf("error: received `%v` wanted `%v`.\n", string(s3), string(s2))
	}

	s1 = []byte("Hello,世 界 ")
	s2 = []byte(" 界 世,olleH")

	s3 = ReverseUtf8(s1)
	if bytes.Compare(s3, s2) != 0 {
		t.Errorf("error: received `%v` wanted `%v`.\n", string(s3), string(s2))
	}
}
