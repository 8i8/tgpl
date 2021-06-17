package csort

import (
	"sort"
	"testing"
)

func TestCustomSort(t *testing.T) {
	const fname = "TestCustomSort"
	data := []int{1, 43, 234, 3, 14, 65, 89, 879}
	exp := []int{1, 3, 14, 43, 65, 89, 234, 879}

	_ = NewSortBuffer(empty)
	inter := wrapInter(data)
	sort.Sort(New(testSortFn(), inter))
	unwrapInter(inter, data)

	for i := range data {
		if data[i] != exp[i] {
			t.Errorf("%s:\nwant %v\ngot  %v", fname, exp, data)
		}
	}
}

// wrapInter returns the given slice as a slice of interfaces.
func wrapInter(data []int) []interface{} {
	inter := make([]interface{}, len(data))
	for i := range data {
		inter[i] = data[i]
	}
	return inter
}

func unwrapInter(inter []interface{}, data []int) {
	for i := range inter {
		data[i] = inter[i].(int)
	}
}

func testSortFn() func(xi, yi interface{}) bool {
	return func(xi, yi interface{}) bool {
		x := xi.(int)
		y := yi.(int)
		if testLess(x, y) {
			return true
		}
		return false
	}
}

func empty(s string) SortFn {
	return func(x, y interface{}) int {
		return 0
	}
}

// Less returns a sort function tha examins the given struct element.
func testLess(x, y int) bool {
	if x == y {
		return false
	}
	if x < y {
		return true
	}
	return false
}
