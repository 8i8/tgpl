package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestHas(t *testing.T) {
	const fname = "TestHas"
	m := &IntSet{
		words: []uint64{0b10001, 0b10001},
	}
	res := m.Has(0)
	if res != true {
		t.Errorf("%s: want true got false", fname)
	}
	res = m.Has(4)
	if res != true {
		t.Errorf("%s: want true got false", fname)
	}
	res = m.Has(64)
	if res != true {
		t.Errorf("%s: want true got false", fname)
	}
	res = m.Has(68)
	if res != true {
		t.Errorf("%s: want true got false", fname)
	}
}

func TestAdd(t *testing.T) {
	const fname = "TestAdd"
	m := &IntSet{}
	m.Add(5)
	res := m.Has(5)
	if res != true {
		t.Errorf("%s: want true got false", fname)
	}
	m.Add(89)
	res = m.Has(89)
	if res != true {
		t.Errorf("%s: want true got false", fname)
	}
}

func TestUnionWith(t *testing.T) {
	const fname = "TestUnionWith"
	m := &IntSet{
		words: []uint64{0b10001},
	}
	m2 := &IntSet{
		words: []uint64{0, 0b10001},
	}
	m.UnionWith(m2)
	res := m.Has(0)
	if res != true {
		t.Errorf("%s: want true got false", fname)
	}
	res = m.Has(4)
	if res != true {
		t.Errorf("%s: want true got false", fname)
	}
	res = m.Has(64)
	if res != true {
		t.Errorf("%s: want true got false", fname)
	}
	res = m.Has(68)
	if res != true {
		t.Errorf("%s: want true got false", fname)
	}
}

func TestString(t *testing.T) {
	const fname = "TestString"
	m := &IntSet{
		words: []uint64{0b10001, 0b10001, 0},
	}
	str := fmt.Sprintf("%s", m)
	exp := "{0 4 64 68}"
	if str != exp {
		t.Errorf("%s: want %q got %q", fname, "", str)
	}
}

func TestLen(t *testing.T) {
	const fname = "TestLen"
	m := &IntSet{
		words: []uint64{0b10001, 0b10001, 0, 0b10101010},
	}
	l := m.Len()
	exp := 8
	if l != exp {
		t.Errorf("%s: want %d got %d", fname, exp, l)
	}
}

func TestRemove(t *testing.T) {
	const fname = "TestRemove"
	m := &IntSet{
		words: []uint64{0b10001, 0b10001, 0, 0b10101010},
	}
	m.Remove(68)
	exp := "{0 4 64 193 195 197 199}"
	if strings.Compare(m.String(), exp) != 0 {
		t.Errorf("%s: want %q got %q", fname, exp, m)
	}
}

func TestClear(t *testing.T) {
	const fname = "TestClear"
	m := &IntSet{
		words: []uint64{0b10001, 0b10001, 0, 0b10101010},
	}
	m.Clear()
	exp := "{}"
	if strings.Compare(m.String(), exp) != 0 {
		t.Errorf("%s: want %q got %q", fname, exp, m)
	}
}

func TestCopy(t *testing.T) {
	const fname = "TestCopy"
	m := &IntSet{
		words: []uint64{0b10001, 0b10001, 0, 0b10101010},
	}
	m2 := m.Copy()
	str := m.String()
	str2 := m2.String()
	if strings.Compare(str, str2) != 0 {
		t.Errorf("%s: want %s got %s", fname, str, str2)
	}
	m.Add(400)
	str = m.String()
	str2 = m2.String()
	if strings.Compare(str, str2) == 0 {
		t.Errorf("%s: want != got ==\n%s\n%s", fname, str, str2)
	}
}

func TestAddAll(t *testing.T) {
	const fname = "TestAddAll"
	m := &IntSet{
		words: []uint64{0b10001, 0b10001, 0, 0b10101010},
	}
	m.AddAll(3, 5, 9, 1, 505)
	str := m.String()
	exp := "{0 1 3 4 5 9 64 68 193 195 197 199 505}"
	if strings.Compare(str, exp) != 0 {
		t.Errorf("%s: want == got !=\n%s\n%s", fname, str, exp)
	}
}

func TestIntersectWith(t *testing.T) {
	const fname = "TestIntersectWith"
	m := &IntSet{
		words: []uint64{0b10001, 0b10001, 0, 0b10101010},
	}
	m2 := &IntSet{
		words: []uint64{0, 0b10001},
	}
	m.IntersectWith(m2)
	str := m.String()
	exp := "{64 68}"
	if strings.Compare(str, exp) != 0 {
		t.Errorf("%s: want == got !=\n%s\n%s", fname, str, exp)
	}
}

func TestDifferenceWith(t *testing.T) {
	const fname = "TestDifferenceWith"
	m := &IntSet{
		words: []uint64{0b10001, 0b10101, 0, 0b10101010, 0b101},
	}
	m2 := &IntSet{
		words: []uint64{0b10001, 0b10001, 0b01010, 0b10101010},
	}
	m.DifferenceWith(m2)
	str := m.String()
	exp := "{66 256 258}"
	if strings.Compare(str, exp) != 0 {
		t.Errorf("%s: want == got !=\n%s\n%s", fname, str, exp)
	}
}

func TestSymmetricDifferenceWith(t *testing.T) {
	const fname = "TestSymmetricDifference"
	m := &IntSet{
		words: []uint64{0b10001, 0b10101, 0b00000, 0b10101010, 0b0101},
	}
	m2 := &IntSet{
		words: []uint64{0b10001, 0b10001, 0b01010, 0b10101010, 0b1010, 0b1100},
	}
	m3 := &IntSet{
		words: []uint64{0b00000, 0b00100, 0b01010, 0b00000000, 0b1111, 0b1100},
	}
	m.SymmetricDifferenceWith(m2)
	str := m.String()
	exp := "{66 129 131 256 257 258 259 322 323}"
	exp = m3.String()
	if strings.Compare(str, exp) != 0 {
		t.Errorf("%s: want == got !=\n%s\n%s", fname, str, exp)
	}
}

func TestElems(t *testing.T) {
	const fname = "TestElems"
	m := &IntSet{
		words: []uint64{0b00000, 0b00100, 0b01010, 0b00000000, 0b1111, 0b1100},
	}
	elems := m.Elems()
	exp := []int{66, 129, 131, 256, 257, 258, 259, 322, 323}
	for i := range elems {
		if elems[i] != exp[i] {
			t.Errorf("%s: want == got !=\n%v\n%v", fname, elems, exp)
			break
		}
	}
}