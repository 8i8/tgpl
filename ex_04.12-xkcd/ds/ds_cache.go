package ds

import (
	"fmt"
	"io"
	"os"

	"tgpl/ex_04.12-xkcd/msg"
)

func serialiseTrieToFile(t *Trie, name string) error {

	// Open cache file for writing.
	file, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("file.Create: %v", err)
	}

	w := msg.NewErrWriter(file)

	err = t.SerialiseTrie(w)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	// Finish up
	err = file.Close()
	if err != nil {
		return fmt.Errorf("file.Close: %v", err)
	}
	return err
}

// deserialiseTrieFromFile reconstitutes a trie data structure from its serialised cache
// file.
func deserialiseTrieFromFile(t *Trie, name string) error {

	if d&VERBOSE > 0 {
		fmt.Printf("ds: trie deserialisation started ...\n")
	}

	t.start = new(node)

	// Open cache file for reading.
	file, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("file.Open: %v", err)
	}
	rw := msg.NewErrReader(file)

	// Perform deserialisation.
	_, err = deserialiseTrie(t.start, rw)
	if err != nil && err != io.EOF {
		return fmt.Errorf("deserialise: %v", err)
	} else if err == io.EOF {
		if d&VERBOSE > 0 {
			println("EOF reached")
		}
	}

	// Finish up.
	err = file.Close()
	if err != nil {
		return fmt.Errorf("file.Close: %v", err)
	}

	if d&VERBOSE > 0 {
		fmt.Printf("ds: ... trie deserialisation done\n")
	}

	return err
}
