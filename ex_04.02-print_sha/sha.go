package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
	"os"
	"unicode"
	"unicode/utf8"
)

// Input flags.
var SHA256 = flag.Bool("sha256", false, "Output a sha512 hash.")
var SHA384 = flag.Bool("sha384", false, "Output a sha384 hash.")
var SHA512 = flag.Bool("sha512", false, "Output a sha512 hash.")

// firstLetterToUpper converts the first occurance of a letter to upper case.
func firstLetterToUpper(in []byte) []byte {

	fmt.Printf("bytes in:  %x\n", in)
	i := 0
	count := 0

	for len(in) > 0 {
		fmt.Printf("char:    %c\n", in)
		char, size := utf8.DecodeLastRune(in)
		if unicode.IsLetter(char) {
			fmt.Printf("char:    %c\n", char)

			t := len(in)
			in = in[:len(in)-size]

			n := utf8.EncodeRune(in[i:], char)
			if size != n {
				fmt.Errorf("rune to long for insertion. index: %d.", i)
				return nil
			}
			break
		}
		i += size
	}
	fmt.Printf("runes out: %x\n", in)

	in = []byte(string(in))
	fmt.Printf("bytes out: %x\n", in)

	return in
}

func main() {

	// Get flags from the os on program start.
	flag.Parse()

	// Read stdin.
	scanner := bufio.NewScanner(os.Stdin)

	//temp := "hello"
	//scanner := bufio.NewScanner(strings.NewReader(temp))

	if !scanner.Scan() {
		log.Printf("Failed to read: %v", scanner.Err())
		return
	}
	input := scanner.Bytes()

	// Set the appropriate hash for the input, the given flag or lack there
	// of, set the hashing algorithm. The resulting hash is copied to the system
	// clipboard and then written to stdout.
	if *SHA256 {
		sha := sha256.Sum256([]byte(input))
		str := firstLetterToUpper(sha[:])
		//clipboard.WriteAll(fmt.Sprintf("%x", str))
		fmt.Printf("%x\n", str)
	} else if *SHA384 {
		sha := sha512.Sum384([]byte(input))
		str := firstLetterToUpper(sha[:])
		//clipboard.WriteAll(fmt.Sprintf("%x", str))
		fmt.Printf("%x\n", str)
	} else if *SHA512 {
		sha := sha512.Sum512([]byte(input))
		str := firstLetterToUpper(sha[:])
		//clipboard.WriteAll(fmt.Sprintf("%x", str))
		fmt.Printf("%x\n", str)
	} else {
		// Default hash is md5
		sha := md5.Sum([]byte(input))
		str := firstLetterToUpper(sha[:])
		//clipboard.WriteAll(fmt.Sprintf("%x", str))
		fmt.Printf("%x\n", str)
	}
}
