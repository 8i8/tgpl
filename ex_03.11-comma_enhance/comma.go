package main

import (
	"bytes"
	"fmt"
	"strings"
)

func commaRec(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commaRec(s[:n-3]) + "," + s[n-3:]
}

// Displeased with the reseting in the above function; I decided to write
// another version, looking to invrease efficincy.
func commaBuf(s string) string {

	var buf bytes.Buffer
	var sign, tail, rem, i int

	// Check if there is a sign at the start of the string if so copy it
	// into the buffer and keep track of the space.
	if s[0] == '-' || s[0] == '+' {
		sign++
	}

	// Check for decimal point. If present store the number length before
	// the point in n and after the pont in tail.
	n := strings.LastIndex(s, ".")
	if n > 0 {
		tail = len(s[n:])
	} else {
		n = len(s)
	}

	// Is the number below 1000?
	if n <= 3+sign {
		return s
	}

	// Write the sign if it exists, to the buffer.
	if sign > 0 {
		buf.WriteString(s[:sign])
	}

	// Check for numbers that are between the thousands denomination and
	// add any that are present to the buffer followed by a comma.
	rem = (n - sign) % 3
	if rem > 0 {
		buf.WriteString(s[sign:rem+sign] + ",")
	}

	// For every factor of one thousand, copy into the buffer and add a
	// comma.
	for i = rem + sign; i < n-3; i += 3 {
		buf.WriteString(s[i:i+3] + ",")
	}

	// Add the last thousand that does not require a comma, and the decimal
	// point if it is there.
	buf.WriteString(s[i:n])

	// If there is a decimal tail, add the trailing decimal digits.
	if tail > 0 {
		buf.WriteString(s[n:])
	}

	return buf.String()
}

func main() {
	fmt.Println("a: " + commaRec("-1000"))
	fmt.Println("b: " + commaBuf("10000.0"))
}
