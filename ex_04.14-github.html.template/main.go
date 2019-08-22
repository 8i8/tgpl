package main

import "tgpl/ex_04.14-github.text.template/github"

import (
	"html/template"
	"log"
	"os"
	"time"
)

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  HTML report, write to a file and then view in a browser.
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<tr style='text-align: left'>
<table>
<tr style='text-align: left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
</tr>
{{range .Items}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func HtmlReport() {

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  Basic report, run in the terminal to see stdout.
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

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

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

func main() {

	Report()
}
