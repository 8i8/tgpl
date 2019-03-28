package xkcd

import (
	"fmt"
	"strconv"

	"tgpl/ex_04.12-xkcd/quest"
)

// GetComic returns the latest xkcd comic.
func GetComic(i uint) (Comic, error) {

	var req quest.HttpQuest
	var comic Comic
	var err error

	if i == 0 {
		req, err = req.QuestGET(cLastURL)
	} else {
		req, err = req.QuestGET(cBaseURL + strconv.Itoa(int(i)) + cTailURL)
	}
	if err != nil {
		return comic, fmt.Errorf("QuestGET: %v", err)
	}

	comic, err = xkcdDecode(req)
	if err != nil {
		return comic, fmt.Errorf("xkcdDecode: %v", err)
	}

	return comic, nil
}

// GetLatest returns the latest xkcd comic.
func GetLatest() (Comic, error) {
	return GetComic(0)
}

// GetLatestNumber returns the referance number of the latest xkcd comic.
func GetLatestNumber() (uint, error) {

	latest, err := GetLatest()
	if err != nil {
		return 0, fmt.Errorf("GetLatestNumber: %v", err)
	}
	return latest.Num, err
}
