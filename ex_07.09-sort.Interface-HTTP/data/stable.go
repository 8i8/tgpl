package data

import (
	"log"
)

type Tlist struct {
	tr   []*Track
	less func(x, y *Track) bool
}

func (t Tlist) Len() int           { return len(t.tr) }
func (t Tlist) Less(i, j int) bool { return t.less(t.tr[i], t.tr[j]) }
func (t Tlist) Swap(i, j int)      { t.tr[i], t.tr[j] = t.tr[j], t.tr[i] }

func (t *Tlist) stableSort(tracks Tlist, cmd string) *Tlist {
	t.less = func(x, y *Track) bool {
		if res := less(x, y, cmd); res > Equal {
			if res == Left {
				return true
			}
			return false
		}
		return false
	}
	return t
}

const (
	Equal = iota
	Left
	Right
)

func less(x, y *Track, cmd string) int {
	switch cmd {
	case "title":
		if x.Title != y.Title {
			if x.Title < y.Title {
				return Left
			}
			return Right
		}
	case "title-rev":
		if x.Title != y.Title {
			if x.Title > y.Title {
				return Left
			}
			return Right
		}
	case "artist":
		if x.Artist != y.Artist {
			if x.Artist < y.Artist {
				return Left
			}
			return Right
		}
	case "artist-rev":
		if x.Artist != y.Artist {
			if x.Artist > y.Artist {
				return Left
			}
			return Right
		}
	case "album":
		if x.Album != y.Album {
			if x.Album < y.Album {
				return Left
			}
			return Right
		}
	case "album-rev":
		if x.Album != y.Album {
			if x.Album > y.Album {
				return Left
			}
			return Right
		}
	case "year":
		if x.Year != y.Year {
			if x.Year < y.Year {
				return Left
			}
			return Right
		}
	case "year-rev":
		if x.Year != y.Year {
			if x.Year > y.Year {
				return Left
			}
			return Right
		}
	case "length":
		if x.Length != y.Length {
			if x.Length < y.Length {
				return Left
			}
			return Right
		}
	case "length-rev":
		if x.Length != y.Length {
			if x.Length > y.Length {
				return Left
			}
			return Right
		}
	case "", "NUL":
	default:
		log.Fatal("unknown command: " + cmd)
	}
	return Equal
}
