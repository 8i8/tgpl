package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	words := make(map[string]int)
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	for in.Scan() {
		// Scan input splitting between words.
		b := []byte(in.Text())
		// Remove all extraneous characters, replacing punctuation with
		// spaces.
		str := string(removeNonAlphaNumeric(b))
		// Split again to catch words that were punctuated.
		s := strings.Split(str, " ")
		for _, str := range s {
			if str != "" {
				words[strings.ToLower(strings.TrimSpace(str))]++
			}
		}
	}
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	// Make array of words and sort alphabetically.
	word := make([]string, 0, len(words))
	for w := range words {
		word = append(word, w)
	}
	sort.Strings(word)

	// Print out the final wordcount.
	fmt.Printf("word\tcount\n")
	for _, w := range word {
		fmt.Printf("%s\t%d\n", w, words[w])
	}
}

func removeNonAlphaNumeric(b []byte) []byte {

	count := 0
	tally := 0
	l := len(b)

	for len(b) > 0 {
		r, size := utf8.DecodeLastRune(b)
		if conditional(r) {
			for conditional(r) {
				b = b[:len(b)-size]
				count += size
				r, size = utf8.DecodeLastRune(b)
			}
			// Take a temporary measure of the current length.
			t := len(b)
			// Activate the entire buffer.
			b = b[:l]
			// append the buffer from the end of the unichar
			// whitespace onto the end of the next valid unichar
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

func conditional(r rune) bool {
	return unicode.IsPunct(r) || unicode.IsSpace(r)
}
