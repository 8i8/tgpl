/*
Package github - Command line client for the github issue API.

SYNOPSIS
	github [user | repo | number][Oauth2][options]

DESCRIPTION
	github is a github client designed for raising and tracking and
	updating github issues on the github platform from the users command
	line by way of the github HTTP API. Giving the user access from the
	command line or their favorite editor application.

MAIN
	The github program has essentially five running modes, the mode is set
	from the main function according to the flags set state, defined in the
	SetState() function and three response reactions translating into three
	more sub states the combination of which defines the running of the
	program.

PROGRAM STATES
	Table representation of program states, the program has essentially
	five different primary states and three further states which comprise
	all subroutines, all of which is preset by the SetState() function
	establishing the type of HTTP request that is required. The second
	defines the formation of the expectation and treatment of the HTTP
	response.
URL MODE
	There are tow types of url formation, uAddr in which the url provides
	an explicit location and uSear with which the search server folder
	is defined and search values are given by way of query key vale pares.

	┌─────┬─────┬─────┬─────┬─────┬─────┬───────┬───────┬───────┐
	│     │     │     │     │-r   │     │       │       │       │
	│-o or│     │     │     │-l lo│     │       │       │       │
	│-a au│repo │numbe│token│-e ed│edito│ State │ url   │ Respo │
	│-u us│-r   │-n   │-t   │-x ra│-d   │       │ type  │       │
	├─────┼─────┼─────┼─────┼─────┼─────┼───────┼───────┼───────┤
	│ N/A │ N/A │ N/A │     │ -r  │ all │ mRaw  │ user  │ input │
	├─────┼─────┼─────┼─────┼─────┼─────┼───────┼───────┼───────┤
	│ yes │     │     │ N/A │ N/A │ all │ mList │ uSear │ rMany │
	├─────┼─────┼─────┼─────┼─────┼─────┼───────┼───────┼───────┤
	│     │ yes │     │ N/A │ N/A │ all │ mList │ uSear │ rMany │
	├─────┼─────┼─────┼─────┼─────┼─────┼───────┼───────┼───────┤
	│ yes │ yes │     │ N/A │ N/A │ all │ mList │ uSear │ rMany │
	├─────┼─────┼─────┼─────┼─────┼─────┼───────┼───────┼───────┤
	│ yes │     │ yes │ N/A │ N/A │ all │ mList │ uSear │ rMany │
	├─────┼─────┼─────┼─────┼─────┼─────┼───────┼───────┼───────┤
	│     │ yes │ yes │ N/A │ N/A │ all │ mList │ uSear │ rMany │
	├─────┼─────┼─────┼─────┼─────┼─────┼───────┼───────┼───────┤
	│ yes │ yes │ yes │ N/A │ N/A │ all │ mRead │ uAdRe │ rLone │
	├─────┼─────┼─────┼─────┼─────┼─────┼───────┼───────┼───────┤
	│ yes │ yes │ yes │ yes │ -x  │ all │ mRais │ uAdRw │ rNone │
	├─────┼─────┼─────┼─────┼─────┼─────┼───────┼───────┼───────┤
	│ yes │ yes │ yes │ yes │ -e  │ all │ mEdit │ uAdRe │ rLone │
	│     │     │     │     │     │     │ mEdit │ uAdWr │ rNone │
	├─────┼─────┼─────┼─────┼─────┼─────┼───────┼───────┼───────┤
	│ yes │ yes │ yes │ yes │ -l  │ all │ mLock │ uAdRe │ rLone │
	│     │     │     │     │     │     │ mLock │ uAdWr │ rNone │
	└─────┴─────┴─────┴─────┴─────┴─────┴───────┴───────┴───────┘

	-o org   -r repo  -n number
	-a auth
	-u user

	-l lock
	-m lock reason
	-e edit
	-x raise
	-d editor
	-v displays verbose report of the programs actions.

HTTP REQUESTS
	┌───────┬───────┬───────┬───────┬───────┬───────┐
	│       │ GET   │ POST  │ PATCH │ PUT   │ DELETE│
	├───────┼───────┼───────┼───────┼───────┼───────┤
	│ list  │   1   │       │       │       │       │
	├───────┼───────┼───────┼───────┼───────┼───────┤
	│ read  │   1   │       │       │       │       │
	├───────┼───────┼───────┼───────┼───────┼───────┤
	│ raise │       │   1   │       │       │       │
	├───────┼───────┼───────┼───────┼───────┼───────┤
	│ edit  │       │       │   1   │       │       │
	├───────┼───────┼───────┼───────┼───────┼───────┤
	│ lock  │       │       │       │   1   │       │
	├───────┼───────┼───────┼───────┼───────┼───────┤
	│unlock │       │       │       │       │   1   │
	└───────┴───────┴───────┴───────┴───────┴───────┘

	GET    /issues
 	GET    /user/issues
 	GET    /orgs/:org/issues
	GET    /search/issues?q= user:[user] | repo:[repo] | author:[author]
 	GET    /repos/:owner/:repo/issues
 	GET    /repos/:owner/:repo/issues/:number
 	POST   /repos/:owner/:repo/issues
 	PATCH  /repos/:owner/:repo/issues/:number
 	PUT    /repos/:owner/:repo/issues/:number/lock?lock_reason=[reason]
	DELETE /repos/:owner/:repo/issues/:number/lock

	https://api.github.com/search/issues

*/
package github

// GET https://api.github.com/repos/golang/go/issues?q=json+decoder

// var query1 []string
// var query2 []string

// func init() {
// 	query1 = append(query1, "is:open")
// 	query2 = append(query2, "repo:golang/go")
// 	query2 = append(query2, "is:open")
// }

//URL := "https://api.github.com/users/octocat/orgs"
//URL := "https://api.github.com/orgs/octokit/repos"
//URL := "https://api.github.com/search/issues?q=repo:8i8/test"
//URL := "https://api.github.com/repos/8i8/test/issues"
//URL := "https://api.github.com/repos/8i8/test/issues"

//URL := "https://api.github.com/users/octocat/orgs"
//URL := "https://api.github.com/orgs/octokit/repos"
//URL := "https://api.github.com/search/issues?q=repo:8i8/test"
//URL := "https://api.github.com/repos/8i8/test/issues"
//URL := "https://api.github.com/repos/8i8/test/issues"

// var query []string
// query = append(query, "repo:8i8/test")
// query = append(query, "is:open")
// q := url.QueryEscape(strings.Join(terms, " "))
