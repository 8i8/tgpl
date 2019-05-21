/*
Package github - Command line client for the github issue API.

SYNOPSIS
	github [user | repo | number][Oauth2][options]

DESCRIPTION
	github is a github client built for raising, tracking and updating git
	issues on the github platform. Run from the users command line
	accessing the github API; Affording the user access from the command
	line or their favorite editor application.

MAIN
	The github program has essentially five running modes, the mode is set
	from the main function according to the flags set state, defined in the
	SetBitState() function and three response reactions translating into three
	more sub states the combination of which defines the running of the
	program.

PROGRAM STATES
	Table representation of program states, the program has essentially
	five different primary states and three further states which comprise
	all subroutines, all of which is preset in the SetBitState() function
	essentialy establishing the type of HTTP request that is required. The
	second defines the formation of the expectation and treatment of the
	HTTP response.

URL MODE
	There are two types of url formation, uAddr in which the url provides
	an explicit location and uSear with which the search server folder
	is defined and search values given by way of query key value pairs.

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
	E lock/unlock record

MODES
	github -[mode]

	-read	Read an existing issue.
	-list	List all issues for a specific repo or user.
	-edit	Edit an existing issue.
	-raise	Raise a new issue.
	-lock	Lock an issue.
	-unlock	Unlock an issue.

FLAGS
	github	-[flag]

	-v	Verbose mode, gives detailed description of the programs
		actions.
	-h
	-help	Prints out the programs help file.

	github	-[flag] [value]

	-u	User name.
	-a	Author.
	-o	Organisation name.
	-r	Repository name.
	-n	Issue number, requires that author and repository also be
		defined.
	-t	OAuth2 token.
	-d	External editor launch command.
	-l	Provide a reason for locking.

	[lock rational]

		* off-topic
		* too heated
		* resolved
		* spam

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

	URL usable by the API

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

/*
	API endpoints

curl https://api.github.com
{
  "current_user_url": "https://api.github.com/user",
  "current_user_authorizations_html_url": "https://github.com/settings/connections/applications{/client_id}",
  "authorizations_url": "https://api.github.com/authorizations",
  "code_search_url": "https://api.github.com/search/code?q={query}{&page,per_page,sort,order}",
  "commit_search_url": "https://api.github.com/search/commits?q={query}{&page,per_page,sort,order}",
  "emails_url": "https://api.github.com/user/emails",
  "emojis_url": "https://api.github.com/emojis",
  "events_url": "https://api.github.com/events",
  "feeds_url": "https://api.github.com/feeds",
  "followers_url": "https://api.github.com/user/followers",
  "following_url": "https://api.github.com/user/following{/target}",
  "gists_url": "https://api.github.com/gists{/gist_id}",
  "hub_url": "https://api.github.com/hub",
  "issue_search_url": "https://api.github.com/search/issues?q={query}{&page,per_page,sort,order}",
  "issues_url": "https://api.github.com/issues",
  "keys_url": "https://api.github.com/user/keys",
  "notifications_url": "https://api.github.com/notifications",
  "organization_repositories_url": "https://api.github.com/orgs/{org}/repos{?type,page,per_page,sort}",
  "organization_url": "https://api.github.com/orgs/{org}",
  "public_gists_url": "https://api.github.com/gists/public",
  "rate_limit_url": "https://api.github.com/rate_limit",
  "repository_url": "https://api.github.com/repos/{owner}/{repo}",
  "repository_search_url": "https://api.github.com/search/repositories?q={query}{&page,per_page,sort,order}",
  "current_user_repositories_url": "https://api.github.com/user/repos{?type,page,per_page,sort}",
  "starred_url": "https://api.github.com/user/starred{/owner}{/repo}",
  "starred_gists_url": "https://api.github.com/gists/starred",
  "team_url": "https://api.github.com/teams",
  "user_url": "https://api.github.com/users/{user}",
  "user_organizations_url": "https://api.github.com/user/orgs",
  "user_repositories_url": "https://api.github.com/users/{user}/repos{?type,page,per_page,sort}",
  "user_search_url": "https://api.github.com/search/users?q={query}{&page,per_page,sort,order}"
}
*/
