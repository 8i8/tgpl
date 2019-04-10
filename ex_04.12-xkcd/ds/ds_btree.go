package ds

import "fmt"

// BtreeNode is the primary construct of a binary tree.
type BtreeNode struct {
	index uint
	prev  *BtreeNode
	left  *BtreeNode
	right *BtreeNode
}

type Count map[uint]int

// BtreeAdd adds a new node to a binary tree.
func BtreeAdd(node *BtreeNode, data BtreeNode) {

	// End node; Add new node and return.
	if node == nil {
		node = &data
		return
	}

	// Equal.
	if data.index == node.index {
		return
	}

	// Less than.
	if data.index < node.index {
		next := node.left
		next.prev = node
		BtreeAdd(next, data)
		return
	}

	// Greater than.
	if data.index > node.index {
		next := node.right
		next.prev = node
		BtreeAdd(next, data)
		return
	}
}

// BtreeCount generates a map of all comic indexes counting how many times they
// are added, used to check if the word exists in multiple btrees.
func BtreeCount(node *BtreeNode, m Count) Count {

	if node == nil {
		return m
	}

	m[node.index]++

	fmt.Println(node.index)

	next := node.left
	BtreeCount(next, m)

	next = node.right
	BtreeCount(next, m)

	return m
}
