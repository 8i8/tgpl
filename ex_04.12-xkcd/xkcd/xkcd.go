package xkcd

const cLastURL = "https://xkcd.com/info.0.json"
const cBaseURL = "https://xkcd.com/"
const cTailURL = "/info.0.json"

// Database file name.
var cNAME = "xkcd.json"

// Verbouse program output whilst running.
var (
	VERBOSE   bool
	UPDATE    bool
	SEARCH    bool
	DBGET  uint
	WEBGET uint
	TESTRUN   uint
)

// Comics is an array of xkcd cartoons.
type Comics struct {
	Len     uint
	Edition []Comic
}

// Comic contains an xkcd cartoon.
type Comic struct {
	Month      string
	Number     uint `json:"num"`
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

// Print returns a single xkcd commic edition specified by referance number.
func (c Comics) Print(i uint) Comic {
	return c.Edition[i]
}

// Num returns the comic number.
func (c Comic) Num() uint {
	return c.Number
}
