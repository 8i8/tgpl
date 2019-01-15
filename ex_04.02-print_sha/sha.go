package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	// A simple program that uses xclip to write to xservers clipboard on a
	// linux system, I have not tested on other machines.
	"github.com/atotto/clipboard"
	"log"
	"os"
)

// Input flags.
var SHA256 = flag.Bool("sha256", false, "Output a sha512 hash.")
var SHA384 = flag.Bool("sha384", false, "Output a sha384 hash.")
var SHA512 = flag.Bool("sha512", false, "Output a sha512 hash.")

func main() {

	// Get flags from the os on program start.
	flag.Parse()

	// Read stdin.
	scanner := bufio.NewScanner(os.Stdin)
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
		clipboard.WriteAll(fmt.Sprintf("%x", sha[:]))
		fmt.Printf("%x\n", sha)
	} else if *SHA384 {
		sha := sha512.Sum384([]byte(input))
		clipboard.WriteAll(fmt.Sprintf("%x", sha[:]))
		fmt.Printf("%x\n", sha)
	} else if *SHA512 {
		sha := sha512.Sum512([]byte(input))
		clipboard.WriteAll(fmt.Sprintf("%x", sha[:]))
		fmt.Printf("%x\n", sha)
	} else {
		// Default hash is md5
		sha := md5.Sum([]byte(input))
		clipboard.WriteAll(fmt.Sprintf("%x", sha[:]))
		fmt.Printf("%x\n", sha)
	}
}
