package github

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
)

var config = Config{}
var def = Config{
	User:    User{},
	Request: Request{Name: User{}, Repo: "8i8/test", Queries: "is:open"},
	Token:   "",
	Editor:  "gvim",
}

// Load the last saved config, if no file exists create with the default
// settings.
func LoadConfig(c Config) error {
	// Read file if present.
	file, err := os.Open("config.json")
	if err != nil {
		// If not present create file
		if err = setConfig(def); err != nil {
			return err
		}
	}
	// decode json from file.
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		msg := "Config struct not created"
		fmt.Fprintf(os.Stderr, "error: %s: %v %v: %s\n", msg, file,
			line, err.Error())
	}

	// Check for modifications, if required resave.
	compConfig(config, c)

	return err
}

// Compaire and update the configuration if required.
func compConfig(c1, c2 Config) {
	if c1 == c2 {
		return
	}
	if c2.Repo != "" {
		c1.Repo = c2.Repo
	}
	if c2.Queries != "" {
		c1.Queries = c2.Queries
	}
	if c2.Token != "" {
		c1.Token = c2.Token
	}
	if c2.Editor != "" {
		c1.Editor = c2.Editor
	}
	setConfig(c1)
}

// Create a Config file, used in the case that no config file is present.
func setConfig(c Config) error {

	var data []byte

	// Set the config from.
	config = c

	// Make the file.
	file, err := os.Create("config.json")
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		fmt.Fprintf(os.Stderr, "error: %v %v: %s\n", file, line,
			err.Error())
		return err
	}

	// Prepare the json in readable fashon.
	if data, err = json.MarshalIndent(config, "", "	"); err != nil {
		_, file, line, _ := runtime.Caller(0)
		fmt.Fprintf(os.Stderr, "error: %v %v: %s\n", file, line,
			err.Error())
		return err
	}

	// Write json to file.
	if _, err = file.Write(data); err != nil {
		_, file, line, _ := runtime.Caller(0)
		fmt.Fprintf(os.Stderr, "error: %v %v: %s\n", file, line,
			err.Error())
		return err
	}

	return nil
}
