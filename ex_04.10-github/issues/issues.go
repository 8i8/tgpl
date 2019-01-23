package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"tgpl/ex_04.10-github/github"
)

type date struct {
	r int
	y int
	m time.Month
	d int
	p bool
}

func main() {
	ny, nm, nd := time.Now().Date()
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	// Make and fill a map from the result array.
	issue := make(map[int]github.Issue)
	for _, item := range result.Items {
		issue[item.Number] = *item
	}

	// Make and fill an array to index the issues.
	index := make([]date, 0, len(issue))
	for r, item := range issue {
		y, m, d := item.CreatedAt.Date()
		p := false
		issue := date{r, y, m, d, p}
		index = append(index, issue)
	}

	// Sort the results, order by referance numbers.
	sort.Slice(index, func(i, j int) bool {
		return index[i].r > index[j].r
	})

	// 1) The same calendar month and year.
	// 2) The same calendar year with one month difference; lesser day than
	//	the day issued.
	// 3) One calendar year difference; January issue date December; lesser
	//	day today than the day issued.
	fmt.Println("\nLess than a month")
	for n, i := range index {
		item := issue[i.r]
		if i.y == ny && i.m == nm ||
			i.y == ny-1 && i.m == nm && i.d > nd ||
			i.y == ny-1 && nm == 1 && i.m == 12 && i.d > nd {
			fmt.Printf("#%-5d %9.9s %55.55s  %.2d %s %d\n",
				item.Number, item.User.Login, item.Title, i.d, i.m, i.y)
			// Verify date logic
			if i.p == true {
				fmt.Println("error: this line has already been printed")
			}
			index[n].p = true
		}
	}

	// 1) The same calendar year; Greater month by more than one
	//	calendar month than the date issued.
	// 2) The same calendar year; The month is one greater than the issued
	//	month; equal or greater day today than the day issued.
	// 3) One calendar years difference; January the date of issue in
	//	December; Equal or greater day today than the day issued.
	// 4) One calendar years difference; The current calendar month is lesser than
	//	the month of the issue.
	// 5) One calendar years difference; The same calendar month; lesser
	//	day today than the day issued.
	fmt.Println("\nLess than a year")
	for n, i := range index {
		item := issue[i.r]
		if i.y == ny && i.m < nm-1 ||
			i.y == ny && i.m == nm-1 && i.d <= nd ||
			i.y == ny-1 && nm == 1 && i.m == 12 && i.d <= nd ||
			i.y == ny-1 && i.m > nm ||
			i.y == ny-1 && i.m == nm && i.d > nd {
			fmt.Printf("#%-5d %9.9s %55.55s  %.2d %s %d\n",
				item.Number, item.User.Login, item.Title, i.d, i.m, i.y)
			// Verify date logic
			if i.p == true {
				fmt.Println("error: this line has already been printed")
			}
			index[n].p = true
		}
	}

	// 1) One calendar years difference; The same calendar month: Equal or
	//	greater day than the day issued.
	// 2) Greater than one years difference.
	fmt.Println("\nOlder than a year")
	for n, i := range index {
		item := issue[i.r]
		if i.y < ny-1 && i.m == nm && i.d <= nd ||
			i.y < ny-1 {
			fmt.Printf("#%-5d %9.9s %55.55s  %.2d %s %d\n",
				item.Number, item.User.Login, item.Title, i.d, i.m, i.y)
			// Verify date logic
			if i.p == true {
				fmt.Println("error: this line has already been printed")
			}
			index[n].p = true
		}
	}

	// Verify date logic
	for _, i := range index {
		if i.p == false {
			fmt.Println("error: this line has not been printed", i.d, i.m, i.y)
		}
	}
}
