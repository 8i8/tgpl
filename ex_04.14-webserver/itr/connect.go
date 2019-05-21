package itr

import (
	"bufio"
	"fmt"
	"os"
)

// connectDetails loads the github api conection details from file.
func connectDetails(path string) (user, repo, token string, err error) {

	file, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf("os.Open: %v", err)
		return
	}
	defer file.Close()

	content := bufio.NewScanner(file)
	var str []string
	for i := 0; i < 3; i++ {
		content.Scan()
		str = append(str, content.Text())
	}
	user = str[0]
	repo = str[1]
	token = str[2]

	return
}
