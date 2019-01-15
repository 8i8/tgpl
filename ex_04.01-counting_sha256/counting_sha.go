package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

var pc [256]byte

const (
	SHA256 = 32
)

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func BitCount(x uint64) int {
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}

// Looping bitwise operations and the uint64 is bing drawn out of the byte
// slice using binary.LittleEndian.Uint64.
func BitComp1(c1, c2 [SHA256]byte) int {
	var n uint64

	for i := 0; i < 4; i++ {
		j := i * 8
		x := binary.LittleEndian.Uint64(c1[j : j+8])
		y := binary.LittleEndian.Uint64(c2[j : j+8])
		x = x ^ y
		x = x - ((x >> 1) & 0x5555555555555555)
		x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
		x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
		x = x + (x >> 8)
		x = x + (x >> 16)
		x = x + (x >> 32)
		n += x & 0x7f
	}

	return int(n)
}

// Fast bitwise operations and the uint64 is bing drawn out of the byte slice
// using binary.LittleEndian.Uint64.
func BitComp2(c1, c2 [SHA256]byte) int {

	var n uint64
	x := binary.LittleEndian.Uint64(c1[0:8])
	y := binary.LittleEndian.Uint64(c2[0:8])

	x = x ^ y
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	n += x & 0x7f

	x = binary.LittleEndian.Uint64(c1[8:16])
	y = binary.LittleEndian.Uint64(c2[8:16])

	x = x ^ y
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	n += x & 0x7f

	x = binary.LittleEndian.Uint64(c1[16:24])
	y = binary.LittleEndian.Uint64(c2[16:24])

	x = x ^ y
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	n += x & 0x7f

	x = binary.LittleEndian.Uint64(c1[24:32])
	y = binary.LittleEndian.Uint64(c2[24:32])

	x = x ^ y
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	n += x & 0x7f

	return int(n)
}

// To test if it is any faster calling only one line for all the counting.
func BitComp3(c1, c2 [SHA256]byte) int {

	x := binary.LittleEndian.Uint64(c1[0:8])
	y := binary.LittleEndian.Uint64(c2[0:8])
	x2 := binary.LittleEndian.Uint64(c1[8:16])
	y2 := binary.LittleEndian.Uint64(c2[8:16])
	x3 := binary.LittleEndian.Uint64(c1[16:24])
	y3 := binary.LittleEndian.Uint64(c2[16:24])
	x4 := binary.LittleEndian.Uint64(c1[24:32])
	y4 := binary.LittleEndian.Uint64(c2[24:32])

	return int(pc[byte(x>>(0*32))] ^ pc[byte(y>>(0*32))] +
		pc[byte(x>>(1*32))] ^ pc[byte(y>>(1*32))] +
		pc[byte(x>>(2*32))] ^ pc[byte(y>>(2*32))] +
		pc[byte(x>>(3*32))] ^ pc[byte(y>>(3*32))] +
		pc[byte(x>>(4*32))] ^ pc[byte(y>>(4*32))] +
		pc[byte(x>>(5*32))] ^ pc[byte(y>>(5*32))] +
		pc[byte(x>>(6*32))] ^ pc[byte(y>>(6*32))] +
		pc[byte(x>>(7*32))] ^ pc[byte(y>>(7*32))] +
		pc[byte(x2>>(8*32))] ^ pc[byte(y2>>(8*32))] +
		pc[byte(x2>>(9*32))] ^ pc[byte(y2>>(9*32))] +
		pc[byte(x2>>(10*32))] ^ pc[byte(y2>>(10*32))] +
		pc[byte(x2>>(11*32))] ^ pc[byte(y2>>(11*32))] +
		pc[byte(x2>>(12*32))] ^ pc[byte(y2>>(12*32))] +
		pc[byte(x2>>(13*32))] ^ pc[byte(y2>>(13*32))] +
		pc[byte(x2>>(14*32))] ^ pc[byte(y2>>(14*32))] +
		pc[byte(x2>>(15*32))] ^ pc[byte(y2>>(15*32))] +
		pc[byte(x3>>(16*32))] ^ pc[byte(y3>>(16*32))] +
		pc[byte(x3>>(17*32))] ^ pc[byte(y3>>(17*32))] +
		pc[byte(x3>>(18*32))] ^ pc[byte(y3>>(18*32))] +
		pc[byte(x3>>(19*32))] ^ pc[byte(y3>>(19*32))] +
		pc[byte(x3>>(20*32))] ^ pc[byte(y3>>(20*32))] +
		pc[byte(x3>>(21*32))] ^ pc[byte(y3>>(21*32))] +
		pc[byte(x3>>(22*32))] ^ pc[byte(y3>>(22*32))] +
		pc[byte(x3>>(23*32))] ^ pc[byte(y3>>(23*32))] +
		pc[byte(x4>>(24*32))] ^ pc[byte(y4>>(24*32))] +
		pc[byte(x4>>(25*32))] ^ pc[byte(y4>>(25*32))] +
		pc[byte(x4>>(26*32))] ^ pc[byte(y4>>(26*32))] +
		pc[byte(x4>>(27*32))] ^ pc[byte(y4>>(27*32))] +
		pc[byte(x4>>(28*32))] ^ pc[byte(y4>>(28*32))] +
		pc[byte(x4>>(29*32))] ^ pc[byte(y4>>(29*32))] +
		pc[byte(x4>>(30*32))] ^ pc[byte(y4>>(30*32))] +
		pc[byte(x4>>(31*32))] ^ pc[byte(y4>>(31*32))])
}

