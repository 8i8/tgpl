package csort

import "testing"

func TestBufferFnList(t *testing.T) {
	const fname = "TestBufferFnList"
	buf := NewSortBuffer()
	buf.Add("title")
	buf.Add("year")
	buf.Add("artist")
	buf.Add("length")
	l := buf.List()
	if l == nil {
		t.Errorf("%s: %+v", fname, l)
	}
}
