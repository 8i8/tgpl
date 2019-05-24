package github

import (
	"fmt"
	"testing"
)

func TestConnectDetails(t *testing.T) {
	user, repo, token, err := ConnectDetails("../connect")
	if err != nil {
		t.Errorf("TestConnectionDetails: %v", err)
	}
	fmt.Println(user)
	fmt.Println(repo)
	fmt.Println(token)
}
