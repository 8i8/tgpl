package main

import (
	"io"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	const fname = "TestLimitReader"
	str := "This is a test string for testing the LimitReader"
	l := int64(14)
	r := LimitReader(strings.NewReader(str), l)
	buf := make([]byte, l+1)
	n, err := r.Read(buf)
	if err != io.EOF {
		t.Errorf("%s: want EOF got %s", fname, err)
	}
	if int64(n) != l {
		t.Errorf("%s: want %d got %d", fname, l, n)
	}
	exp := str[:l]
	str = string(buf[:l])
	if strings.Compare(str, exp) != 0 {
		t.Errorf("%s: want %s got %s", fname, exp, str)
	}
}
