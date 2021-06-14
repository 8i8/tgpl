// Exercise 07.02: Write a function CountingWriter with the signature
// below that, given an io.Writer, returns a Writer that wraps the
// original, and a pointer to an int64 variable that at any moment
// contains the number of bytes written to the new Writer.
//	func CountingWriter(w io.Writer) (io.Writer, *int64)
package main

import "io"

type ByteCounter struct {
	c *int64
	w io.Writer
}

func (b *ByteCounter) Write(p []byte) (int, error) {
	*(b.c) += int64(len(p))
	return b.w.Write(p)
}

// CountingWriter wraps the given writer with another writer that keeps
// track of how many bytes in total have been written.
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var c int64
	b := &ByteCounter{w: w, c: &c}
	return b, &c
}

func main() {
}
