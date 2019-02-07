package github

import (
	"encoding/json"
	"fmt"
	"os"
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
			return fmt.Errorf("setConfig: ", err)
		}
	}

	// decode json from within the file.
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return fmt.Errorf("decoder: %v", err)
	}

	// Check for modifications, with input settings, resave if required.
	compConfig(c, def)

	return err
}

// Compaire and update the configuration as required.
func compConfig(c, def Config) {

	if len(c.Login) == 0 {
		c.Login = def.Login
	}
	if len(c.Editor) == 0 {
		c.Editor = def.Editor
	}
	if len(c.Mode) == 0 {
		c.Mode = def.Mode
	}
	setConfig(c)
}

// Create a base Config file, used in the case that no config file is present.
func setConfig(c Config) error {

	var data []byte

	// Set the global config from the input.
	config = c

	// Make the file.
	file, err := os.Create("config.json")
	if err != nil {
		return fmt.Errorf("File create failed: %v", err)
	}
	defer file.Close()

	// Prepare the json in readable fashon.
	c.Token = ""
	if data, err = json.MarshalIndent(config, "", "	"); err != nil {
		return fmt.Errorf("MarshalIndent: %v", err)
	}

	// Write json to file.
	if _, err = file.Write(data); err != nil {
		return fmt.Errorf("File write failed: %v", err)
	}

	return err
}
