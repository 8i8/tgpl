package gitish

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
   Configuration
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// Config is a struct specific to the program that contains the principal
// program settings.
type Config struct {
	User   string // Repository owner,
	Token  string // Flag defined Oauth token.
	Editor string // Flag defined external editor.
	Req           // Stores the users request data.
	State         // The state of the program.
}

// State is an anonymous struct of only one single instance per request, it is
// contained within the Config struct.
type State struct {
	Mode    int    // Program running mode.
	Edit    bool   // Signal request to edit an issue.
	Lock    bool   // Lock state.
	Reason  string // Reason for lock.
	Verbose bool   // Signals the program print out extra detail.
}

// Request is a struct containing the details of a particular request.
type Req struct {
	Author  string   // Author user name.
	Org     string   // Organisation.
	Repo    string   // Repository name.
	Number  string   // Issue number.
	Queries []string // GET queries retrieved from the Args[] array.
}

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
