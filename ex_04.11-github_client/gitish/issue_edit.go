package gitish

// EditIssue opens and edits and existing issue.
// TODO NOW is it possible to have only one result possibele and negate the need
// for a test for 1?
func EditIssue(conf Config) error {

	// Serch for requested issue.
	// var issue []*Issue
	// issue = searchIssues(conf)
	// if err != nil {
	// 	return fmt.Errorf("edit: %v", err)
	// }

	// Validate results array.
	// err = checkResults(conf, issue)
	// if err != nil {
	// 	return fmt.Errorf("checkResults: %v", err)
	// }

	// // Edit existing data.
	// json, err := editIssue(conf, *issue[0])
	// if err != nil {
	// 	return fmt.Errorf("editIssue: %v", err)
	// }

	// // Update the issue with new data.
	// err = writeIssue(conf, json)
	// if err != nil {
	// 	return fmt.Errorf("writeIssue: %v", err)
	// }

	return nil
}

// checkResults verifies the resulting array of issues.
func checkResults(conf Config, issue []*Issue) error {
	// If issue not found or to many or locked.
	// l := len(issue)
	// if l == 0 {
	// 	return fmt.Errorf("issue %s not found", conf.Number)
	// } else if l > 1 {
	// 	return fmt.Errorf("multiple issues returned, require one")
	// }
	// if issue[0].Locked {
	// 	return fmt.Errorf("the issue is currently locked")
	// }
	return nil
}
