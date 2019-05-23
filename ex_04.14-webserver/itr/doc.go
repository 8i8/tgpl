/*
	Exercise 4.14 ~ Create a web server that queries GitHub once and then allows
	navigation of the list of bug reports, milestones, and users.

	Connection details
		Set user name and repo.
		Set autentication token.

	Load relevent data from cache.

	Query github
		If no data in cache download all.
		Update current records if they do exist; Only issues updated at
		or after this time are returned.
			This is a timestamp in ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ.
			?since=YYYY-MM-DDTHH:MM:SSZ
		https://developer.github.com/v3/issues/#parameters-1

		If data already in cache, download any that have been modified
		since the last update date and time.

		Updade any issues that have been modified by using the ETag or
		Last-Modified header.
		'If-None-Match: "644b5b0155e6404a9cc4bd9d8b1ae730"'

		or

		"If-Modified-Since: Thu, 05 Jul 2012 15:31:30 GMT"
		use Issue.UpdatedAt for last update time.
		https://developer.github.com/v3/#conditional-requests

		use
			?state=all to get both open and closed issues.
			&per_page=100 to increase the pageinate amount.

		generate indexe lists and sort

	Write new data to local cache.
	Open bug reports page in list view.
	Open milestones
*/

package itr
