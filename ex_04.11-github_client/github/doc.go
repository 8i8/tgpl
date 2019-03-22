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
	├─────┼─────┼─────┼─────┼─────┼─────┼───────┼───────┼───────┤
	│ yes │     │     │ N/A │ N/A │ all │ cLIST │ uSear │ cMANY │ A
	├─────┼─────┼─────┼─────┼─────┼─────┼───────┼───────┼───────┤
	│     │ yes │     │ N/A │ N/A │ all │ cLIST │ uSear │ cMANY │ A
	├─────┼─────┼─────┼─────┼─────┼─────┼───────┼───────┼───────┤
	│ yes │ yes │     │ N/A │ N/A │ all │ cLIST │ uSear │ cMANY │ A
	├─────┼─────┼─────┼─────┼─────┼─────┼───────┼───────┼───────┤
	│ yes │     │ yes │ N/A │ N/A │ all │ cLIST │ uSear │ cMANY │ A
	├─────┼─────┼─────┼─────┼─────┼─────┼───────┼───────┼───────┤
	│     │ yes │ yes │ N/A │ N/A │ all │ cLIST │ uSear │ cMANY │ A
	├─────┼─────┼─────┼─────┼─────┼─────┼───────┼───────┼───────┤
	│ yes │ yes │ yes │ N/A │ N/A │ all │ cREAD │ uAdRe │ cLONE │ B
	├─────┼─────┼─────┼─────┼─────┼─────┼───────┼───────┼───────┤
	│ yes │ yes │ yes │ yes │ -x  │ all │ cRAISE│ uAdWr │ cNONE │ C
	├─────┼─────┼─────┼─────┼─────┼─────┼───────┼───────┼───────┤
	│ yes │ yes │ yes │ yes │ -e  │ all │ cREAD │ uAdRe │ cLONE │ B
	│     │     │     │     │     │     │ cEDIT │ uAdWr │ cNONE │ D
	├─────┼─────┼─────┼─────┼─────┼─────┼───────┼───────┼───────┤
	│ yes │ yes │ yes │ yes │ -l  │ all │ cREAD │ uAdRe │ cLONE │ B
	│     │     │     │     │     │     │ cLOCK │ uAdWr │ cNONE │ E
	└─────┴─────┴─────┴─────┴─────┴─────┴───────┴───────┴───────┘

	A read list
	B read record
	C raise new
	D edit record
	E Lock record

	[mode]
	edit
	raise
	search
	read

	[flag|data]
	-o org   -r repo  -n number
	-a auth
	-u user

	[mode/one of the following reasons]
	lock
		* off-topic
		* too heated
		* resolved
		* spam

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

	GET    /issues          (login required, lists all issues assigned to user)
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
