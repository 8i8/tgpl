// Exercise 7.8: Many GUI's provide a table widget with a stateful
// multi-tier sort: the primary sort key is the most recently clicked
// column head, and so on. Define an implementation of sort.Interface
// for use by such a table. Compare that approach with repeated sorting
// using sort.Stable.
package main

import (
	"fmt"

	"sortInterface/data"
)

func main() {
	buf, tracks := data.Init()

	buf.Add("year")
	buf.Add("title")
	data.Sort(buf, tracks)
	data.Print(tracks)
	fmt.Println()

	buf.Add("year")
	buf.Add("year")
	buf.Add("title")
	data.Sort(buf, tracks)
	data.Print(tracks)
	fmt.Println()

	data.StableSort(tracks, "year")
	data.StableSort(tracks, "title")
	data.Print(tracks)
	fmt.Println()

	data.StableSortRev(tracks, "year")
	data.StableSort(tracks, "title")
	data.Print(tracks)
}
