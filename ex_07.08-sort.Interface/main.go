package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"

	"github.com/8i8/csort"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
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

// sortBy returns a sort function tha examins the given struct element.
func sortBy(cmd string) csort.SortFn {
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
		case "artist":
			if x.Artist != y.Artist {
				if x.Artist < y.Artist {
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
		case "year":
			if x.Year != y.Year {
				if x.Year < y.Year {
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
		default:
			panic("unknown command: " + cmd)
		}
		return csort.Equal
	}
}

func main() {
	buf := csort.NewSortFnBuffer(sortBy)
	// buf.Add("length")
	// buf.Add("year")
	buf.Add("title")
	sort.Sort(csort.New(csort.SortFunction(buf), tracks))
	printTracks(tracks)
}
