package xkcd

import (
	"encoding/json"
	"fmt"
	"os"
)

func loadDatabase() (Comics, error) {

	if VERBOSE {
		fmt.Printf("xkcd: loading database\n")
	}

	// Open file for reading if present.
	file, err := os.Open(cNAME)
	if err != nil {
		if VERBOSE {
			fmt.Printf("xkcd: data file not found ...\n")
		}
		//return nil, fmt.Errorf("Open: %v", err)
		// TODO ask user whether a new database is to be made.
		comics, err := generateDatabase()
		if err != nil {
			return comics, fmt.Errorf("generateDatabase: %v", err)
		}
		return comics, nil
	}

	comics, err := readFile(file)
	if err != nil {
		return comics, fmt.Errorf("readFile: %v", err)
	}

	if VERBOSE {
		fmt.Printf("xkcd: database loaded\n")
	}

	return comics, err
}

func growDatastructure(comics Comics, l uint) Comics {

	// If required, make a new data structure and copy over any comics
	// present.
	if cap(comics.Edition) < int(l) {
		newComics := Comics{}
		newComics.Edition = make([]Comic, l)
		copy(newComics.Edition, comics.Edition)
		return newComics
	}
	// Adjust length.
	comics.Edition = comics.Edition[:int(l)]
	return comics
}

// UpdateDatabase retrieves any editions that are not currently in the
// database.
func UpdateDatabase(comics Comics) (Comics, bool, error) {

	if VERBOSE {
		fmt.Printf("xkcd: updating database ...\n")
	}

	// Check whether update is required.
	l, err := getLatestNumber()
	if err != nil {
		return comics, false, fmt.Errorf("GetLatestNumber: %v", err)
	}

	if l <= comics.Len {
		if VERBOSE {
			fmt.Printf("\nxkcd: ... up to date\n")
		}
		return comics, false, nil
	}

	// Record current length and grow.
	c := comics.Len
	comics = growDatastructure(comics, l)

	// Fill array with new comics.
	for i := uint(c); i < l; i++ {
		comic, code, err := getComic(i + 1)
		if err != nil && code != 404 {
			return comics, false, fmt.Errorf("GetComicNum: %d: %v", i, err)
		}
		if code != 404 {
			comics.Edition[i] = comic
		}
		comics.Len++
	}

	// Adds \n after the quest.http status responces which use \r to avoid
	// flooding.
	if VERBOSE {
		fmt.Printf("\n")
	}

	// Prepare for reading.
	data, err := json.MarshalIndent(comics, "", "	")
	if err != nil {
		return comics, false, fmt.Errorf("MarshalIndent: %v", err)
	}

	// Open file for writing.
	file, err := os.Create(cNAME)
	if err != nil {
		return comics, false, fmt.Errorf("Create: %v", err)
	}

	// Write data to file.
	_, err = file.Write(data)
	if err != nil {
		return comics, false, fmt.Errorf("Write: %v", err)
	}
	file.Close()

	if VERBOSE {
		fmt.Printf("xkcd: ... database updated, %d records added\n", l-c)
	}

	return comics, true, nil
}

// generateDatabase pulls the entire xkcd database from the website.
func generateDatabase() (Comics, error) {

	var comics Comics

	if VERBOSE {
		fmt.Printf("xkcd: generating database ...\n")
	}

	// Get the total quantity of comics.
	l, err := getLatestNumber()
	if err != nil || l == 0 {
		return comics, fmt.Errorf("GetLatestNum: %v", err)
	}

	// Data store.
	comics.Edition = make([]Comic, l)
	comics.Len = l

	// Fill array with comics.
	for i := uint(0); i < l && i < 400; i++ {
		comic, code, err := getComic(i + 1)
		if err != nil && code != 404 {
			return comics, fmt.Errorf("GetComicNum: %d: %v", i, err)
		}
		if code == 404 {
			fmt.Printf("http: 404 page %d not found\n", i+1)
		} else {
			comics.Edition[i] = comic
		}
	}

	if VERBOSE {
		fmt.Printf("\n")
	}

	// Prepare for reading.
	data, err := json.MarshalIndent(comics, "", "	")
	if err != nil {
		return comics, fmt.Errorf("MarshalIndent: %v", err)
	}

	// Open file for writing.
	file, err := os.Create(cNAME)
	if err != nil {
		return comics, fmt.Errorf("Create: %v", err)
	}
	defer file.Close()

	// Write data to file.
	_, err = file.Write(data)
	if err != nil {
		return comics, fmt.Errorf("Write: %v", err)
	}

	if VERBOSE {
		fmt.Printf("xkcd: ... database written\n")
	}

	return comics, err
}
