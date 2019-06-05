package github

import (
	"bufio"
	"fmt"
	"os"
)

// ConnectDetails loads the github api conection details from file.
func ConnectDetails(path string) (tokens []string, err error) {

	file, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf("os.Open: %v", err)
		return
	}
	defer file.Close()

	content := bufio.NewScanner(file)
	var str []string
	ok := true
	for ok {
		ok = content.Scan()
		str = append(str, content.Text())
	}

	return str, nil
}
