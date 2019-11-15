package main

import (
	"bytes"
	"fmt"
)

// Recursive comma writing.
func commaRec(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commaRec(s[:n-3]) + "," + s[n-3:]
}

// Comma inserts commas into a non-negative integer string using a non recursive
// function applying the bytes.Buffer in place of recursion and concatenation.
func commaBf1(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var b bytes.Buffer
	for i := n - 3; i > 0; i -= 3 {
		b.Reset()
		b.WriteString(s[:i] + "," + s[i:])
		s = b.String()
	}
	return b.String()
}

// Displeased with the reseting in the above function; I decided to write
// another version, looking to increase efficiency.
func commaBf2(s string) string {

	// If less than a thousand, no commas are required, return.
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	var i int

	// Mod 3 gives us the number of digits after any thousands, the position
	// of the first comma.
	rem := n % 3
	if rem > 0 {
		buf.WriteString(s[:rem] + ",")
	}
	// Multiples of a thousand, each group of 3 zeros requires a comma.
	for i = rem; i < n-rem-3; i += 3 {
		buf.WriteString(s[i:i+3] + ",")
	}
	// Add the final remaining thousand.
	buf.WriteString(s[i:])
	return buf.String()
}

// Another version, this time splitting the iteration by dividing the length
// by 3 for each multiple of a thousand.
func commaBf3(s string) string {

	// If less than a thousand return.
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	var i, rem, div int

	// Get intermediary numbers and then multiples of a thousand.
	rem = n % 3
	div = n / 3
	// Add the intermediary to the buffer.
	if rem > 0 {
		buf.WriteString(s[:rem] + ",")
	}
	s = s[rem:]
	// Add all thousands that require a comma.
	for i = 0; i < div-1; i++ {
		j := i * 3
		buf.WriteString(s[j:j+3] + ",")
	}
	// Add the last thousand.
	buf.WriteString(s[i*3:])
	return buf.String()
}

func main() {
	fmt.Println("a: " + commaRec("10000000000"))
	fmt.Println("b: " + commaBf1("10000000000"))
	fmt.Println("c: " + commaBf2("10000000000"))
	fmt.Println("d: " + commaBf3("10000000000"))
}
