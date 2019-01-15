package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func ReverseBytes(b []byte) []byte {
	for l, r := 0, len(b)-1; l < r; l, r = l+1, r-1 {
		b[l], b[r] = b[r], b[l]
	}
	return b
}

func ReverseUtf8(b []byte) []byte {

	out := b
	for cl := len(b); cl > 1; cl = len(b) {

		// Get the length of the left and rightmost runes.
		lr, ls := utf8.DecodeRune(b)
		rr, rs := utf8.DecodeLastRune(b)
		if ls == rs {
			// Direct swap.
			utf8.EncodeRune(b[:rs], rr)
			utf8.EncodeRune(b[cl-ls:], lr)
		} else {
			// Shunt the array when runes are differing sizes.
			b = append(b[:rs], b[ls:cl+ls-rs]...)
			utf8.EncodeRune(b[:rs], rr)
			utf8.EncodeRune(b[cl-ls:cl], lr)
		}
		b = b[rs : cl-ls]
	}
	return out
}

func main() {

	var s1 []byte = []byte("desrever eb dluohs sihT")
	var s2 []byte = []byte("This should be reversed")

	s3 := ReverseUtf8(s1)
	if bytes.Compare(s3, s2) != 0 {
		fmt.Printf("error: received `%v` wanted `%v`.\n", string(s3), string(s2))
	} else {
		fmt.Println("control: ", string(s2))
		fmt.Println("result: ", string(s3))
	}

	var s4 []byte = []byte("界世 ,olleH")
	var s5 []byte = []byte("Hello, 世界")

	s6 := ReverseUtf8(s4)
	if bytes.Compare(s6, s5) != 0 {
		fmt.Printf("error: recieved `%v` wanted `%v`.", string(s6), string(s5))
	} else {
		fmt.Println("control: ", string(s5))
		fmt.Println("result: ", string(s6))
	}

	var s7 []byte = []byte("Hello, 世界")
	var s8 []byte = []byte("界世 ,olleH")

	s9 := ReverseUtf8(s7)
	if bytes.Compare(s9, s8) != 0 {
		fmt.Printf("error: received `%v` wanted `%v`.\n", string(s9), string(s8))
	} else {
		fmt.Println("control: ", string(s8))
		fmt.Println("result: ", string(s9))
	}
}
