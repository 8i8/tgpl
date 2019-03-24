package gitish

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
   Configuration
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

type Flags uint

// Config is a struct specific to the program that contains the principal
// program settings.
type Config struct {
	Mode   Flags  // Program running mode.
	User   string // Repository owner,
	Token  string // Flag defined Oauth token.
	Editor string // Flag defined external editor.
	Reason string // Reason for lock.
	Req           // Stores the users request data.
}

// Req request is a struct containing all of the details of a particular request.
type Req struct {
	Author  string   // Author user name.
	Org     string   // Organisation.
	Repo    string   // Repository name.
	Number  string   // Issue number.
	Queries []string // Queries that have been retrieved from the Args[] array.
}

// TODO clean up the configs role
// // config holds the programs global configuration.
// var config = Config{}

// // def: Default configuration.
// var def = Config{
// 	User:    "8i8",
// 	Editor:  "gvim",
// 	Request: Request{Author: "8i8", Repo: "test"},
// }

// func init() {
// 	var queries []string
// 	queries = append(queries, "json")
// 	queries = append(queries, "decoder")
// 	def.Queries = queries
// }

// // LoadConfig load the last saved config, if no file exists create with the
// // default settings.
// func LoadConfig(c Config) error {

// 	// Open file for reading if present.
// 	file, err := os.Open("config.json")
// 	if err != nil {
// 		// If not present create file
// 		if err = setConfig(def); err != nil {
// 			return fmt.Errorf("setConfig: %v", err)
// 		}
// 	}

// 	// decode json from within the file.
// 	decoder := json.NewDecoder(file)
// 	err = decoder.Decode(&config)
// 	if err != nil {
// 		return fmt.Errorf("decoder: %v", err)
// 	}

// 	// Check for modifications, with input settings, re-save if required.
// 	// compConfig(c, def)

// 	return err
// }

// // // Compare and update the configuration as required.
// // func compConfig(c, def Config) {

// // 	if len(c.Author) == 0 {
// // 		c.Author = def.Author
// // 	}
// // 	if len(c.Editor) == 0 {
// // 		c.Editor = def.Editor
// // 	}
// // 	if len(c.Mode) == 0 {
// // 		c.Mode = def.Mode
// // 	}
// // 	_ = setConfig(c)
// // }

// // Create a base Config file, use in the case that no configuration file present.
// func setConfig(c Config) error {

// 	var data []byte

// 	// Set the global configuration from the input.
// 	config = c

// 	// Make the file.
// 	file, err := os.Create("config.json")
// 	if err != nil {
// 		return fmt.Errorf("File create failed: %v", err)
// 	}
// 	defer file.Close()

// 	// Prepare the json in readable fashon.
// 	c.Token = ""
// 	if data, err = json.MarshalIndent(config, "", "	"); err != nil {
// 		return fmt.Errorf("MarshalIndent: %v", err)
// 	}

// 	// Write json to file.
// 	if _, err = file.Write(data); err != nil {
// 		return fmt.Errorf("File write failed: %v", err)
// 	}

// 	return err
// }
