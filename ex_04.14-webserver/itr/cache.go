package itr

import (
	"encoding/json"
	"fmt"
	"os"

	"tgpl/ex_04.14-webserver/github"
)

// loadCache loads a repos data from a local cache file.
func loadCache(name, repo string) (*[]github.Issue, bool, error) {

	filepath := "cache/" + name + "." + repo + ".json"

	if VERBOSE {
		fmt.Printf("loadCache: checking for cache file\n")
	}

	// If the  file is not ther return.
	if _, err := os.Stat(filepath); err != nil {
		if VERBOSE {
			fmt.Printf("loadCache: cache not found\n")
		}
		return nil, false, nil
	}

	if VERBOSE {
		fmt.Printf("loadCache: opening cache file for reading\n")
	}

	file, err := os.Open(filepath)
	if err != nil {
		return nil, false, fmt.Errorf("os.Open: %v", err)
	}

	var data []github.Issue
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return nil, false, fmt.Errorf("Decode: %v", err)
	}

	if VERBOSE {
		fmt.Printf("loadCache: cache loaded\n")
	}

	return &data, true, nil
}

// writeCache writes the current cache data to file.
func writeCache(name, repo string, data *[]github.Issue) error {

	if VERBOSE {
		fmt.Printf("writeCache: opening cache file for writing\n")
	}

	jsondata, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return fmt.Errorf("MarshalIndetn: %v", err)
	}

	file, err := os.Create("cache/" + name + "." + repo + ".json")
	if err != nil {
		return fmt.Errorf("os.Create: %v", err)
	}
	defer file.Close()

	_, err = file.Write(jsondata)
	if err != nil {
		return fmt.Errorf("file.Write: %v", err)
	}

	if VERBOSE {
		fmt.Printf("writeCache: cache writen\n")
	}

	return nil
}
