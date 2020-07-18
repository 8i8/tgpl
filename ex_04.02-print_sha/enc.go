// enc requires the package github.com/atotto/clipboard
package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"

	"github.com/atotto/clipboard"
)

// MD5 Input flags.
var MD5 = flag.Bool("md5", false, "Output a md5 hash.")

// NEWLINE Input flags.
var NEWLINE = flag.Bool("n", false, "Output the hash generated when a newline char is included in the string to be hashed.")

// SHA256 Input flags.
var SHA256 = flag.Bool("sha256", false, "Output a sha512 hash.")

// SHA384 Input flags.
var SHA384 = flag.Bool("sha384", false, "Output a sha384 hash.")

// SHA512 Input flags.
var SHA512 = flag.Bool("sha512", false, "Output a sha512 hash.")

func main() {

	// Get flags from the os on program start.
	flag.Parse()

	for {
		// Read stdin.
		scanner := bufio.NewScanner(os.Stdin)

		if !scanner.Scan() {
			return
		}

		input := scanner.Bytes()
		if len(input) == 0 {
			return
		}
		if *NEWLINE {
			input = append(input, '\n')
		}

		// Set the appropriate hash for the input, the given flag or lack there
		// of, set the hashing algorithm. The resulting hash is copied to the system
		// clipboard and then written to stdout.
		if *MD5 {
			str := md5.Sum(input)
			clipboard.WriteAll(fmt.Sprintf("%x", str))
			fmt.Printf("%x\n", str)
		} else if *SHA256 {
			str := sha256.Sum256(input)
			clipboard.WriteAll(fmt.Sprintf("%x", str))
			fmt.Printf("%x\n", str)
		} else if *SHA384 {
			str := sha512.Sum384(input)
			clipboard.WriteAll(fmt.Sprintf("%x", str))
			fmt.Printf("%x\n", str)
		} else if *SHA512 {
			str := sha512.Sum512(input)
			clipboard.WriteAll(fmt.Sprintf("%x", str))
			fmt.Printf("%x\n", str)
		} else {
			// Default hash is md5
			str := md5.Sum(input)
			clipboard.WriteAll(fmt.Sprintf("%x", str))
			fmt.Printf("%x\n", str)
		}
	}
}
