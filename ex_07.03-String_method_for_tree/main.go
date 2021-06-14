package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

func write(b *bytes.Buffer, t *tree) {
	if t.left != nil {
		write(b, t.left)
	}
	b.WriteString(" " + strconv.Itoa(t.value))
	if t.right != nil {
		write(b, t.right)
	}
}

func (t *tree) String() string {
	if t == nil {
		return ""
	}
	buf := &bytes.Buffer{}
	write(buf, t)
	b := buf.Bytes()
	return string(b[1:])
}

// Sort sorts values in place, by using a binary tree and insertion
// sort.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order and returns
// the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

// add adds a value to the binary tree in the appropriate leaf node.
func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	a := []int{75, 24, 85, 34, 5, 734, 85, 323, 6, 8, 34, 8543, 56, 241}
	Sort(a)
	fmt.Println(a)
}
