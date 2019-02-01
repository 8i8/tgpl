package github

import "testing"

const t_IssuesAddrURL = "https://api.github.com/repos/golang/go/issues"
const t_IssuesQueryURL = "https://api.github.com/search/issues"

// GET https://api.github.com/repos/golang/go/issues?q=json+decoder

var query1 []string
var query2 []string

func init() {
	query1 = append(query1, "is:open")
	query2 = append(query2, "repo:golang/go")
	query2 = append(query2, "is:open")
}

func BenchmarkSearchIssuesQuery(b *testing.B) {
	for i := 0; i < 5; i++ {
		SearchIssues(t_IssuesQueryURL, query2)
	}
}

func BenchmarkSearchIssuesAddr(b *testing.B) {
	for i := 0; i < 5; i++ {
		SearchIssues(t_IssuesAddrURL, query1)
	}
}
