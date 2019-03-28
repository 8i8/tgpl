package xkcd

import (
	"encoding/json"
	"fmt"

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
