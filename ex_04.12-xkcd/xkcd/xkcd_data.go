package xkcd

import (
	"encoding/json"
	"fmt"
	"os"
)

// DataBase is an array of xkcd cartoons.
type DataBase struct {
	Len     int
	Edition []Comic
}

// xkcdInit loads the xkcd index data from the data file into program memory.
func xkcdInit() (*DataBase, error) {

	// Data store for program whilst running.
	var comics DataBase

	// load db into memory.
	err := loadDatabase(&comics, cNAME)
	if err != nil {
		return &comics, fmt.Errorf("loadDatabase: %v", err)
	}

	return &comics, err
}

// loadDatabase loads xlcd data from a file if one exists, downloading and
// creating the file if it is not present.
func loadDatabase(comics *DataBase, path string) error {

	if VERBOSE {
		fmt.Printf("xkcd: loading database\n")
	}

	// Open file for reading if present.
	file, err := os.Open(path)
	if err != nil {
		if VERBOSE {
			fmt.Printf("xkcd: data file not found ...\n")
		}
		//return nil, fmt.Errorf("Open: %v", err)
		// TODO ask user whether a new database is to be made.
		comics, err = downloadAllXkcd(comics)

		if err != nil {
			return fmt.Errorf("downloadAllXkcd: %v", err)
		}
		return nil
	}

	comics, err = readFile(comics, file)
	if err != nil {
		return fmt.Errorf("readFile: %v", err)
	}

	if VERBOSE {
		fmt.Printf("xkcd: database loaded\n")
	}

	return err
}

// growDatastructure expands the array that holds the comic structs to a
// specified length, the latest comic is retrievable as such the required
// length is known.
func growDatastructure(comics *DataBase, l int) *DataBase {

	// If required, make a new data structure and copy over any comics
	// present.
	if cap(comics.Edition) < int(l) {
		newDataBase := &DataBase{}
		newDataBase.Edition = make([]Comic, l)
		copy(newDataBase.Edition, comics.Edition)
		return newDataBase
	}
	// Adjust length.
	comics.Edition = comics.Edition[:int(l)]
	return comics
}

// updateDatabase retrieves any editions that are not currently in the
// database and adds them.
func updateDatabase(comics *DataBase) (*DataBase, bool, error) {

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

	// Record current length and grow to accommodate new records if
	// required.
	c := comics.Len
	comics = growDatastructure(comics, l)

	// Fill array with new comics.
	for i := int(c); i < l; i++ {
		comic, code, err := newComicHTTP(i + 1)
		if err != nil && code != 404 {
			return comics, false, fmt.Errorf("newComicHTTP: %d: %v", i, err)
		}
		if code != 404 {
			comics.Edition[i] = comic
		}
		comics.Len++
	}

	// Adds \n after the quest.http status responses which use \r to avoid
	// flooding.
	if VERBOSE {
		fmt.Printf("\n")
	}

	err = writeDatabase(comics)
	if err != nil {
		return comics, false, fmt.Errorf("writeDatabase: %v", err)
	}

	if VERBOSE {
		fmt.Printf("xkcd: ... database updated, %d records added\n", l-c)
	}

	return comics, true, nil
}

// downloadAllXkcd pulls the entire xkcd database from the website.
func downloadAllXkcd(comics *DataBase) (*DataBase, error) {

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
	for i := int(0); i < l && i < 400; i++ {
		comic, code, err := newComicHTTP(i + 1)
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

	err = writeDatabase(comics)
	if err != nil {
		return comics, fmt.Errorf("writeDatabase: %v", err)
	}

	return comics, nil
}

// writeDatabase writes a database struct to file in a readable json format.
func writeDatabase(comics *DataBase) error {

	if VERBOSE {
		fmt.Printf("xkcd: opening database ...\n")
	}

	// Prepare for reading.
	data, err := json.MarshalIndent(comics, "", "	")
	if err != nil {
		return fmt.Errorf("MarshalIndent: %v", err)
	}

	// Open file for writing.
	file, err := os.Create(cNAME)
	if err != nil {
		return fmt.Errorf("Create: %v", err)
	}
	defer file.Close()

	// Write data to file.
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("Write: %v", err)
	}

	if VERBOSE {
		fmt.Printf("xkcd: ... database written\n")
	}
	return err
}

// Update updates the comic database with the latest comic descriptions,
// returns true if the database has been updated or false if there are no new
// comics.
func (d *DataBase) Update() bool {

	_, ok, err := updateDatabase(d)
	if err != nil {
		fmt.Printf("error: updateDatabase: %v\n", err)
		return false
	}
	return ok
}

// DbGet prints out the given comic description from the database, requesting
// comic 0 will raise the most recent comic.
func (d *DataBase) DbGet(n int) {

	if VERBOSE {
		fmt.Printf("xkcd: database access ~~~\n\n")
	}
	if d.Len >= n && n > 0 {
		d.printComic(n)
		// If DBGET == 0 print the most recent comic.
	} else if n == 0 {
		d.printComic(d.Len)
	} else {
		fmt.Printf("xkcd: most recent addition is Number %d.", d.Len)
	}
	if VERBOSE {
		fmt.Printf("\nxkcd: ~~~ database done\n")
	}
}
