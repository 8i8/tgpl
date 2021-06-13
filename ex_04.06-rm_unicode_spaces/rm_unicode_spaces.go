package main

import (
	"unicode"
	"unicode/utf8"
)

// Function to remove all adjacent unicode space characters in a UTF-8-encoded
// []byte slice into a single ASCII space.
func adjDuplicates(b []byte) []byte {

	count := 0
	tally := 0
	l := len(b)

	for len(b) > 0 {
		r, size := utf8.DecodeLastRune(b)
		if unicode.IsSpace(r) {
			for unicode.IsSpace(r) {
				b = b[:len(b)-size]
				count += size
				r, size = utf8.DecodeLastRune(b)
			}
			// Take a temporary measure of the currnet length.
			t := len(b)
			// Activate the entire buffer.
			b = b[:l]
			// append the buffer from the end of the unichar
			// whitespace onto the end of the next valid unispace
			// character, leaving one space in which to add an
			// ascii space character.
			b = append(b[:t+1], b[t+count:]...)
			// Add one ASCII space to replace the unichar spaces.
			b[t] = ' '
			// Set to the new shortened length.
			b = b[:t]
		}
		// Id the count has been used then characters were removed,
		// take the number of removed char from the total length, minus
		// one for the space that was added.
		if count > 0 {
			tally += count - 1
		}
		count = 0
		b = b[:len(b)-size]
	}
	// Remove the total of overwritten spaces from the overall length.
	b = b[:l-tally]
	return b
}
