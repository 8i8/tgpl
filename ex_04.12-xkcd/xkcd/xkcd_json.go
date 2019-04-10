package xkcd

import (
	"encoding/json"
	"fmt"
	"os"

	"tgpl/ex_04.12-xkcd/quest"
)

// xkcdDecode decodes the raw json responce and marshals it into a Comic
// struct.
func xkcdDecode(req quest.HttpQuest) (Comic, error) {

	var comic Comic
	err := json.Unmarshal(req.Msg.(json.RawMessage), &comic)
	if err != nil {
		return comic, fmt.Errorf("Unmarshal: %v", err)
	}

	return comic, nil
}

// readFile is a wraper around the json decode function.
func readFile(comics *DataBase, file *os.File) (*DataBase, error) {

	err := json.NewDecoder(file).Decode(comics)
	if err != nil {
		return comics, fmt.Errorf("decoder: %v", err)
	}
	return comics, nil
}
