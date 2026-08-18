package main

import (
	"encoding/json"
	"os"
	"symmer/utils"
)

// Config type for json structure
type Config struct {
	CurrentConfigLocation string `json:"config"`
	FuturedSymLocation string `json:"symlink"`
}

// Mapped config of the json structure with a string acting as the title of the config
// Example: Opencode { "config": "location", "symlink": "location"}
type MapConfig map[string]Config

// Checks if the config for symmer exists, if not, it creates one in the current directory.
func ConfigCheck() *os.File {
	// Attempts to open the json file if it exists.
	file, err := os.Open("symmer.json")

	// Error check if it cannot be opened
	if err != nil {
		// If the file doesn't exist, then it'll create the file and exit
		if os.IsNotExist(err) {
			utils.SymmerPrint("Missing config, creating...")
			file = CreateConfig()
			utils.SymmerPrint("Config created, fill it out and rerun symmer.")
			os.Exit(0)
		}
	}
	return file
}

// Creates a symmer json config file when called.
func CreateConfig() *os.File {
	// Attempts to create config file
	file, err := os.Create("symmer.json")
	
	// if the file exists then it'll skip the error check
	if os.IsExist(err) {
		goto Return
	} else {
		s := "{}\n"
		file.Write([]byte(s))
	}

	// Presents the user with an error stating that the json config can't be created
	if err != nil {
		utils.SymmerError("Unable to create json config file in current directory. Review error below")
		utils.SymmerError(err.Error())
		os.Exit(1)
	}
	
	Return:
	return file
}

// Opens the json config file, decodes it using the Config type template, returns it as a typed object for parsing.
func OpenJSON(file *os.File) MapConfig {
	// Creates a config object to hold the future decoded json
	var cfg MapConfig
	
	// Attempts to decode the json and presents the user with an error message if unable to create it.
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		utils.SymmerError("Error decoding json")
		utils.SymmerPrint("Check if your config has been filled out with applications and sym locations")
		os.Exit(1)
	}

	
	// Success message and returns the decoded config
	utils.SymmerPrint("Config loaded successfully.")
	return cfg
}

// Creates a symlink based on the configuration options given in the json config file
func CreateSymLink(config string, symlinkLocation string, configName string) error {
	// Attempts to create the symlink. If unable to create one, it will print an error message for the user and return the error.
	if err := os.Symlink(config, symlinkLocation); err != nil {
		utils.SymmerError("Unable to create a symlink for " + configName + ". Review error below.")
		utils.SymmerError(err.Error())
		return err
	} else {
		// Returns nil if there is no issue with creating the symlink
		return nil
	}
}

func main() {
	// Opens the file after it checks if it exists
	file := ConfigCheck()
	defer file.Close()
	
	// Assigns cfg as the parsed config file
	cfg := OpenJSON(file)

	// Loop over every entry in config file 
	for k, v := range cfg {
		// Attempts to create a symlink via the given paths
		err := CreateSymLink(v.CurrentConfigLocation, v.FuturedSymLocation, k)
		
		// If an error occurs, prompts the user that symmer cant make the symlink and provides the error for the user to review
		if err != nil {
			utils.SymmerError("Unable to create symlink for " + k + ". Review error below.")
			utils.SymmerError(err.Error())
		} else {
			utils.SymmerPrint("Symlink for " + k + " created successfully!")
		}
	}
}
