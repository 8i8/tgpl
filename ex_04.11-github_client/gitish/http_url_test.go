package gitish

// go run main.go -u 8i8
const A = "GET https://api.gitish.com/search/issues?q=user%3A8i8"

// go run main.go -a 8i8
const B = "https://api.gitish.com/search/issues?q=author%3A8i8"

// go run main.go -u 8i8 -r test
const C = "https://api.gitish.com/search/issues?q=repo%3A8i8%2Ftest"
