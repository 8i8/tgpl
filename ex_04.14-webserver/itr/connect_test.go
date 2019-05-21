package itr

import (
	"fmt"
	"testing"
)

func TestConnectDetails(t *testing.T) {
	user, repo, token, err := connectDetails("../connect")
	if err != nil {
		t.Errorf("TestConnectionDetails: %v", err)
	}
	fmt.Println(user)
	fmt.Println(repo)
	fmt.Println(token)
}
