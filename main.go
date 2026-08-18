package main

import (
	"encoding/json"
	"fmt"
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
	file, err := os.Open("symmer.json")
	if err != nil {
		if os.IsNotExist(err) {
			utils.SymmerPrint("Missing config, creating...")
			file = CreateConfig()
			fmt.Println("Config created, fill it out and rerun symmer.")
			os.Exit(1)
		}
	}
	return file
}

// Creates a symmer json config file when called.
func CreateConfig() *os.File {
	file, err := os.Create("symmer.json")

	if err != nil {
		utils.SymmerError("Unable to create json config file in current directory. Review error below")
		utils.SymmerError(err.Error())
		os.Exit(1)
	}

	return file
}

// Opens the json config file, decodes it using the Config type template, returns it as a typed object for parsing.
func OpenJSON(file *os.File) MapConfig {
	var cfg MapConfig
	
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		utils.SymmerError("Error decoding json")
		utils.SymmerPrint("Check if your config has been filled out with applications and sym locations")
		os.Exit(1)
	}


	utils.SymmerPrint("Config loaded successfully.")
	return cfg
}

// Creates a symlink based on the configuration options given in the json config file
func CreateSymLink(config string, symlinkLocation string, configName string) error {
	if err := os.Symlink(config, symlinkLocation); err != nil {
		utils.SymmerError("Unable to create a symlink for " + configName + ". Review error below.")
		utils.SymmerError(err.Error())
		return err
	} else {
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
