package xkcd

import "strings"

// cleanArgs breaks the arguments list down in to single words and if the
// appropriate flag is set coverts them all to lower case glyphs.
func cleanArgs(args []string) []string {

	var wordlist []string
	for _, arg := range args {
		words := strings.Split(arg, " ")
		for _, word := range words {
			wordlist = append(wordlist, strings.ToLower(word))
		}
	}
	return wordlist
}
