// Exercise 5.16 Write a variadic version of strings.Join
package main

import "strings"

// Join concatenates the elements of its second argument separated by
// the first argument into one string, the final concatenated string is
// then returned.
func join(sep string, elems ...string) string {

	// Extraneous cases.
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return elems[0]
	}

	// Generate the appropriate sized buffer.
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(elems[i])
	}
	out := strings.Builder{}
	out.Grow(n)

	// Concatenate.
	out.WriteString(elems[0])
	for _, elem := range elems[1:] {
		out.WriteString(sep)
		out.WriteString(elem)
	}
	return out.String()
}

func main() {
}
