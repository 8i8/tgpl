package xkcd

import (
	"fmt"
	"strconv"

	"tgpl/ex_04.12-xkcd/quest"
)

var comics Comics

// getComic returns the latest xkcd comic.
func getComic(i uint) (Comic, int, error) {

	var req quest.HttpQuest
	var comic Comic
	var err error

	if i == 0 {
		req, err = req.QuestGET(cLastURL)
	} else {
		req, err = req.QuestGET(cBaseURL + strconv.Itoa(int(i)) + cTailURL)
	}
	if err != nil {
		return comic, req.Code, fmt.Errorf("QuestGET: %v", err)
	}

	comic, err = xkcdDecode(req)
	if err != nil {
		return comic, req.Code, fmt.Errorf("xkcdDecode: %v", err)
	}

	return comic, req.Code, nil
}

// getLatest returns the latest xkcd comic.
func getLatest() (Comic, int, error) {
	return getComic(0)
}

// getLatestNumber returns the referance number of the latest xkcd comic.
func getLatestNumber() (uint, error) {

	if TESTRUN > 0 {
		return TESTRUN, nil
	}

	latest, _, err := getLatest()
	if err != nil {
		return 0, fmt.Errorf("GetLatestNumber: %v", err)
	}
	return latest.Num(), nil
}

// Init initalises the xkcd datastructure from the database, checking for
// updates if requested, or if the last update was over a week ago.
func Init() {

	var err error
	if VERBOSE {
		quest.VERBOSE = true
	}
	if TESTRUN > 0 {
		cNAME = "test.json"
	}

	if RECORD > 0 {
		comic, code, err := getComic(RECORD)
		if err != nil && code != 404 {
			fmt.Printf("error: getComic: %v\n")
		}
		if code == 404 {
			fmt.Printf("http: 404 page not found\n")
		} else {
			PrintSingle(comic)
		}
		return
	}

	comics, err = LoadDatabase()
	if err != nil {
		fmt.Printf("error: LoadDatabase: %v\n", err)
		return
	}

	if UPDATE {
		Update()
	}

}

func Update() {

	var err error
	comics, _, err = UpdateDatabase(comics)
	if err != nil {
		fmt.Printf("error: UpdateDatabase: %v\n", err)
		return
	}
}

/*
month	"3"
num	2128
link	""
year	"2019"
news	""
safe_title	"New Robot"
transcript	""
alt	"\"Some worry that we'll soon have a surplus of search and rescue robots, compared to the number of actual people in situations requiring search and rescue. That's where our other robot project comes in...\""
img	"https://imgs.xkcd.com/comics/new_robot.png"
title	"New Robot"
day	"25"
*/
func PrintComic(i uint) {
	fmt.Println(comics.Edition[i])
}

func PrintSingle(comic Comic) {
	fmt.Println(comic)
}
