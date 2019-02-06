package github

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
)

var config = Config{}
var def = Config{
	Login:   "8i8",
	Editor:  "gvim",
	Mode:    "list",
	Request: Request{Owner: "8i8", Repo: "test"},
}

func init() {
	var queries []string
	queries = append(queries, "json")
	queries = append(queries, "decoder")
	def.Queries = queries
}

// Load the last saved config, if no file exists create with the default
// settings.
func LoadConfig(c Config) error {

	// Open file for reading if present.
	file, err := os.Open("config.json")
	if err != nil {
		// If not present create file
		if err = setConfig(def); err != nil {
			return err
		}
	}

	// decode json from within the file.
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		msg := "Config struct not created"
		fmt.Fprintf(os.Stderr, "error: %s: %v %v: %s\n", msg, file,
			line, err.Error())
	}

	// Check for modifications, with input settings, resave if required.
	compConfig(c, def)

	return err
}

// Compaire and update the configuration as required.
func compConfig(c, def Config) {
	if c.Repo == "" {
		c.Repo = def.Repo
	}
	if len(c.Queries) == 0 {
		c.Queries = def.Queries
	}
	if c.Token == "" {
		c.Token = def.Token
	}
	if c.Editor == "" {
		c.Editor = def.Editor
	}
	setConfig(c)
}

// Create a base Config file, used in the case that no config file is present.
func setConfig(c Config) error {

	var data []byte

	// Set the config from input.
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
	c.Token = ""
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
	file.Close()

	return nil
}
