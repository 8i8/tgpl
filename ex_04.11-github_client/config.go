package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
)

type Configuration struct {
	Repo    string
	Queries string
	Token   string
}

func (c Configuration) Strings() []string {
	var s []string
	s = append(s, c.Repo)
	s = append(s, c.Queries)
	return s
}

var Config = Configuration{}

// Load the previous config used, else create a default.
func loadConfig() error {
	file, err := os.Open("config.json")
	if err != nil {
		if err = setDefConfig(); err != nil {
			panic(err)
		}
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		msg := "Config struct not created"
		fmt.Fprintf(os.Stderr, "error: %s: %v %v: %s\n", msg, file, line, err.Error())
	}
	return err
}

// Create a Config file, for the case of no file present.
func setDefConfig() error {

	Config.Repo = "repo:golang/go"
	Config.Queries = "is:open json decoder"
	var data []byte

	file, err := os.Create("config.json")
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		fmt.Fprintf(os.Stderr, "error: %v %v: %s\n", file, line, err.Error())
		return err
	}

	if data, err = json.MarshalIndent(Config, "", "	"); err != nil {
		_, file, line, _ := runtime.Caller(0)
		fmt.Fprintf(os.Stderr, "error: %v %v: %s\n", file, line, err.Error())
		return err
	}

	if _, err = file.Write(data); err != nil {
		_, file, line, _ := runtime.Caller(0)
		fmt.Fprintf(os.Stderr, "error: %v %v: %s\n", file, line, err.Error())
		return err
	}
	return nil
}
