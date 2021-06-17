package data

import (
	"sortInterface/data/csort"
	"testing"
)

func TestSortBuffer(t *testing.T) {
	const fname = "TestSortBuffer"
	buf := csort.NewSortBuffer(Less)
	buf.Add("title", "artist", "length")
	str := buf.String()
	exp := "title,artist,length"
	if str != exp {
		t.Errorf("%s:\nwant %s\ngot  %s", fname, exp, str)
	}
}
