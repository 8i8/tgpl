package xkcd

const cLastURL = "https://xkcd.com/info.0.json"
const cBaseURL = "https://xkcd.com/"
const cTailURL = "/info.0.json"

// Comics is an array of xkcd cartoons.
type Comics struct {
	Len     uint
	Edition []Comic
}

// Comic contains an xkcd cartoon.
type Comic struct {
	Month      string
	Num        uint
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

/*

https://xkcd.com/json.html

If you want to fetch comics and metadata automatically,
you can use the JSON interface. The URLs look like this:

http://xkcd.com/info.0.json (current comic)

or:

http://xkcd.com/614/info.0.json (comic #614)

Those files contain, in a plaintext and easily-parsed format: comic titles,
URLs, post dates, transcripts (when available), and other metadata.

~~~

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

~~~

month	"7"
num	614
link	""
year	"2009"
news	""
safe_title	"Woodpecker"
transcript	"[[A man with a beret and a woman are standing on a boardwalk,
	leaning on a handrail.]]\nMan: A woodpecker!\n<<Pop pop pop>>\nWoman:
	Yup.\n\n[[The woodpecker is banging its head against a tree.]]\nWoman:
	He hatched about this time last year.\n<<Pop pop pop pop>>\n\n[[The
	woman walks away.  The man is still standing at the handrail.]]\n\nMan:
	...  woodpecker?\nMan: It's your birthday!\n\nMan: Did you
	know?\n\nMan: Did... did nobody tell you?\n\n[[The man stands,
	looking.]]\n\n[[The man walks away.]]\n\n[[There is a tree.]]\n\n[[The
	man approaches the tree with a present in a box, tied up with
	ribbon.]]\n\n[[The man sets the present down at the base of the tree
	and looks up.]]\n\n[[The man walks away.]]\n\n[[The present is sitting
	at the bottom of the tree.]]\n\n[[The woodpecker looks down at the
	present.]]\n\n[[The woodpecker sits on the present.]]\n\n[[The
	woodpecker pulls on the ribbon tying the present closed.]]\n\n((full
	width panel))\n[[The woodpecker is flying, with an electric drill
	dangling from its feet, held by the cord.]]\n\n{{Title text: If you
	don't have an extension cord I can get that too.  Because we're
	friends!  Right?}}"
alt	"If you don't have an extension cord I can get that too.  Because we're
	friends!  Right?"
img	"https://imgs.xkcd.com/comics/woodpecker.png"
title	"Woodpecker"
day	"24"

*/
