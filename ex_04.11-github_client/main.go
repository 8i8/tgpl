package issues

import (
	"flag"

	"tgpl/ex_04.11-github_client/github"
)

var ioUpdate = flag.Bool("u", false, "update issue")
var ioIssue = flag.Bool("n", false, "new issue")
var ioAddress = flag.Bool("s", false, "set repo address")
var ioToken = flag.Bool("t", false, "set token")
var ioEditor = flag.Bool("e", false, "set editor")
var ioDelete = flag.Bool("d", false, "delete issue")

func main() {

	if err := loadConfig(); err != nil {
		panic(err)
	}

	flag.Parse()
	if *ioUpdate {
	} else if *ioAddress {
	} else if *ioToken {
	} else if *ioEditor {
	} else if *ioIssue {
		github.RaiseIssue()
	} else if *ioDelete {
	} else {
		github.ListIssues(Config.Strings())
	}
}
