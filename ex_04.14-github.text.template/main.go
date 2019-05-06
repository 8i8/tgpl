package main

import "tgpl/ex_04.14-github.text.template/github"

import (
	"html/template"
	"log"
	"os"
	"time"

	"tgpl/ex_04.14-github.text.template/dates"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}--------------------------------------------------------------------------
Number:	{{.Number}}
User:	{{.User.Login}}
Title:	{{.Title | printf "%.64s"}}
Age:	{{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func Report() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func ReportList() {
	// Run search issues to get data from github.
	reply, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	// Make and fill a map from the result array.
	err = dates.ListIssues(reply)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	Report()
}
