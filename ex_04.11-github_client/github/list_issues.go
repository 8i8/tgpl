package github

import (
	"fmt"
	"log"
	"sort"
	"time"
)

type date struct {
	r int        // reference
	y int        // year
	m time.Month // month
	d int        // day
	p bool
}

// listIssues Retrieves a list of issues from the given repo that meet the
// given search criteria.
func ListIssues(terms []string) {
	// Retrieve data
	result, err := SearchIssues(terms)
	if err != nil {
		log.Fatal(err)
	}
	// Make and fill a map from the result array.
	issue := make(map[int]Issue)
	for _, item := range result.Items {
		issue[item.Number] = *item
	}
	// Make and fill an array to index the map.
	index := make([]date, 0, len(issue))
	for r, item := range issue {
		y, m, d := item.CreatedAt.Date()
		p := false
		issue := date{r, y, m, d, p}
		index = append(index, issue)
	}
	// Sort the results, order by reference numbers.
	sort.Slice(index, func(i, j int) bool {
		return index[i].r > index[j].r
	})

	printIssues(issue, index)
}

// Print to terminal order by date separating under one month and under one
// year.
func printIssues(issue map[int]Issue, index []date) {

	now := date{}
	now.y, now.m, now.d = time.Now().Date()
	fmt.Printf("%d issues:\n", len(index))

	// Print issues that are less than a month old.
	fmt.Println("\nLess than a month")
	for n, date := range index {
		item := issue[date.r]
		if lessThanMonth(date, now) {
			fmt.Printf("#%-5d %9.9s %55.55s  %.2d %s %d\n",
				item.Number, item.User.Login, item.Title,
				date.d, date.m, date.y)
			// Verify date logic
			if date.p == true {
				fmt.Println("error: this line has already been printed")
			}
			index[n].p = true
		}
	}

	// Print issues that are less than a year old.
	fmt.Println("\nLess than a year")
	for n, date := range index {
		item := issue[date.r]
		if monthToAYear(date, now) {
			fmt.Printf("#%-5d %9.9s %55.55s  %.2d %s %d\n",
				item.Number, item.User.Login, item.Title,
				date.d, date.m, date.y)
			// Verify date logic
			if date.p == true {
				fmt.Println("error: this line has already been printed")
			}
			index[n].p = true
		}
	}

	// Print issues that are older than a year.
	fmt.Println("\nOlder than a year")
	for n, date := range index {
		item := issue[date.r]
		if yearOnward(date, now) {
			fmt.Printf("#%-5d %9.9s %55.55s  %.2d %s %d\n",
				item.Number, item.User.Login, item.Title,
				date.d, date.m, date.y)
			// Verify date logic
			if date.p == true {
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

// 1) The same calendar month and year.
// 2) The same calendar year with one month difference; lesser day than the day
//	issued.
// 3) One calendar year difference; January issue date December; lesser day
//	today than the day issued.
func lessThanMonth(d1, d2 date) bool {
	return d1.y == d2.y && d1.m == d2.m ||
		d1.y == d2.y-1 && d1.m == d2.m && d1.d > d2.d ||
		d1.y == d2.y-1 && d2.m == 1 && d1.m == 12 && d1.d > d2.d
}

// 1) The same calendar year; Greater month by more than one calendar month
//	than the date issued.
// 2) The same calendar year; The month is one greater than the issued month;
//	equal or greater day today than the day issued.
// 3) One calendar years difference; January the date of issue in December;
//	Equal or greater day today than the day issued.
// 4) One calendar years difference; The current calendar month is lesser than
//	the month of the issue.
// 5) One calendar years difference; The same calendar month; lesser day today
//	than the day issued.
func monthToAYear(d1, d2 date) bool {
	return d1.y == d2.y && d1.m < d2.m-1 ||
		d1.y == d2.y && d1.m == d2.m-1 && d1.d <= d2.d ||
		d1.y == d2.y-1 && d2.m == 1 && d1.m == 12 && d1.d <= d2.d ||
		d1.y == d2.y-1 && d1.m > d2.m ||
		d1.y == d2.y-1 && d1.m == d2.m && d1.d > d2.d
}

// 1) One calendar years difference; The same calendar month: Equal or greater
//	day than the day issued.
// 2) Greater than one years difference.
func yearOnward(d1, d2 date) bool {
	return d1.y < d2.y-1 && d1.m == d2.m && d1.d <= d2.d ||
		d1.y < d2.y-1
}
