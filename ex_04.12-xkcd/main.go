package main

import (
	"encoding/json"
	"fmt"

	"tgpl/ex_04.12-xkcd/httpm"
	"tgpl/ex_04.12-xkcd/xkcd"
)

func main() {
	var err error
	req := httpm.HttpRequest{}
	req.GET("https://xkcd.com/108/info.0.json")
	req, err = httpm.Request(req, nil)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	var comic xkcd.Comic
	if err := json.Unmarshal(req.Msg.(json.RawMessage), &comic); err != nil {
		fmt.Printf("json decoder Msg failed: %v\n", err)
	}

	fmt.Printf("%+v", comic)
}
