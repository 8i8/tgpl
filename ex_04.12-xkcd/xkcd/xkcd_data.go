package xkcd

import (
	"encoding/json"
	"fmt"
	"os"
)

// LoadConfig load the last saved config, if no file exists create with the
// default settings.
func LoadDatabase() (Comics, error) {

	// Open file for reading if present.
	file, err := os.Open("xkcd.json")
	if err != nil {
		// If not present create file
		comics, err := generate()
		if err != nil {
			return comics, fmt.Errorf("generate: %v", err)
		}
		return comics, nil
	}

	// Read json file.
	var comics Comics
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&comics)
	if err != nil {
		return comics, fmt.Errorf("decoder: %v", err)
	}

	return comics, err
}

// UpdateDatabase retries any editions that are not currently in the database.
// func UpdateDatabase() (Comics, error) {
// }

// generate pull the entire xkcd database from the website.
func generate() (Comics, error) {

	var comics Comics
	// Make the file.
	file, err := os.Create("xkcd.json")
	if err != nil {
		return comics, fmt.Errorf("Create: %v", err)
	}
	defer file.Close()

	// Get the total quantity of comics.
	// l, err := GetLatestNum()
	// if err != nil || l == 0 {
	// 	return fmt.Errorf("GetLatestNum: %v", err)
	// }
	var l uint
	l = 1000
	comics.Edition = make([]Comic, l)
	comics.Len = l

	// Fill array with comics.
	for i := uint(0); i < l; i++ {
		comic, err := GetComic(i + 1)
		if err != nil {
			return comics, fmt.Errorf("GetComicNum: %d: %v", i, err)
		}
		comics.Edition[i] = comic
	}

	// Prepare for reading.
	data, err := json.MarshalIndent(comics, "", "	")
	if err != nil {
		return comics, fmt.Errorf("MarshalIndent: %v", err)
	}

	// Write json to file.
	_, err = file.Write(data)
	if err != nil {
		return comics, fmt.Errorf("Write: %v", err)
	}

	return comics, err
}
