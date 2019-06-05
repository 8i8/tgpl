package github

import (
	"fmt"
	"strings"
	"testing"
)

func TestConnectDetails(t *testing.T) {

	s := []string{"golang", "go", "123456"}

	tokens, err := ConnectDetails("connect")
	if err != nil {
		t.Errorf("TestConnectionDetails: %v: %v", err, tokens)
	}

	if len(tokens) == 0 {
		t.Errorf("TestConnectionDetails: tokens is empty: %v", tokens)
		return
	}

	for i, token := range tokens {
		if len(token) != 0 {
			if strings.Compare(token, s[i]) != 0 {
				t.Errorf("error: recieved `%v` expected `%v`", token, s[i])
			}
		}
	}

	for _, token := range tokens {
		fmt.Println(token)
	}
}
