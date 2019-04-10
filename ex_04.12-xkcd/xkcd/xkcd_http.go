package xkcd

import (
	"fmt"
	"strconv"

	"tgpl/ex_04.12-xkcd/quest"
)

// newComicHTTP returns the specified comic from the website.
func newComicHTTP(i uint) (Comic, int, error) {

	var req quest.HttpQuest
	var comic Comic
	var err error

	// If a number has been given get that, else if zero given, get the latest.
	if i == 0 {
		req, err = req.QuestGET(cLastURL)
	} else {
		req, err = req.QuestGET(cBaseURL + strconv.Itoa(int(i)) + cTailURL)
	}
	if err != nil {
		return comic, req.Code, fmt.Errorf("QuestGET: %v", err)
	}

	// If a comic was returned decode it.
	if req.Code != 404 {
		comic, err = xkcdDecode(req)
	}
	if err != nil {
		return comic, req.Code, fmt.Errorf("xkcdDecode: %v", err)
	}

	return comic, req.Code, nil
}

// getLatestHTTP returns the latest xkcd comic.
func newLatestComicHTTP() (Comic, int, error) {
	return newComicHTTP(0)
}

// getLatestNumber returns the referance number of the latest xkcd comic.
func getLatestNumber() (uint, error) {

	if TESTRUN > 0 {
		return TESTRUN, nil
	}

	latest, _, err := newLatestComicHTTP()
	if err != nil {
		return 0, fmt.Errorf("GetLatestNumber: %v", err)
	}
	return latest.Num(), nil
}
