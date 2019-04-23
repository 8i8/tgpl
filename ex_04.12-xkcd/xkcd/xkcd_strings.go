package xkcd

import "strings"

// cleanArgs breaks the arguments list down in to single words coverting
// them all to lower case and removing any empty strings.
func cleanArgs(args []string) ([]string, bool) {

	var wordlist []string
	for _, arg := range args {
		words := strings.Split(arg, " ")
		for _, word := range words {
			// Remove empty strings.
			if len(word) > 0 {
				wordlist = append(wordlist, strings.ToLower(word))
			}
		}
	}
	if len(wordlist) == 0 {
		return wordlist, false
	}
	return wordlist, true
}
