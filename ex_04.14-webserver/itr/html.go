package itr

import (
	"html/template"
	"io"
	"log"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{ len .Issues }} issues</h1>
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

var issueMap = template.Must(template.New("issuelist").Parse(`
<a href='/'>issues</a>
<a href='/users'>users</a>
<a href='/milestones'>milestones</a>
<h1>{{ len .IssuesIndex }} issues</h1>
<tr style='text-align: left'>
<table>
<tr style='text-align: left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
</tr>
{{ range $index, $value := .IssuesIndex }}
<tr>
	<td><a href='{{ (index $.Issues $value).HtmlURL }}'>{{ (index $.Issues $value).Number }}</a></td>
	<td>{{(index $.Issues $value).State}}</td>
 	<td><a href='{{ (index $.Users (index $.Issues $value).User).HtmlURL }}'>{{ (index $.Users (index $.Issues $value).User).Login }}</a></td>
	<td><a href='{{ (index $.Issues $value).HtmlURL }}'>{{ (index $.Issues $value).Title }}</a></td>
</tr>
{{end}}
</table>
`))

func HtmlIssueReport(w io.Writer, cache Cache) {

	if err := issueMap.Execute(w, cache); err != nil {
		log.Fatal(err)
	}
}

var userMap = template.Must(template.New("userlist").Parse(`
<a href='/'>issues</a>
<a href='/users'>users</a>
<a href='/milestones'>milestones</a>
<h1>{{ len .UsersIndex }} users</h1>
<tr style='text-align: left'>
<table>
<tr style='text-align: left'>
	<th>#</th>
	<th>One</th>
	<th>Two</th>
	<th>Three</th>
</tr>
{{ range $index, $value := .UsersIndex }}
<tr>
	<td><img src="{{ (index $.Users $value).AvatarURL }}" alt="There should be a picture here" width="60">
	<a href='{{ (index $.Users $value).HtmlURL }}'>{{ (index $.Users $value).Login }}</a></td>
</tr>
{{end}}
</table>
`))

func HtmlUserReport(w io.Writer, cache Cache) {

	if err := userMap.Execute(w, cache); err != nil {
		log.Fatal(err)
	}
}

var milestoneMap = template.Must(template.New("milestonelist").Parse(`
<a href='/'>issues</a>
<a href='/users'>users</a>
<a href='/milestones'>milestones</a>
<h1>{{ len .MilestonesIndex }} milestones</h1>
<tr style='test-align: left'>
<table>
<tr style='test-align: left'>
	<th>milestones</th>
</tr>
{{ range $index, $value := .MilestonesIndex }}
<tr>
	<td><a href='{{ (index $.Milestones $value).HtmlURL }}'>{{ (index $.Milestones $value).Description }}</a></td>
</tr>
{{end}}
</table>
`))

func HtmlMilestoneReport(w io.Writer, cache Cache) {

	if err := milestoneMap.Execute(w, cache); err != nil {
		log.Fatal(err)
	}
}