// A test to see if it is any faster to set all the variables before counting
// all the set bits.
func BitComp4(c1, c2 [SHA256]byte) int {
	var n uint64
	x := binary.LittleEndian.Uint64(c1[0:8])
	y := binary.LittleEndian.Uint64(c2[0:8])
	x2 := binary.LittleEndian.Uint64(c1[8:16])
	y2 := binary.LittleEndian.Uint64(c2[8:16])
	x3 := binary.LittleEndian.Uint64(c1[16:24])
	y3 := binary.LittleEndian.Uint64(c2[16:24])
	x4 := binary.LittleEndian.Uint64(c1[24:32])
	y4 := binary.LittleEndian.Uint64(c2[24:32])

	x = x ^ y
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	n += x & 0x7f

	x = x2 ^ y2
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	n += x & 0x7f

	x = x3 ^ y3
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	n += x & 0x7f

	x = x4 ^ y4
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	n += x & 0x7f

	return int(n)
}

type data struct {
	A uint64
	B uint64
	C uint64
	D uint64
}

// This code uses the encode.Read function in the calling function to put all
// the uint64 into a struct which is pased into the function. It is far to slow
// when called inside the function.
func BitComp5(d1, d2 data) int {
	var n uint64

	x, x2, x3, x4 := d1.A, d1.B, d1.C, d1.D
	y, y2, y3, y4 := d2.A, d2.B, d2.C, d2.D

	x = x ^ y
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	n += x & 0x7f

	x = x2 ^ y2
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	n += x & 0x7f

	x = x3 ^ y3
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	n += x & 0x7f

	x = x4 ^ y4
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	n += x & 0x7f

	return int(n)
}

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  Funcions found on stack overflow.
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// An exaple of a looped version of the code taken from stack overflow.
func BitsDifference(h1, h2 *[SHA256]byte) int {
	n := 0
	for i := range h1 {
		for b := h1[i] ^ h2[i]; b != 0; b &= b - 1 {
			n++
		}
	}
	return n
}

// bitCount counts the number of bits set in x
func bitCount(x uint8) int {
	count := 0
	for x != 0 {
		x &= x - 1
		count++
	}
	return count
}

// An exaple of a looped version of the code taken from stack overflow.
func DifferentBits(c1, c2 [SHA256]uint8) int {
	var counter int
	for x := range c1 {
		counter += bitCount(c1[x] ^ c2[x])
	}
	return counter
}

func main() {
	var n int
	c1 := sha256.Sum256([]byte("This"))
	c2 := sha256.Sum256([]byte("That"))

	n = BitComp2(c1, c2)
	fmt.Printf("there are %d differance between the sha's\n", n)
	n = BitComp4(c1, c2)
	fmt.Printf("there are %d differance between the sha's\n", n)
}
