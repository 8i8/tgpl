package data

import (
	"sortInterface/data/csort"
	"testing"
)

func TestSortBuffer(t *testing.T) {
	const fname = "TestSortBuffer"
	buf := csort.NewSortBuffer(Less)
	buf.Load("title", "artist", "length")
	str := buf.String()
	exp := "title,artist,length"
	if str != exp {
		t.Errorf("%s:\nwant %s\ngot  %s", fname, exp, str)
	}
}

func TestSortBufferReverse(t *testing.T) {
	const fname = "TestSortBufferReverse"
	buf := csort.NewSortBuffer(Less)
	buf.Load("title", "artist", "length", "length")
	buf.Add("length")
	str := buf.String()
	exp := "title,artist,length,length-rev"
	if str != exp {
		t.Errorf("%s:\nwant %s\ngot  %s", fname, exp, str)
	}
}
