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

// setConfig sets required state variables for desired program run mode.
func setConfig() {

	// sets quest package to verbose.
	if VERBOSE {
		quest.VERBOSE = true
	}

	// Sets program to generate a test database.
	if TESTRUN > 0 {
		cNAME = "test.json"
	}
}

// dbInit loads the xkcd index data from the data file into program memory.
func dbInit() Comics {

	// load db into memory.
	comics, err := loadDatabase()
	if err != nil {
		fmt.Printf("error: loadDatabase: %v\n", err)
		return comics
	}
	return comics
}

// DbGet prints our the given comic description from the database.
func DbGet(n uint) {

	if VERBOSE {
		fmt.Printf("xkcd: database access ~~~\n\n")
	}
	if comics.Len > DBGET {
		printSingle(comics.Edition[n])
	}
	if VERBOSE {
		fmt.Printf("\nxkcd: ~~~ database done\n")
	}
}

// WebGet outputs the given comic from the internet by making a http request.
func WebGet(n uint) {

	if VERBOSE {
		fmt.Printf("xkcd: make http request\n")
	}

	comic, code, err := getComic(n)
	if err != nil && code != 404 {
		fmt.Printf("error: getComic: %v\n", err)
	} else if code == 404 {
		fmt.Printf("xkcd: http: 404 comic not found\n")
	}

	if VERBOSE {
		fmt.Printf("xkcd: http request closed\n")
	}

	if err == nil {
		printSingle(comic)
	}
}

// Search searched the local database of comic descriptions for the given
// arguments.
func Search(args []string) {
	if VERBOSE {
		fmt.Printf("xkcd: output start ~~~\n\n")
	}
	m := buildSearchMap(comics)
	buildSearchGraph(m)
	res := search(m, comics, args)
	printResults(comics, res)
	if VERBOSE {
		fmt.Printf("\nxkcd: ~~~ output end\n")
	}
}

// Update updates the comic database with the latest comic descriptions.
func Update() {

	var err error
	comics, _, err = UpdateDatabase(comics)
	if err != nil {
		fmt.Printf("error: UpdateDatabase: %v\n", err)
		return
	}
}

// Run is the xkcd main program routine.
func Run(args []string) {

	// Program state.
	setConfig()

	// Load db
	comics = dbInit()

	if DBGET > 0 {
		DbGet(DBGET - 1)
		return
	}

	if SEARCH {
		Search(args)
		return
	}

	if WEBGET > 0 {
		WebGet(WEBGET)
	}

	if UPDATE {
		Update()
	}

	// TODO the numbering needs to be made generic.
	printSingle(comics.Edition[comics.Len-1])
}
