package itr

import (
	"html/template"
	"io"
	"log"
)

// var issueList = template.Must(template.New("issuelist").Parse(`
// <h1>issues</h1>
// <tr style='text-align: left'>
// <table>
// <tr style='text-align: left'>
// 	<th>#</th>
// 	<th>State</th>
// 	<th>User</th>
// 	<th>Title</th>
// </tr>
// {{range .Issues}}
// <tr>
// 	<td><a href='{{.HtmlURL}}'>{{.Number}}</a></td>
// 	<td>{{.State}}</td>
// 	<td><a href='{{index .Users .User}}'>{{index .Users .Login}}</a></td>
// 	<td><a href='{{.HtmlURL}}'>{{.Title}}</a></td>
// </tr>
// {{end}}
// </table>
// `))

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>issues</h1>
<tr style='text-align: left'>
<table>
<tr style='text-align: left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
</tr>
{{range .Issues}}
<tr>
	<td><a href='{{.HtmlURL}}'>{{.Number}}</a></td>
	<td>{{.State}}</td>
 	<td><a href='{{(index $.Users .User).HtmlURL}}'>{{(index $.Users .User).Login}}</a></td>
	<td><a href='{{.HtmlURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func HtmlReport(w io.Writer, cache Cache) {

	if err := issueList.Execute(w, cache); err != nil {
		log.Fatal(err)
	}
}
