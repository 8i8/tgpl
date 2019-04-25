package msg

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

type node struct {
	list []int
}

func makeList() node {
	var n node
	n.list = append(n.list, 1)
	n.list = append(n.list, 2)
	n.list = append(n.list, 3)
	n.list = append(n.list, 4)
	n.list = append(n.list, 5)
	return n
}

func writeList(n *node, w *ErrWriter) error {

	w.WriteRune(' ')
	for i, id := range n.list {
		w.WriteInt(id)
		// Add a comma between all integers, except for the last.
		if i != len(n.list)-1 {
			w.WriteRune(',')
		}
	}
	w.WriteRune(' ')
	if w.Err != nil {
		return fmt.Errorf("WriteRune: %v", w.Err)
	}
	return nil
}

func TestErrWriter(t *testing.T) {
	var buf bytes.Buffer
	w := NewErrWriter(&buf)
	n := makeList()

	writeList(&n, w)

	s1 := buf.String()
	s2 := " 1,2,3,4,5 "
	if strings.Compare(s1, s2) != 0 {
		t.Errorf("test: writeList: recieved: %v expected: %v", s1, s2)
	}
}
