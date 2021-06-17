package data

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"
	"time"

	"sortInterface/data/csort"
)

var verbose bool

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var Tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

// Less returns a sort function tha examins the given struct element.
func Less(cmd string) csort.SortFn {
	return func(xi, yi interface{}) int {
		x, ok := xi.(*Track)
		if !ok {
			panic("x interface conversion failed")
		}
		y, ok := yi.(*Track)
		if !ok {
			panic("y interface conversion failed")
		}
		switch cmd {
		case "title":
			if x.Title != y.Title {
				if x.Title < y.Title {
					return csort.Left
				}
				return csort.Right
			}
		case "title-rev":
			if x.Title != y.Title {
				if x.Title > y.Title {
					return csort.Left
				}
				return csort.Right
			}
		case "artist":
			if x.Artist != y.Artist {
				if x.Artist < y.Artist {
					return csort.Left
				}
				return csort.Right
			}
		case "artist-rev":
			if x.Artist != y.Artist {
				if x.Artist > y.Artist {
					return csort.Left
				}
				return csort.Right
			}
		case "album":
			if x.Album != y.Album {
				if x.Album < y.Album {
					return csort.Left
				}
				return csort.Right
			}
		case "album-rev":
			if x.Album != y.Album {
				if x.Album > y.Album {
					return csort.Left
				}
				return csort.Right
			}
		case "year":
			if x.Year != y.Year {
				if x.Year < y.Year {
					return csort.Left
				}
				return csort.Right
			}
		case "year-rev":
			if x.Year != y.Year {
				if x.Year > y.Year {
					return csort.Left
				}
				return csort.Right
			}
		case "length":
			if x.Length != y.Length {
				if x.Length < y.Length {
					return csort.Left
				}
				return csort.Right
			}
		case "length-rev":
			if x.Length != y.Length {
				if x.Length > y.Length {
					return csort.Left
				}
				return csort.Right
			}
		case "":
		default:
			log.Fatal("unknown command: " + cmd)
		}
		return csort.Equal
	}
}

// wrapInter returns the given slice as a slice of interfaces.
func wrapInter(tracks []*Track) []interface{} {
	inter := make([]interface{}, len(tracks))
	for i := range tracks {
		inter[i] = tracks[i]
	}
	return inter
}

func unwrapInter(inter []interface{}, tracks []*Track) {
	for i := range inter {
		tracks[i] = inter[i].(*Track)
	}
}

func Sort(buf *csort.SortBuffer, tracks []*Track) []*Track {
	inter := wrapInter(Tracks)
	sort.Sort(csort.New(buf.LoadSortFn(), inter))
	unwrapInter(inter, Tracks)
	return Tracks
}

func StableSort(tracks []*Track, cmd string) []*Track {
	t := Tlist{tr: tracks}
	sort.Stable(t.stableSort(t, cmd))
	return t.tr
}

func Data() (*csort.SortBuffer, []*Track) {
	buf := csort.NewSortBuffer(Less)
	return buf, Tracks
}
